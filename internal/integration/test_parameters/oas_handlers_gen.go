// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleComplicatedParameterNameGetRequest handles GET /complicatedParameterName operation.
//
// GET /complicatedParameterName
func (s *Server) handleComplicatedParameterNameGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/complicatedParameterName"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ComplicatedParameterNameGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ComplicatedParameterNameGet",
			ID:   "",
		}
	)
	params, err := decodeComplicatedParameterNameGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *ComplicatedParameterNameGetOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "ComplicatedParameterNameGet",
			OperationSummary: "",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "=",
					In:   "query",
				}: params.Eq,
				{
					Name: "+",
					In:   "query",
				}: params.Plus,
				{
					Name: "question?",
					In:   "query",
				}: params.Question,
				{
					Name: "and&",
					In:   "query",
				}: params.And,
				{
					Name: "percent%",
					In:   "query",
				}: params.Percent,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ComplicatedParameterNameGetParams
			Response = *ComplicatedParameterNameGetOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackComplicatedParameterNameGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ComplicatedParameterNameGet(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.ComplicatedParameterNameGet(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeComplicatedParameterNameGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleContentParametersRequest handles contentParameters operation.
//
// GET /contentParameters/{path}
func (s *Server) handleContentParametersRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("contentParameters"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/contentParameters/{path}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ContentParameters",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ContentParameters",
			ID:   "contentParameters",
		}
	)
	params, err := decodeContentParametersParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *ContentParameters
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "ContentParameters",
			OperationSummary: "",
			OperationID:      "contentParameters",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "query",
					In:   "query",
				}: params.Query,
				{
					Name: "path",
					In:   "path",
				}: params.Path,
				{
					Name: "X-Header",
					In:   "header",
				}: params.XHeader,
				{
					Name: "cookie",
					In:   "cookie",
				}: params.Cookie,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ContentParametersParams
			Response = *ContentParameters
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackContentParametersParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ContentParameters(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ContentParameters(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeContentParametersResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleCookieParameterRequest handles cookieParameter operation.
//
// Test for cookie param.
//
// GET /cookieParameter
func (s *Server) handleCookieParameterRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("cookieParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/cookieParameter"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CookieParameter",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "CookieParameter",
			ID:   "cookieParameter",
		}
	)
	params, err := decodeCookieParameterParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Value
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "CookieParameter",
			OperationSummary: "",
			OperationID:      "cookieParameter",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "value",
					In:   "cookie",
				}: params.Value,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = CookieParameterParams
			Response = *Value
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackCookieParameterParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CookieParameter(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.CookieParameter(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCookieParameterResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleHeaderParameterRequest handles headerParameter operation.
//
// Test for header param.
//
// GET /headerParameter
func (s *Server) handleHeaderParameterRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("headerParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/headerParameter"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "HeaderParameter",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "HeaderParameter",
			ID:   "headerParameter",
		}
	)
	params, err := decodeHeaderParameterParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Value
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "HeaderParameter",
			OperationSummary: "",
			OperationID:      "headerParameter",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "X-Value",
					In:   "header",
				}: params.XValue,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = HeaderParameterParams
			Response = *Value
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackHeaderParameterParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.HeaderParameter(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.HeaderParameter(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeHeaderParameterResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleObjectCookieParameterRequest handles objectCookieParameter operation.
//
// GET /objectCookieParameter
func (s *Server) handleObjectCookieParameterRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("objectCookieParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/objectCookieParameter"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ObjectCookieParameter",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ObjectCookieParameter",
			ID:   "objectCookieParameter",
		}
	)
	params, err := decodeObjectCookieParameterParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *OneLevelObject
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "ObjectCookieParameter",
			OperationSummary: "",
			OperationID:      "objectCookieParameter",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "value",
					In:   "cookie",
				}: params.Value,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ObjectCookieParameterParams
			Response = *OneLevelObject
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackObjectCookieParameterParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ObjectCookieParameter(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ObjectCookieParameter(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeObjectCookieParameterResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleObjectQueryParameterRequest handles objectQueryParameter operation.
//
// GET /objectQueryParameter
func (s *Server) handleObjectQueryParameterRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("objectQueryParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/objectQueryParameter"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ObjectQueryParameter",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ObjectQueryParameter",
			ID:   "objectQueryParameter",
		}
	)
	params, err := decodeObjectQueryParameterParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *ObjectQueryParameterOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "ObjectQueryParameter",
			OperationSummary: "",
			OperationID:      "objectQueryParameter",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "formObject",
					In:   "query",
				}: params.FormObject,
				{
					Name: "deepObject",
					In:   "query",
				}: params.DeepObject,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ObjectQueryParameterParams
			Response = *ObjectQueryParameterOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackObjectQueryParameterParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ObjectQueryParameter(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ObjectQueryParameter(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeObjectQueryParameterResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handlePathParameterRequest handles pathParameter operation.
//
// Test for path param.
//
// GET /pathParameter/{value}
func (s *Server) handlePathParameterRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pathParameter"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/pathParameter/{value}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PathParameter",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "PathParameter",
			ID:   "pathParameter",
		}
	)
	params, err := decodePathParameterParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Value
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "PathParameter",
			OperationSummary: "",
			OperationID:      "pathParameter",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "value",
					In:   "path",
				}: params.Value,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = PathParameterParams
			Response = *Value
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackPathParameterParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.PathParameter(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.PathParameter(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePathParameterResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSameNameRequest handles sameName operation.
//
// Parameters with different location, but with the same name.
//
// GET /same_name/{param}
func (s *Server) handleSameNameRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("sameName"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/same_name/{param}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SameName",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SameName",
			ID:   "sameName",
		}
	)
	params, err := decodeSameNameParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *SameNameOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "SameName",
			OperationSummary: "parameters with different location, but with the same name",
			OperationID:      "sameName",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "param",
					In:   "path",
				}: params.PathParam,
				{
					Name: "param",
					In:   "query",
				}: params.QueryParam,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SameNameParams
			Response = *SameNameOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSameNameParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.SameName(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.SameName(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSameNameResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSimilarNamesRequest handles similarNames operation.
//
// Parameters with different location, but with similar names.
//
// GET /similarNames
func (s *Server) handleSimilarNamesRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("similarNames"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/similarNames"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SimilarNames",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SimilarNames",
			ID:   "similarNames",
		}
	)
	params, err := decodeSimilarNamesParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *SimilarNamesOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "SimilarNames",
			OperationSummary: "parameters with different location, but with similar names",
			OperationID:      "similarNames",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "x-param",
					In:   "query",
				}: params.QueryXParam,
				{
					Name: "X-Param",
					In:   "header",
				}: params.HeaderXParam,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SimilarNamesParams
			Response = *SimilarNamesOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSimilarNamesParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.SimilarNames(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.SimilarNames(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSimilarNamesResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
