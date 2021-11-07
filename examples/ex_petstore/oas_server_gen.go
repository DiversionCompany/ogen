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

// Server handles operations described by OpenAPI v3 specification.
type Server interface {
	// CreatePets implements createPets operation.
	CreatePets(ctx context.Context) (CreatePetsRes, error)
	// ListPets implements listPets operation.
	ListPets(ctx context.Context, params ListPetsParams) (ListPetsRes, error)
	// ShowPetById implements showPetById operation.
	ShowPetById(ctx context.Context, params ShowPetByIdParams) (ShowPetByIdRes, error)
}

type HTTPServer struct {
	s   Server
	mux *chi.Mux
	cfg config
}

func NewServer(s Server, opts ...Option) *HTTPServer {
	srv := &HTTPServer{
		s:   s,
		mux: chi.NewMux(),
		cfg: newConfig(opts...),
	}
	srv.setupRoutes()
	return srv
}

func (s *HTTPServer) setupRoutes() {
	s.mux.MethodFunc("POST", "/pets", s.HandleCreatePetsRequest)
	s.mux.MethodFunc("GET", "/pets", s.HandleListPetsRequest)
	s.mux.MethodFunc("GET", "/pets/{petId}", s.HandleShowPetByIdRequest)
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
