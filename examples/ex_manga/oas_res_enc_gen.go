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
	"net/netip"
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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/nonrecording"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = bytes.NewReader
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = io.Copy
	_ = math.Mod
	_ = big.Rat{}
	_ = bits.LeadingZeros64
	_ = net.IP{}
	_ = http.MethodGet
	_ = netip.Addr{}
	_ = url.URL{}
	_ = regexp.MustCompile
	_ = sort.Ints
	_ = strconv.ParseInt
	_ = strings.Builder{}
	_ = sync.Pool{}
	_ = time.Time{}

	_ = errors.Is
	_ = jx.Null
	_ = uuid.UUID{}
	_ = otel.GetTracerProvider
	_ = attribute.KeyValue{}
	_ = codes.Unset
	_ = metric.MeterConfig{}
	_ = syncint64.Counter(nil)
	_ = nonrecording.NewNoopMeterProvider
	_ = trace.TraceIDFromHex

	_ = conv.ToInt32
	_ = ht.NewRequest
	_ = json.Marshal
	_ = otelogen.Version
	_ = uri.PathEncoder{}
	_ = validate.Int{}
)

func encodeGetBookResponse(response GetBookRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Book:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *GetBookForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))
		return nil
	default:
		return errors.Errorf("/api/gallery/{book_id}"+`: unexpected response type: %T`, response)
	}
}

func encodeSearchResponse(response SearchRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SearchOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *SearchForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))
		return nil
	default:
		return errors.Errorf("/api/galleries/search"+`: unexpected response type: %T`, response)
	}
}

func encodeSearchByTagIDResponse(response SearchByTagIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *SearchByTagIDOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *SearchByTagIDForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))
		return nil
	default:
		return errors.Errorf("/api/galleries/tagged"+`: unexpected response type: %T`, response)
	}
}
