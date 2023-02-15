// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleDisjointSecurityRequest handles disjointSecurity operation.
//
// GET /disjointSecurity
func (s *Server) handleDisjointSecurityRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("disjointSecurity"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/disjointSecurity"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DisjointSecurity",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "DisjointSecurity",
			ID:   "disjointSecurity",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBasicAuth(ctx, "DisjointSecurity", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BasicAuth",
					Err:              err,
				}
				recordError("Security:BasicAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}
		{
			sctx, ok, err := s.securityQueryKey(ctx, "DisjointSecurity", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "QueryKey",
					Err:              err,
				}
				recordError("Security:QueryKey", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 1
				ctx = sctx
			}
		}
		{
			sctx, ok, err := s.securityCookieKey(ctx, "DisjointSecurity", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "CookieKey",
					Err:              err,
				}
				recordError("Security:CookieKey", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 2
				ctx = sctx
			}
		}
		{
			sctx, ok, err := s.securityHeaderKey(ctx, "DisjointSecurity", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "HeaderKey",
					Err:              err,
				}
				recordError("Security:HeaderKey", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 3
				ctx = sctx
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
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}

	var response *DisjointSecurityOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "DisjointSecurity",
			OperationID:   "disjointSecurity",
			Body:          nil,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *DisjointSecurityOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.DisjointSecurity(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.DisjointSecurity(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDisjointSecurityResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleIntersectSecurityRequest handles intersectSecurity operation.
//
// GET /intersectSecurity
func (s *Server) handleIntersectSecurityRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("intersectSecurity"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/intersectSecurity"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "IntersectSecurity",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "IntersectSecurity",
			ID:   "intersectSecurity",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBasicAuth(ctx, "IntersectSecurity", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BasicAuth",
					Err:              err,
				}
				recordError("Security:BasicAuth", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}
		{
			sctx, ok, err := s.securityHeaderKey(ctx, "IntersectSecurity", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "HeaderKey",
					Err:              err,
				}
				recordError("Security:HeaderKey", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 1
				ctx = sctx
			}
		}
		{
			sctx, ok, err := s.securityBearerToken(ctx, "IntersectSecurity", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "BearerToken",
					Err:              err,
				}
				recordError("Security:BearerToken", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 2
				ctx = sctx
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
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}

	var response *IntersectSecurityOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "IntersectSecurity",
			OperationID:   "intersectSecurity",
			Body:          nil,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *IntersectSecurityOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.IntersectSecurity(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.IntersectSecurity(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeIntersectSecurityResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleOptionalSecurityRequest handles optionalSecurity operation.
//
// GET /optionalSecurity
func (s *Server) handleOptionalSecurityRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("optionalSecurity"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/optionalSecurity"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "OptionalSecurity",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "OptionalSecurity",
			ID:   "optionalSecurity",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityQueryKey(ctx, "OptionalSecurity", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "QueryKey",
					Err:              err,
				}
				recordError("Security:QueryKey", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
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
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}

	var response *OptionalSecurityOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "OptionalSecurity",
			OperationID:   "optionalSecurity",
			Body:          nil,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *OptionalSecurityOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.OptionalSecurity(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.OptionalSecurity(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOptionalSecurityResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
