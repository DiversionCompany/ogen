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
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// HandleCreatePetsRequest handles createPets operation.
//
// POST /pets
func (s *Server) handleCreatePetsRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createPets"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePets",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.CreatePets(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleListPetsRequest handles listPets operation.
//
// GET /pets
func (s *Server) handleListPetsRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listPets"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPets",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeListPetsParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, span, otelAttrs, err)
		return
	}

	response, err := s.h.ListPets(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleShowPetByIdRequest handles showPetById operation.
//
// GET /pets/{petId}
func (s *Server) handleShowPetByIdRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("showPetById"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ShowPetById",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodeShowPetByIdParams(args, r)
	if err != nil {
		s.badRequest(ctx, w, span, otelAttrs, err)
		return
	}

	response, err := s.h.ShowPetById(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeShowPetByIdResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

func (s *Server) badRequest(ctx context.Context, w http.ResponseWriter, span trace.Span, otelAttrs []attribute.KeyValue, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, "BadRequest")
	s.errors.Add(ctx, 1, otelAttrs...)
	respondError(w, http.StatusBadRequest, err)
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}
