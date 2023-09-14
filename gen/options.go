package gen

import (
	"context"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"

	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/internal/location"
	"github.com/ogen-go/ogen/internal/urlpath"
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/openapi"
)

// RemoteOptions is remote reference resolver options.
type RemoteOptions = jsonschema.ExternalOptions

// Options is Generator options.
type Options struct {
	// NoClient disables client generation.
	NoClient bool
	// NoServer disables server generation.
	NoServer bool
	// NoWebhookClient disables webhook client generation.
	NoWebhookClient bool
	// NoWebhookServer disables webhook server generation.
	NoWebhookServer bool

	// GenerateExampleTests whether to generate decoding tests using schema examples.
	GenerateExampleTests bool
	// SkipTestRegex is regex to skip generated tests.
	SkipTestRegex *regexp.Regexp
	// SkipUnimplemented disables generation of unimplemented Handler, like UnimplementedServer generated by gRPC.
	SkipUnimplemented bool
	// InferSchemaType enables type inference for schemas. Schema parser will try to detect schema type
	// by its properties.
	InferSchemaType bool
	// SchemaDepthLimit is maximum depth of schema generation. Default is 1000.
	SchemaDepthLimit int
	// CustomFormats sets custom formats.
	CustomFormats CustomFormatsMap

	// AllowRemote enables remote references resolving.
	//
	// See https://github.com/ogen-go/ogen/issues/385.
	AllowRemote bool
	// RootURL is root URL for remote references resolving.
	RootURL *url.URL
	// Remote is remote reference resolver options.
	Remote RemoteOptions

	// Filters contains filters to skip operations.
	Filters Filters
	// IgnoreNotImplemented contains ErrNotImplemented messages to ignore.
	IgnoreNotImplemented []string
	// NotImplementedHook is hook for ErrNotImplemented errors.
	NotImplementedHook func(name string, err error)

	// ConvenientErrors control Convenient Errors feature.
	//
	// Default value is `auto` (0), NewError handler will be generated if possible.
	//
	// If value > 0 forces feature. An error will be returned if generator is unable to find common error pattern.
	//
	// If value < 0 disables feature entirely.
	ConvenientErrors ConvenientErrors

	// ContentTypeAliases contains content type aliases.
	ContentTypeAliases ContentTypeAliases

	// ExpandSpec is a path to expanded spec.
	ExpandSpec string

	// ShouldSkipDefaultResp whether to skip default responses.
	SkipDefaultResp bool

	// OptionalIsNullable is whether optional fields are nullable.
	OptionalIsNullable bool

	SkipSecurities bool

	// File is the file that is being parsed.
	//
	// Used for error messages.
	File location.File
	// Logger to use.
	Logger *zap.Logger
}

// SetLocation sets File, RootURL and RemoteOptions using given path or URL
// and returns file data.
func (o *Options) SetLocation(p string, opts RemoteOptions) ([]byte, error) {
	o.Remote = opts
	r := jsonschema.NewExternalResolver(opts)

	containsDrive := runtime.GOOS == "windows" && filepath.VolumeName(p) != ""
	if u, _ := url.Parse(p); u != nil && !containsDrive && u.Scheme != "" {
		switch u.Scheme {
		case "http", "https":
			_, fileName := path.Split(u.Path)

			// FIXME(tdakkota): pass context.
			data, err := r.Get(context.Background(), p)
			if err != nil {
				return nil, err
			}

			o.RootURL = u
			o.File = location.NewFile(fileName, p, data)
			// Guard against reading local files in remote mode.
			o.Remote.ReadFile = func(p string) ([]byte, error) {
				return nil, errors.New("local files are not supported in remote mode")
			}

			return data, nil
		case "file":
			toPath := opts.URLToFilePath
			if toPath == nil {
				toPath = urlpath.URLToFilePath
			}

			converted, err := toPath(u)
			if err != nil {
				return nil, errors.Wrap(err, "convert url to file path")
			}
			p = converted
		default:
			return nil, errors.Errorf("unsupported scheme %q", u.Scheme)
		}
	}
	p = filepath.Clean(p)

	abs, err := filepath.Abs(p)
	if err != nil {
		return nil, err
	}
	_, fileName := filepath.Split(p)

	readFile := o.Remote.ReadFile
	if readFile == nil {
		readFile = os.ReadFile
	}

	data, err := readFile(p)
	if err != nil {
		return nil, err
	}

	u, err := urlpath.URLFromFilePath(abs)
	if err != nil {
		return nil, errors.Wrap(err, "convert file path to url")
	}

	o.RootURL = u
	o.File = location.NewFile(fileName, p, data)
	return data, nil
}

func (o *Options) setDefaults() {
	if o.SchemaDepthLimit <= 0 {
		o.SchemaDepthLimit = defaultSchemaDepthLimit
	}
	if o.Logger == nil {
		o.Logger = zap.NewNop()
	}
}

// ConvenientErrors is an option type to control `Convenient Errors` feature.
type ConvenientErrors int

// IsDisabled whether Convenient Errors is disabled.
func (c ConvenientErrors) IsDisabled() bool {
	return c < 0
}

// IsForced whether Convenient Errors is forced.
func (c ConvenientErrors) IsForced() bool {
	return c > 0
}

// String implements fmt.Stringer.
func (c ConvenientErrors) String() string {
	switch {
	case c < 0:
		return "off"
	case c > 0:
		return "on"
	default:
		return "auto"
	}
}

// IsBoolFlag implements flag.boolFlag.
func (c *ConvenientErrors) IsBoolFlag() bool {
	return true
}

// Set implements flag.Value.
func (c *ConvenientErrors) Set(s string) error {
	switch s {
	case "auto":
		*c = 0
		return nil
	case "on", "true":
		*c = 1
		return nil
	case "off", "false":
		*c = -1
		return nil
	default:
		return errors.Errorf(`expected "on", "off" or "auto", got %q`, s)
	}
}

// ContentTypeAliases maps content type to concrete ogen encoding.
type ContentTypeAliases map[string]ir.Encoding

// String implements fmt.Stringer.
func (m ContentTypeAliases) String() string {
	var (
		b     strings.Builder
		first = true
	)
	for k, v := range m {
		if first {
			first = false
		} else {
			b.WriteString(",")
		}
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(v.String())
	}
	return b.String()
}

// Set implements flag.Value.
func (m *ContentTypeAliases) Set(value string) error {
	if *m == nil {
		*m = ContentTypeAliases{}
	}
	before, after, ok := strings.Cut(value, "=")
	if !ok {
		return errors.Errorf("invalid mapping %q", value)
	}
	(*m)[before] = ir.Encoding(after)
	return nil
}

// Filters contains filters to skip operations.
type Filters struct {
	PathRegex *regexp.Regexp
	Methods   []string
}

func (f Filters) accept(op *openapi.Operation) bool {
	if f.PathRegex != nil && !f.PathRegex.MatchString(op.Path.String()) {
		return false
	}

	if len(f.Methods) > 0 {
		return slices.ContainsFunc(f.Methods, func(m string) bool { return strings.EqualFold(m, op.HTTPMethod) })
	}

	return true
}

// CustomFormatsMap is map of custom formats.
type CustomFormatsMap = map[jsonschema.SchemaType]map[string]CustomFormatDef

// CustomFormatDef defines custom format type.
type CustomFormatDef struct {
	typ  reflect.Type
	json reflect.Type
	text reflect.Type
}

// CustomFormat returns custom format definition.
func CustomFormat[
	T any,
	JSON interface {
		~struct{} // Enforce implementation without state.

		EncodeJSON(*jx.Encoder, T)
		DecodeJSON(*jx.Decoder) (T, error)
	},
	Text interface {
		~struct{} // Enforce implementation without state.

		EncodeText(T) string
		DecodeText(string) (T, error)
	},
]() CustomFormatDef {
	return CustomFormatDef{
		typ:  reflect.TypeOf(new(T)).Elem(),
		json: reflect.TypeOf(new(JSON)).Elem(),
		text: reflect.TypeOf(new(Text)).Elem(),
	}
}
