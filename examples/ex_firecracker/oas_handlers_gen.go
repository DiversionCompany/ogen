// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/errors"
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
	_ = chi.Context{}
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
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
)

func (s *HTTPServer) HandleCreateSnapshotRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `CreateSnapshot`,
		trace.WithAttributes(otelogen.OperationID(`createSnapshot`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreateSnapshotRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.CreateSnapshot(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateSnapshotResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleCreateSyncActionRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `CreateSyncAction`,
		trace.WithAttributes(otelogen.OperationID(`createSyncAction`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreateSyncActionRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.CreateSyncAction(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateSyncActionResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleDescribeBalloonConfigRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `DescribeBalloonConfig`,
		trace.WithAttributes(otelogen.OperationID(`describeBalloonConfig`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.s.DescribeBalloonConfig(ctx)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDescribeBalloonConfigResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleDescribeBalloonStatsRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `DescribeBalloonStats`,
		trace.WithAttributes(otelogen.OperationID(`describeBalloonStats`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.s.DescribeBalloonStats(ctx)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDescribeBalloonStatsResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleDescribeInstanceRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `DescribeInstance`,
		trace.WithAttributes(otelogen.OperationID(`describeInstance`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.s.DescribeInstance(ctx)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDescribeInstanceResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleGetExportVmConfigRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `GetExportVmConfig`,
		trace.WithAttributes(otelogen.OperationID(`getExportVmConfig`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.s.GetExportVmConfig(ctx)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeGetExportVmConfigResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleGetMachineConfigurationRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `GetMachineConfiguration`,
		trace.WithAttributes(otelogen.OperationID(`getMachineConfiguration`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.s.GetMachineConfiguration(ctx)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeGetMachineConfigurationResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleLoadSnapshotRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `LoadSnapshot`,
		trace.WithAttributes(otelogen.OperationID(`loadSnapshot`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeLoadSnapshotRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.LoadSnapshot(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeLoadSnapshotResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleMmdsConfigPutRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `MmdsConfigPut`,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeMmdsConfigPutRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.MmdsConfigPut(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeMmdsConfigPutResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleMmdsGetRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `MmdsGet`,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.s.MmdsGet(ctx)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeMmdsGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleMmdsPatchRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `MmdsPatch`,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeMmdsPatchRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.MmdsPatch(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeMmdsPatchResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandleMmdsPutRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `MmdsPut`,
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeMmdsPutRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.MmdsPut(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeMmdsPutResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePatchBalloonRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PatchBalloon`,
		trace.WithAttributes(otelogen.OperationID(`patchBalloon`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePatchBalloonRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PatchBalloon(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchBalloonResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePatchBalloonStatsIntervalRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PatchBalloonStatsInterval`,
		trace.WithAttributes(otelogen.OperationID(`patchBalloonStatsInterval`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePatchBalloonStatsIntervalRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PatchBalloonStatsInterval(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchBalloonStatsIntervalResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePatchGuestDriveByIDRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PatchGuestDriveByID`,
		trace.WithAttributes(otelogen.OperationID(`patchGuestDriveByID`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePatchGuestDriveByIDParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodePatchGuestDriveByIDRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PatchGuestDriveByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchGuestDriveByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePatchGuestNetworkInterfaceByIDRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PatchGuestNetworkInterfaceByID`,
		trace.WithAttributes(otelogen.OperationID(`patchGuestNetworkInterfaceByID`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePatchGuestNetworkInterfaceByIDParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodePatchGuestNetworkInterfaceByIDRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PatchGuestNetworkInterfaceByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchGuestNetworkInterfaceByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePatchMachineConfigurationRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PatchMachineConfiguration`,
		trace.WithAttributes(otelogen.OperationID(`patchMachineConfiguration`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePatchMachineConfigurationRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PatchMachineConfiguration(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchMachineConfigurationResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePatchVmRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PatchVm`,
		trace.WithAttributes(otelogen.OperationID(`patchVm`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePatchVmRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PatchVm(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePatchVmResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePutBalloonRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PutBalloon`,
		trace.WithAttributes(otelogen.OperationID(`putBalloon`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutBalloonRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PutBalloon(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutBalloonResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePutGuestBootSourceRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PutGuestBootSource`,
		trace.WithAttributes(otelogen.OperationID(`putGuestBootSource`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutGuestBootSourceRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PutGuestBootSource(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutGuestBootSourceResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePutGuestDriveByIDRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PutGuestDriveByID`,
		trace.WithAttributes(otelogen.OperationID(`putGuestDriveByID`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePutGuestDriveByIDParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodePutGuestDriveByIDRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PutGuestDriveByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutGuestDriveByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePutGuestNetworkInterfaceByIDRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PutGuestNetworkInterfaceByID`,
		trace.WithAttributes(otelogen.OperationID(`putGuestNetworkInterfaceByID`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodePutGuestNetworkInterfaceByIDParams(r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodePutGuestNetworkInterfaceByIDRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PutGuestNetworkInterfaceByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutGuestNetworkInterfaceByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePutGuestVsockRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PutGuestVsock`,
		trace.WithAttributes(otelogen.OperationID(`putGuestVsock`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutGuestVsockRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PutGuestVsock(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutGuestVsockResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePutLoggerRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PutLogger`,
		trace.WithAttributes(otelogen.OperationID(`putLogger`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutLoggerRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PutLogger(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutLoggerResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePutMachineConfigurationRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PutMachineConfiguration`,
		trace.WithAttributes(otelogen.OperationID(`putMachineConfiguration`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutMachineConfigurationRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PutMachineConfiguration(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutMachineConfigurationResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func (s *HTTPServer) HandlePutMetricsRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `PutMetrics`,
		trace.WithAttributes(otelogen.OperationID(`putMetrics`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodePutMetricsRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.s.PutMetrics(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodePutMetricsResponse(response, w, span); err != nil {
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
