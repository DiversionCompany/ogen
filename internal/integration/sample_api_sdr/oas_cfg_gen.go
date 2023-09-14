// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"fmt"
	"math/big"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/ogenregex"
	"github.com/ogen-go/ogen/otelogen"
)

var regexMap = map[string]ogenregex.Regexp{
	"^\\d-\\d$":                      ogenregex.MustCompile("^\\d-\\d$"),
	"^variant3_[^\r\n\u2028\u2029]*": ogenregex.MustCompile("^variant3_[^\r\n\u2028\u2029]*"),
	"foo[^\r\n\u2028\u2029]*":        ogenregex.MustCompile("foo[^\r\n\u2028\u2029]*"),
	"string_[^\r\n\u2028\u2029]*":    ogenregex.MustCompile("string_[^\r\n\u2028\u2029]*"),
}
var ratMap = map[string]*big.Rat{
	"10": func() *big.Rat {
		r, ok := new(big.Rat).SetString("10")
		if !ok {
			panic(fmt.Sprintf("rat %q: can't parse", "10"))
		}
		return r
	}(),
	"5": func() *big.Rat {
		r, ok := new(big.Rat).SetString("5")
		if !ok {
			panic(fmt.Sprintf("rat %q: can't parse", "5"))
		}
		return r
	}(),
}
var (
	// Allocate option closure once.
	clientSpanKind = trace.WithSpanKind(trace.SpanKindClient)
	// Allocate option closure once.
	serverSpanKind = trace.WithSpanKind(trace.SpanKindServer)
)

type (
	optionFunc[C any] func(*C)
	otelOptionFunc    func(*otelConfig)
)

type otelConfig struct {
	TracerProvider trace.TracerProvider
	Tracer         trace.Tracer
	MeterProvider  metric.MeterProvider
	Meter          metric.Meter
}

func (cfg *otelConfig) initOTEL() {
	if cfg.TracerProvider == nil {
		cfg.TracerProvider = otel.GetTracerProvider()
	}
	if cfg.MeterProvider == nil {
		cfg.MeterProvider = otel.GetMeterProvider()
	}
	cfg.Tracer = cfg.TracerProvider.Tracer(otelogen.Name,
		trace.WithInstrumentationVersion(otelogen.SemVersion()),
	)
	cfg.Meter = cfg.MeterProvider.Meter(otelogen.Name)
}

// ErrorHandler is error handler.
type ErrorHandler = ogenerrors.ErrorHandler

type serverConfig struct {
	otelConfig
	NotFound           http.HandlerFunc
	MethodNotAllowed   func(w http.ResponseWriter, r *http.Request, allowed string)
	ErrorHandler       ErrorHandler
	Prefix             string
	Middleware         Middleware
	MaxMultipartMemory int64
}

// ServerOption is server config option.
type ServerOption interface {
	applyServer(*serverConfig)
}

var _ = []ServerOption{
	(optionFunc[serverConfig])(nil),
	(otelOptionFunc)(nil),
}

func (o optionFunc[C]) applyServer(c *C) {
	o(c)
}

func (o otelOptionFunc) applyServer(c *serverConfig) {
	o(&c.otelConfig)
}

func newServerConfig(opts ...ServerOption) serverConfig {
	cfg := serverConfig{
		NotFound: http.NotFound,
		MethodNotAllowed: func(w http.ResponseWriter, r *http.Request, allowed string) {
			w.Header().Set("Allow", allowed)
			w.WriteHeader(http.StatusMethodNotAllowed)
		},
		ErrorHandler:       ogenerrors.DefaultErrorHandler,
		Middleware:         nil,
		MaxMultipartMemory: 32 << 20, // 32 MB
	}
	for _, opt := range opts {
		opt.applyServer(&cfg)
	}
	cfg.initOTEL()
	return cfg
}

type baseServer struct {
	cfg      serverConfig
	requests metric.Int64Counter
	errors   metric.Int64Counter
	duration metric.Float64Histogram
}

func (s baseServer) notFound(w http.ResponseWriter, r *http.Request) {
	s.cfg.NotFound(w, r)
}

func (s baseServer) notAllowed(w http.ResponseWriter, r *http.Request, allowed string) {
	s.cfg.MethodNotAllowed(w, r, allowed)
}

func (cfg serverConfig) baseServer() (s baseServer, err error) {
	s = baseServer{cfg: cfg}
	if s.requests, err = s.cfg.Meter.Int64Counter(otelogen.ServerRequestCount); err != nil {
		return s, err
	}
	if s.errors, err = s.cfg.Meter.Int64Counter(otelogen.ServerErrorsCount); err != nil {
		return s, err
	}
	if s.duration, err = s.cfg.Meter.Float64Histogram(otelogen.ServerDuration); err != nil {
		return s, err
	}
	return s, nil
}

type ClientErrorMiddleware func(ctx context.Context, opName, opMethod, opPath string, err error) error

type clientConfig struct {
	otelConfig
	Client          ht.Client
	errorMiddleware ClientErrorMiddleware
}

// ClientOption is client config option.
type ClientOption interface {
	applyClient(*clientConfig)
}

var _ = []ClientOption{
	(optionFunc[clientConfig])(nil),
	(otelOptionFunc)(nil),
}

func (o optionFunc[C]) applyClient(c *C) {
	o(c)
}

func (o otelOptionFunc) applyClient(c *clientConfig) {
	o(&c.otelConfig)
}

func newClientConfig(opts ...ClientOption) clientConfig {
	cfg := clientConfig{
		Client: http.DefaultClient,
	}
	for _, opt := range opts {
		opt.applyClient(&cfg)
	}
	cfg.initOTEL()
	return cfg
}

type baseClient struct {
	cfg      clientConfig
	requests metric.Int64Counter
	errors   metric.Int64Counter
	duration metric.Float64Histogram
}

func (cfg clientConfig) baseClient() (c baseClient, err error) {
	c = baseClient{cfg: cfg}
	if c.requests, err = c.cfg.Meter.Int64Counter(otelogen.ClientRequestCount); err != nil {
		return c, err
	}
	if c.errors, err = c.cfg.Meter.Int64Counter(otelogen.ClientErrorsCount); err != nil {
		return c, err
	}
	if c.duration, err = c.cfg.Meter.Float64Histogram(otelogen.ClientDuration); err != nil {
		return c, err
	}
	return c, nil
}

// Option is config option.
type Option interface {
	ServerOption
	ClientOption
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
//
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return otelOptionFunc(func(cfg *otelConfig) {
		if provider != nil {
			cfg.TracerProvider = provider
		}
	})
}

// WithMeterProvider specifies a meter provider to use for creating a meter.
//
// If none is specified, the otel.GetMeterProvider() is used.
func WithMeterProvider(provider metric.MeterProvider) Option {
	return otelOptionFunc(func(cfg *otelConfig) {
		if provider != nil {
			cfg.MeterProvider = provider
		}
	})
}

// WithClient specifies http client to use.
func WithClient(client ht.Client) ClientOption {
	return optionFunc[clientConfig](func(cfg *clientConfig) {
		if client != nil {
			cfg.Client = client
		}
	})
}

func WithErrorMiddleware(errorMiddleware ClientErrorMiddleware) ClientOption {
	return optionFunc[clientConfig](func(cfg *clientConfig) {
		cfg.errorMiddleware = errorMiddleware
	})
}

// WithNotFound specifies Not Found handler to use.
func WithNotFound(notFound http.HandlerFunc) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		if notFound != nil {
			cfg.NotFound = notFound
		}
	})
}

// WithMethodNotAllowed specifies Method Not Allowed handler to use.
func WithMethodNotAllowed(methodNotAllowed func(w http.ResponseWriter, r *http.Request, allowed string)) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		if methodNotAllowed != nil {
			cfg.MethodNotAllowed = methodNotAllowed
		}
	})
}

// WithErrorHandler specifies error handler to use.
func WithErrorHandler(h ErrorHandler) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		if h != nil {
			cfg.ErrorHandler = h
		}
	})
}

// WithPathPrefix specifies server path prefix.
func WithPathPrefix(prefix string) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		cfg.Prefix = prefix
	})
}

// WithMiddleware specifies middlewares to use.
func WithMiddleware(m ...Middleware) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		switch len(m) {
		case 0:
			cfg.Middleware = nil
		case 1:
			cfg.Middleware = m[0]
		default:
			cfg.Middleware = middleware.ChainMiddlewares(m...)
		}
	})
}

// WithMaxMultipartMemory specifies limit of memory for storing file parts.
// File parts which can't be stored in memory will be stored on disk in temporary files.
func WithMaxMultipartMemory(max int64) ServerOption {
	return optionFunc[serverConfig](func(cfg *serverConfig) {
		if max > 0 {
			cfg.MaxMultipartMemory = max
		}
	})
}
