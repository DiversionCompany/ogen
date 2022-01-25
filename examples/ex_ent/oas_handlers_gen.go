// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
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
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

// HandleCreatePetRequest handles createPet operation.
//
// POST /pets
func (s *Server) handleCreatePetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `CreatePet`,
		trace.WithAttributes(otelogen.OperationID(`createPet`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreatePetRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreatePet(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreatePetCategoriesRequest handles createPetCategories operation.
//
// POST /pets/{id}/categories
func (s *Server) handleCreatePetCategoriesRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `CreatePetCategories`,
		trace.WithAttributes(otelogen.OperationID(`createPetCategories`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCreatePetCategoriesParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeCreatePetCategoriesRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreatePetCategories(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetCategoriesResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreatePetFriendsRequest handles createPetFriends operation.
//
// POST /pets/{id}/friends
func (s *Server) handleCreatePetFriendsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `CreatePetFriends`,
		trace.WithAttributes(otelogen.OperationID(`createPetFriends`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCreatePetFriendsParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeCreatePetFriendsRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreatePetFriends(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetFriendsResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreatePetOwnerRequest handles createPetOwner operation.
//
// POST /pets/{id}/owner
func (s *Server) handleCreatePetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `CreatePetOwner`,
		trace.WithAttributes(otelogen.OperationID(`createPetOwner`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCreatePetOwnerParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeCreatePetOwnerRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreatePetOwner(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleDeletePetRequest handles deletePet operation.
//
// DELETE /pets/{id}
func (s *Server) handleDeletePetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `DeletePet`,
		trace.WithAttributes(otelogen.OperationID(`deletePet`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeletePetParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeletePet(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeletePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleDeletePetOwnerRequest handles deletePetOwner operation.
//
// DELETE /pets/{id}/owner
func (s *Server) handleDeletePetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `DeletePetOwner`,
		trace.WithAttributes(otelogen.OperationID(`deletePetOwner`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeletePetOwnerParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeletePetOwner(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeletePetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListPetRequest handles listPet operation.
//
// GET /pets
func (s *Server) handleListPetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `ListPet`,
		trace.WithAttributes(otelogen.OperationID(`listPet`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListPetParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListPet(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListPetCategoriesRequest handles listPetCategories operation.
//
// GET /pets/{id}/categories
func (s *Server) handleListPetCategoriesRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `ListPetCategories`,
		trace.WithAttributes(otelogen.OperationID(`listPetCategories`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListPetCategoriesParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListPetCategories(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetCategoriesResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListPetFriendsRequest handles listPetFriends operation.
//
// GET /pets/{id}/friends
func (s *Server) handleListPetFriendsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `ListPetFriends`,
		trace.WithAttributes(otelogen.OperationID(`listPetFriends`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListPetFriendsParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListPetFriends(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetFriendsResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleReadPetRequest handles readPet operation.
//
// GET /pets/{id}
func (s *Server) handleReadPetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `ReadPet`,
		trace.WithAttributes(otelogen.OperationID(`readPet`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadPetParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadPet(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadPetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleReadPetOwnerRequest handles readPetOwner operation.
//
// GET /pets/{id}/owner
func (s *Server) handleReadPetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `ReadPetOwner`,
		trace.WithAttributes(otelogen.OperationID(`readPetOwner`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadPetOwnerParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadPetOwner(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadPetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleUpdatePetRequest handles updatePet operation.
//
// PATCH /pets/{id}
func (s *Server) handleUpdatePetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `UpdatePet`,
		trace.WithAttributes(otelogen.OperationID(`updatePet`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeUpdatePetParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeUpdatePetRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.UpdatePet(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdatePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
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
