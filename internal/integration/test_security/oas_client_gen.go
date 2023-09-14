// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// DisjointSecurity invokes disjointSecurity operation.
	//
	// GET /disjointSecurity
	DisjointSecurity(ctx context.Context) error
	// IntersectSecurity invokes intersectSecurity operation.
	//
	// GET /intersectSecurity
	IntersectSecurity(ctx context.Context) error
	// OptionalSecurity invokes optionalSecurity operation.
	//
	// GET /optionalSecurity
	OptionalSecurity(ctx context.Context) error
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
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
		sec:        sec,
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

// DisjointSecurity invokes disjointSecurity operation.
//
// GET /disjointSecurity
func (c *Client) DisjointSecurity(ctx context.Context) error {
	res, err := c.sendDisjointSecurity(ctx)
	_ = res
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "DisjointSecurity", "GET", "/disjointSecurity", err)
	}
	return err
}

func (c *Client) sendDisjointSecurity(ctx context.Context) (res *DisjointSecurityOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("disjointSecurity"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/disjointSecurity"),
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
	ctx, span := c.cfg.Tracer.Start(ctx, "DisjointSecurity",
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
	pathParts[0] = "/disjointSecurity"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BasicAuth"
			switch err := c.securityBasicAuth(ctx, "DisjointSecurity", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BasicAuth\"")
			}
		}
		{
			stage = "Security:QueryKey"
			switch err := c.securityQueryKey(ctx, "DisjointSecurity", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 1
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"QueryKey\"")
			}
		}
		{
			stage = "Security:CookieKey"
			switch err := c.securityCookieKey(ctx, "DisjointSecurity", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 2
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"CookieKey\"")
			}
		}
		{
			stage = "Security:HeaderKey"
			switch err := c.securityHeaderKey(ctx, "DisjointSecurity", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 3
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"HeaderKey\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000011},
				{0b00001100},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeDisjointSecurityResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// IntersectSecurity invokes intersectSecurity operation.
//
// GET /intersectSecurity
func (c *Client) IntersectSecurity(ctx context.Context) error {
	res, err := c.sendIntersectSecurity(ctx)
	_ = res
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "IntersectSecurity", "GET", "/intersectSecurity", err)
	}
	return err
}

func (c *Client) sendIntersectSecurity(ctx context.Context) (res *IntersectSecurityOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("intersectSecurity"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/intersectSecurity"),
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
	ctx, span := c.cfg.Tracer.Start(ctx, "IntersectSecurity",
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
	pathParts[0] = "/intersectSecurity"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BasicAuth"
			switch err := c.securityBasicAuth(ctx, "IntersectSecurity", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BasicAuth\"")
			}
		}
		{
			stage = "Security:HeaderKey"
			switch err := c.securityHeaderKey(ctx, "IntersectSecurity", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 1
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"HeaderKey\"")
			}
		}
		{
			stage = "Security:BearerToken"
			switch err := c.securityBearerToken(ctx, "IntersectSecurity", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 2
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerToken\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000011},
				{0b00000110},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeIntersectSecurityResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OptionalSecurity invokes optionalSecurity operation.
//
// GET /optionalSecurity
func (c *Client) OptionalSecurity(ctx context.Context) error {
	res, err := c.sendOptionalSecurity(ctx)
	_ = res
	if err != nil && c.cfg.errorMiddleware != nil {
		err = c.cfg.errorMiddleware(ctx, "OptionalSecurity", "GET", "/optionalSecurity", err)
	}
	return err
}

func (c *Client) sendOptionalSecurity(ctx context.Context) (res *OptionalSecurityOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("optionalSecurity"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/optionalSecurity"),
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
	ctx, span := c.cfg.Tracer.Start(ctx, "OptionalSecurity",
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
	pathParts[0] = "/optionalSecurity"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:QueryKey"
			switch err := c.securityQueryKey(ctx, "OptionalSecurity", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"QueryKey\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{},
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeOptionalSecurityResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
