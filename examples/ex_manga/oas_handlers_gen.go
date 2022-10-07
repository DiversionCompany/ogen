// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// Allocate option closure once.
var serverSpanKind = trace.WithSpanKind(trace.SpanKindServer)

// handleGetBookRequest handles getBook operation.
//
// GET /api/gallery/{book_id}
func (s *Server) handleGetBookRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBook"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetBook",
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
			Name: "GetBook",
			ID:   "getBook",
		}
	)
	params, err := decodeGetBookParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetBookRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "GetBook",
			OperationID:   "getBook",
			Body:          nil,
			Params: map[string]any{
				"book_id": params.BookID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetBookParams
			Response = GetBookRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetBookParams,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.GetBook(ctx, params)
			},
		)
	} else {
		response, err = s.h.GetBook(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetBookResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleGetPageCoverImageRequest handles getPageCoverImage operation.
//
// GET /galleries/{media_id}/cover.{format}
func (s *Server) handleGetPageCoverImageRequest(args [2]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPageCoverImage"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetPageCoverImage",
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
			Name: "GetPageCoverImage",
			ID:   "getPageCoverImage",
		}
	)
	params, err := decodeGetPageCoverImageParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetPageCoverImageRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "GetPageCoverImage",
			OperationID:   "getPageCoverImage",
			Body:          nil,
			Params: map[string]any{
				"media_id": params.MediaID,
				"format":   params.Format,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPageCoverImageParams
			Response = GetPageCoverImageRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPageCoverImageParams,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.GetPageCoverImage(ctx, params)
			},
		)
	} else {
		response, err = s.h.GetPageCoverImage(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPageCoverImageResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleGetPageImageRequest handles getPageImage operation.
//
// GET /galleries/{media_id}/{page}.{format}
func (s *Server) handleGetPageImageRequest(args [3]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPageImage"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetPageImage",
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
			Name: "GetPageImage",
			ID:   "getPageImage",
		}
	)
	params, err := decodeGetPageImageParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetPageImageRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "GetPageImage",
			OperationID:   "getPageImage",
			Body:          nil,
			Params: map[string]any{
				"media_id": params.MediaID,
				"page":     params.Page,
				"format":   params.Format,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPageImageParams
			Response = GetPageImageRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPageImageParams,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.GetPageImage(ctx, params)
			},
		)
	} else {
		response, err = s.h.GetPageImage(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPageImageResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleGetPageThumbnailImageRequest handles getPageThumbnailImage operation.
//
// GET /galleries/{media_id}/{page}t.{format}
func (s *Server) handleGetPageThumbnailImageRequest(args [3]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPageThumbnailImage"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetPageThumbnailImage",
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
			Name: "GetPageThumbnailImage",
			ID:   "getPageThumbnailImage",
		}
	)
	params, err := decodeGetPageThumbnailImageParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetPageThumbnailImageRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "GetPageThumbnailImage",
			OperationID:   "getPageThumbnailImage",
			Body:          nil,
			Params: map[string]any{
				"media_id": params.MediaID,
				"page":     params.Page,
				"format":   params.Format,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetPageThumbnailImageParams
			Response = GetPageThumbnailImageRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetPageThumbnailImageParams,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.GetPageThumbnailImage(ctx, params)
			},
		)
	} else {
		response, err = s.h.GetPageThumbnailImage(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetPageThumbnailImageResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSearchRequest handles search operation.
//
// GET /api/galleries/search
func (s *Server) handleSearchRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("search"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Search",
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
			Name: "Search",
			ID:   "search",
		}
	)
	params, err := decodeSearchParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response SearchRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "Search",
			OperationID:   "search",
			Body:          nil,
			Params: map[string]any{
				"query": params.Query,
				"page":  params.Page,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SearchParams
			Response = SearchRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSearchParams,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.Search(ctx, params)
			},
		)
	} else {
		response, err = s.h.Search(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSearchResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSearchByTagIDRequest handles searchByTagID operation.
//
// GET /api/galleries/tagged
func (s *Server) handleSearchByTagIDRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("searchByTagID"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SearchByTagID",
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
			Name: "SearchByTagID",
			ID:   "searchByTagID",
		}
	)
	params, err := decodeSearchByTagIDParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response SearchByTagIDRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "SearchByTagID",
			OperationID:   "searchByTagID",
			Body:          nil,
			Params: map[string]any{
				"tag_id": params.TagID,
				"page":   params.Page,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SearchByTagIDParams
			Response = SearchByTagIDRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSearchByTagIDParams,
			func(ctx context.Context, request Request, params Params) (Response, error) {
				return s.h.SearchByTagID(ctx, params)
			},
		)
	} else {
		response, err = s.h.SearchByTagID(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSearchByTagIDResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
