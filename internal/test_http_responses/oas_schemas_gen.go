// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/netip"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/nonrecording"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = bytes.NewReader
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = io.Copy
	_ = math.Mod
	_ = big.Rat{}
	_ = bits.LeadingZeros64
	_ = net.IP{}
	_ = http.MethodGet
	_ = netip.Addr{}
	_ = url.URL{}
	_ = regexp.MustCompile
	_ = sort.Ints
	_ = strconv.ParseInt
	_ = strings.Builder{}
	_ = sync.Pool{}
	_ = time.Time{}

	_ = errors.Is
	_ = jx.Null
	_ = uuid.UUID{}
	_ = otel.GetTracerProvider
	_ = attribute.KeyValue{}
	_ = codes.Unset
	_ = metric.MeterConfig{}
	_ = syncint64.Counter(nil)
	_ = nonrecording.NewNoopMeterProvider
	_ = trace.TraceIDFromHex

	_ = conv.ToInt32
	_ = ht.NewRequest
	_ = json.Marshal
	_ = otelogen.Version
	_ = uri.PathEncoder{}
	_ = validate.Int{}
)

type AnyContentTypeBinaryStringSchemaDefaultDef struct {
	Data io.Reader
}

func (s AnyContentTypeBinaryStringSchemaDefaultDef) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

// AnyContentTypeBinaryStringSchemaDefaultDefStatusCode wraps AnyContentTypeBinaryStringSchemaDefaultDef with StatusCode.
type AnyContentTypeBinaryStringSchemaDefaultDefStatusCode struct {
	StatusCode int
	Response   AnyContentTypeBinaryStringSchemaDefaultDef
}

type AnyContentTypeBinaryStringSchemaOK struct {
	Data io.Reader
}

func (s AnyContentTypeBinaryStringSchemaOK) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

// NewNilInt returns new NilInt with value set to v.
func NewNilInt(v int) NilInt {
	return NilInt{
		Value: v,
	}
}

// NilInt is nullable int.
type NilInt struct {
	Value int
	Null  bool
}

// SetTo sets value to v.
func (o *NilInt) SetTo(v int) {
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o NilInt) IsNull() bool { return o.Null }

// Get returns value and boolean that denotes whether value was set.
func (o NilInt) Get() (v int, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

func (*NilInt) multipleGenericResponsesRes() {}

// NewNilString returns new NilString with value set to v.
func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}

// NilString is nullable string.
type NilString struct {
	Value string
	Null  bool
}

// SetTo sets value to v.
func (o *NilString) SetTo(v string) {
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o NilString) IsNull() bool { return o.Null }

// Get returns value and boolean that denotes whether value was set.
func (o NilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

func (*NilString) multipleGenericResponsesRes() {}

type OctetStreamBinaryStringSchemaOK struct {
	Data io.Reader
}

func (s OctetStreamBinaryStringSchemaOK) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}

type OctetStreamEmptySchemaOK struct {
	Data io.Reader
}

func (s OctetStreamEmptySchemaOK) Read(p []byte) (n int, err error) {
	return s.Data.Read(p)
}
