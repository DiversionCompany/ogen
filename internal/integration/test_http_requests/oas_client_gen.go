// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
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

// AllRequestBodies invokes allRequestBodies operation.
//
// POST /allRequestBodies
func (c *Client) AllRequestBodies(ctx context.Context, request AllRequestBodiesReq) (AllRequestBodiesOK, error) {
	res, err := c.sendAllRequestBodies(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendAllRequestBodies(ctx context.Context, request AllRequestBodiesReq) (res AllRequestBodiesOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("allRequestBodies"),
	}
	// Validate request before sending.
	switch request := request.(type) {
	case *AllRequestBodiesApplicationJSON:
		// Validation is not required for this type.
	case *AllRequestBodiesReqApplicationOctetStream:
		// Validation is not required for this type.
	case *AllRequestBodiesApplicationXWwwFormUrlencoded:
		// Validation is not required for this type.
	case *AllRequestBodiesMultipartFormData:
		// Validation is not required for this type.
	case *AllRequestBodiesReqTextPlain:
		// Validation is not required for this type.
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AllRequestBodies",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/allRequestBodies"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAllRequestBodiesRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAllRequestBodiesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AllRequestBodiesOptional invokes allRequestBodiesOptional operation.
//
// POST /allRequestBodiesOptional
func (c *Client) AllRequestBodiesOptional(ctx context.Context, request AllRequestBodiesOptionalReq) (AllRequestBodiesOptionalOK, error) {
	res, err := c.sendAllRequestBodiesOptional(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendAllRequestBodiesOptional(ctx context.Context, request AllRequestBodiesOptionalReq) (res AllRequestBodiesOptionalOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("allRequestBodiesOptional"),
	}
	// Validate request before sending.
	switch request := request.(type) {
	case *AllRequestBodiesOptionalReqEmptyBody:
		// Validation is not needed for the empty body type.
	case *AllRequestBodiesOptionalApplicationJSON:
		// Validation is not required for this type.
	case *AllRequestBodiesOptionalReqApplicationOctetStream:
		// Validation is not required for this type.
	case *AllRequestBodiesOptionalApplicationXWwwFormUrlencoded:
		// Validation is not required for this type.
	case *AllRequestBodiesOptionalMultipartFormData:
		// Validation is not required for this type.
	case *AllRequestBodiesOptionalReqTextPlain:
		// Validation is not required for this type.
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AllRequestBodiesOptional",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/allRequestBodiesOptional"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAllRequestBodiesOptionalRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAllRequestBodiesOptionalResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// Base64Request invokes base64Request operation.
//
// POST /base64Request
func (c *Client) Base64Request(ctx context.Context, request Base64RequestReq) (Base64RequestOK, error) {
	res, err := c.sendBase64Request(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendBase64Request(ctx context.Context, request Base64RequestReq) (res Base64RequestOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("base64Request"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "Base64Request",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/base64Request"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeBase64RequestRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeBase64RequestResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MaskContentType invokes maskContentType operation.
//
// POST /maskContentType
func (c *Client) MaskContentType(ctx context.Context, request *MaskContentTypeReqWithContentType) (*MaskResponse, error) {
	res, err := c.sendMaskContentType(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendMaskContentType(ctx context.Context, request *MaskContentTypeReqWithContentType) (res *MaskResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("maskContentType"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "MaskContentType",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/maskContentType"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeMaskContentTypeRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeMaskContentTypeResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MaskContentTypeOptional invokes maskContentTypeOptional operation.
//
// POST /maskContentTypeOptional
func (c *Client) MaskContentTypeOptional(ctx context.Context, request *MaskContentTypeOptionalReqWithContentType) (*MaskResponse, error) {
	res, err := c.sendMaskContentTypeOptional(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendMaskContentTypeOptional(ctx context.Context, request *MaskContentTypeOptionalReqWithContentType) (res *MaskResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("maskContentTypeOptional"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "MaskContentTypeOptional",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/maskContentTypeOptional"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeMaskContentTypeOptionalRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeMaskContentTypeOptionalResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
