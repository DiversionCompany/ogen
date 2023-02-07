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
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}

var _ Handler = struct {
	*Client
}{}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
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
	return err
}

func (c *Client) sendDisjointSecurity(ctx context.Context) (res *DisjointSecurityOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("disjointSecurity"),
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
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/disjointSecurity"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BasicAuth"
			switch err := c.securityBasicAuth(ctx, "DisjointSecurity", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BasicAuth\"")
			}
		}
		{
			stage = "Security:QueryKey"
			switch err := c.securityQueryKey(ctx, "DisjointSecurity", r); err {
			case nil:
				satisfied[0] |= 1 << 1
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"QueryKey\"")
			}
		}
		{
			stage = "Security:CookieKey"
			switch err := c.securityCookieKey(ctx, "DisjointSecurity", r); err {
			case nil:
				satisfied[0] |= 1 << 2
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"CookieKey\"")
			}
		}
		{
			stage = "Security:HeaderKey"
			switch err := c.securityHeaderKey(ctx, "DisjointSecurity", r); err {
			case nil:
				satisfied[0] |= 1 << 3
			case ogenerrors.ErrSkipClientSecurity:
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
			return res, errors.New("no security requirement satisfied")
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
	return err
}

func (c *Client) sendIntersectSecurity(ctx context.Context) (res *IntersectSecurityOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("intersectSecurity"),
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
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/intersectSecurity"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BasicAuth"
			switch err := c.securityBasicAuth(ctx, "IntersectSecurity", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BasicAuth\"")
			}
		}
		{
			stage = "Security:HeaderKey"
			switch err := c.securityHeaderKey(ctx, "IntersectSecurity", r); err {
			case nil:
				satisfied[0] |= 1 << 1
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"HeaderKey\"")
			}
		}
		{
			stage = "Security:BearerToken"
			switch err := c.securityBearerToken(ctx, "IntersectSecurity", r); err {
			case nil:
				satisfied[0] |= 1 << 2
			case ogenerrors.ErrSkipClientSecurity:
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
			return res, errors.New("no security requirement satisfied")
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
	return err
}

func (c *Client) sendOptionalSecurity(ctx context.Context) (res *OptionalSecurityOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("optionalSecurity"),
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
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/optionalSecurity"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:QueryKey"
			switch err := c.securityQueryKey(ctx, "OptionalSecurity", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
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
			return res, errors.New("no security requirement satisfied")
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
