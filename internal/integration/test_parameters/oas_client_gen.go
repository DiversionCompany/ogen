// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// ComplicatedParameterNameGet invokes GET /complicatedParameterName operation.
	//
	// GET /complicatedParameterName
	ComplicatedParameterNameGet(ctx context.Context, params ComplicatedParameterNameGetParams) error
	// ContentParameters invokes contentParameters operation.
	//
	// GET /contentParameters/{path}
	ContentParameters(ctx context.Context, params ContentParametersParams) (*ContentParameters, error)
	// CookieParameter invokes cookieParameter operation.
	//
	// Test for cookie param.
	//
	// GET /cookieParameter
	CookieParameter(ctx context.Context, params CookieParameterParams) (*Value, error)
	// HeaderParameter invokes headerParameter operation.
	//
	// Test for header param.
	//
	// GET /headerParameter
	HeaderParameter(ctx context.Context, params HeaderParameterParams) (*Value, error)
	// ObjectCookieParameter invokes objectCookieParameter operation.
	//
	// GET /objectCookieParameter
	ObjectCookieParameter(ctx context.Context, params ObjectCookieParameterParams) (*OneLevelObject, error)
	// ObjectQueryParameter invokes objectQueryParameter operation.
	//
	// GET /objectQueryParameter
	ObjectQueryParameter(ctx context.Context, params ObjectQueryParameterParams) (*ObjectQueryParameterOK, error)
	// PathParameter invokes pathParameter operation.
	//
	// Test for path param.
	//
	// GET /pathParameter/{value}
	PathParameter(ctx context.Context, params PathParameterParams) (*Value, error)
	// SameName invokes sameName operation.
	//
	// Parameters with different location, but with the same name.
	//
	// GET /same_name/{param}
	SameName(ctx context.Context, params SameNameParams) error
	// SimilarNames invokes similarNames operation.
	//
	// Parameters with different location, but with similar names.
	//
	// GET /similarNames
	SimilarNames(ctx context.Context, params SimilarNamesParams) error
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Operations = &Client{}
var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// ComplicatedParameterNameGet invokes GET /complicatedParameterName operation.
//
// GET /complicatedParameterName
func (c *Client) ComplicatedParameterNameGet(ctx context.Context, params ComplicatedParameterNameGetParams) error {
	_, err := c.sendComplicatedParameterNameGet(ctx, params)
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "ComplicatedParameterNameGet", "GET", "/complicatedParameterName", err)
	}
	return err
}

func (c *Client) sendComplicatedParameterNameGet(ctx context.Context, params ComplicatedParameterNameGetParams) (res *ComplicatedParameterNameGetOK, err error) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/complicatedParameterName"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ComplicatedParameterNameGet",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/complicatedParameterName"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "=" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "=",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Eq))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "+" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "+",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Plus))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "question?" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "question?",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Question))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "and&" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "and&",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.And))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "percent%" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "percent%",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Percent))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeComplicatedParameterNameGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ContentParameters invokes contentParameters operation.
//
// GET /contentParameters/{path}
func (c *Client) ContentParameters(ctx context.Context, params ContentParametersParams) (*ContentParameters, error) {
	res, err := c.sendContentParameters(ctx, params)
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "ContentParameters", "GET", "/contentParameters/{path}", err)
	}
	return res, err
}

func (c *Client) sendContentParameters(ctx context.Context, params ContentParametersParams) (res *ContentParameters, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("contentParameters"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/contentParameters/{path}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ContentParameters",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/contentParameters/"
	{
		// Encode "path" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "path",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			enc := jx.GetEncoder()
			func(e *jx.Encoder) {
				params.Path.Encode(e)
			}(enc)
			return e.EncodeValue(string(enc.Bytes()))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "query" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "query",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			enc := jx.GetEncoder()
			func(e *jx.Encoder) {
				params.Query.Encode(e)
			}(enc)
			return e.EncodeValue(string(enc.Bytes()))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Header",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			enc := jx.GetEncoder()
			func(e *jx.Encoder) {
				params.XHeader.Encode(e)
			}(enc)
			return e.EncodeValue(string(enc.Bytes()))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "EncodeCookieParams"
	cookie := uri.NewCookieEncoder(r)
	{
		// Encode "cookie" parameter.
		cfg := uri.CookieParameterEncodingConfig{
			Name:    "cookie",
			Explode: true,
		}

		if err := cookie.EncodeParam(cfg, func(e uri.Encoder) error {
			enc := jx.GetEncoder()
			func(e *jx.Encoder) {
				params.Cookie.Encode(e)
			}(enc)
			return e.EncodeValue(string(enc.Bytes()))
		}); err != nil {
			return res, errors.Wrap(err, "encode cookie")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeContentParametersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CookieParameter invokes cookieParameter operation.
//
// Test for cookie param.
//
// GET /cookieParameter
func (c *Client) CookieParameter(ctx context.Context, params CookieParameterParams) (*Value, error) {
	res, err := c.sendCookieParameter(ctx, params)
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "CookieParameter", "GET", "/cookieParameter", err)
	}
	return res, err
}

func (c *Client) sendCookieParameter(ctx context.Context, params CookieParameterParams) (res *Value, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("cookieParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/cookieParameter"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CookieParameter",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/cookieParameter"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeCookieParams"
	cookie := uri.NewCookieEncoder(r)
	{
		// Encode "value" parameter.
		cfg := uri.CookieParameterEncodingConfig{
			Name:    "value",
			Explode: true,
		}

		if err := cookie.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Value))
		}); err != nil {
			return res, errors.Wrap(err, "encode cookie")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCookieParameterResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// HeaderParameter invokes headerParameter operation.
//
// Test for header param.
//
// GET /headerParameter
func (c *Client) HeaderParameter(ctx context.Context, params HeaderParameterParams) (*Value, error) {
	res, err := c.sendHeaderParameter(ctx, params)
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "HeaderParameter", "GET", "/headerParameter", err)
	}
	return res, err
}

func (c *Client) sendHeaderParameter(ctx context.Context, params HeaderParameterParams) (res *Value, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("headerParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/headerParameter"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "HeaderParameter",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/headerParameter"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Value",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XValue))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeHeaderParameterResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ObjectCookieParameter invokes objectCookieParameter operation.
//
// GET /objectCookieParameter
func (c *Client) ObjectCookieParameter(ctx context.Context, params ObjectCookieParameterParams) (*OneLevelObject, error) {
	res, err := c.sendObjectCookieParameter(ctx, params)
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "ObjectCookieParameter", "GET", "/objectCookieParameter", err)
	}
	return res, err
}

func (c *Client) sendObjectCookieParameter(ctx context.Context, params ObjectCookieParameterParams) (res *OneLevelObject, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("objectCookieParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/objectCookieParameter"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ObjectCookieParameter",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/objectCookieParameter"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeCookieParams"
	cookie := uri.NewCookieEncoder(r)
	{
		// Encode "value" parameter.
		cfg := uri.CookieParameterEncodingConfig{
			Name:    "value",
			Explode: false,
		}

		if err := cookie.EncodeParam(cfg, func(e uri.Encoder) error {
			return params.Value.EncodeURI(e)
		}); err != nil {
			return res, errors.Wrap(err, "encode cookie")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeObjectCookieParameterResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ObjectQueryParameter invokes objectQueryParameter operation.
//
// GET /objectQueryParameter
func (c *Client) ObjectQueryParameter(ctx context.Context, params ObjectQueryParameterParams) (*ObjectQueryParameterOK, error) {
	res, err := c.sendObjectQueryParameter(ctx, params)
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "ObjectQueryParameter", "GET", "/objectQueryParameter", err)
	}
	return res, err
}

func (c *Client) sendObjectQueryParameter(ctx context.Context, params ObjectQueryParameterParams) (res *ObjectQueryParameterOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("objectQueryParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/objectQueryParameter"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ObjectQueryParameter",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/objectQueryParameter"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "formObject" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "formObject",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FormObject.Get(); ok {
				return val.EncodeURI(e)
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "deepObject" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "deepObject",
			Style:   uri.QueryStyleDeepObject,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.DeepObject.Get(); ok {
				return val.EncodeURI(e)
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeObjectQueryParameterResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PathParameter invokes pathParameter operation.
//
// Test for path param.
//
// GET /pathParameter/{value}
func (c *Client) PathParameter(ctx context.Context, params PathParameterParams) (*Value, error) {
	res, err := c.sendPathParameter(ctx, params)
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "PathParameter", "GET", "/pathParameter/{value}", err)
	}
	return res, err
}

func (c *Client) sendPathParameter(ctx context.Context, params PathParameterParams) (res *Value, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pathParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/pathParameter/{value}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "PathParameter",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/pathParameter/"
	{
		// Encode "value" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "value",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Value))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodePathParameterResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SameName invokes sameName operation.
//
// Parameters with different location, but with the same name.
//
// GET /same_name/{param}
func (c *Client) SameName(ctx context.Context, params SameNameParams) error {
	_, err := c.sendSameName(ctx, params)
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "SameName", "GET", "/same_name/{param}", err)
	}
	return err
}

func (c *Client) sendSameName(ctx context.Context, params SameNameParams) (res *SameNameOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("sameName"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/same_name/{param}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "SameName",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/same_name/"
	{
		// Encode "param" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "param",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.PathParam))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "param" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "param",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.QueryParam))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeSameNameResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SimilarNames invokes similarNames operation.
//
// Parameters with different location, but with similar names.
//
// GET /similarNames
func (c *Client) SimilarNames(ctx context.Context, params SimilarNamesParams) error {
	_, err := c.sendSimilarNames(ctx, params)
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "SimilarNames", "GET", "/similarNames", err)
	}
	return err
}

func (c *Client) sendSimilarNames(ctx context.Context, params SimilarNamesParams) (res *SimilarNamesOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("similarNames"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/similarNames"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "SimilarNames",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/similarNames"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "x-param" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "x-param",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.QueryXParam))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Param",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.HeaderXParam))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeSimilarNamesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
