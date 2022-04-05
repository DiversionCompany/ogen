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

// setDefaults set default value of fields.
func (s *DefaultTest) setDefaults() {
	{
		val := string("required")

		s.Required = val
	}
	{
		val := string("str")

		s.Str.SetTo(val)
	}
	{
		s.NullStr.Null = true
	}
	{
		val := DefaultTestEnum("big")

		s.Enum.SetTo(val)
	}
	{
		val, _ := json.DecodeUUID(jx.DecodeStr("\"123e4567-e89b-12d3-a456-426614174000\""))

		s.UUID.SetTo(val)
	}
	{
		val, _ := json.DecodeIP(jx.DecodeStr("\"1.1.1.1\""))

		s.IP.SetTo(val)
	}
	{
		val, _ := json.DecodeIP(jx.DecodeStr("\"1.1.1.1\""))

		s.IPV4.SetTo(val)
	}
	{
		val, _ := json.DecodeIP(jx.DecodeStr("\"2001:db8:85a3::8a2e:370:7334\""))

		s.IPV6.SetTo(val)
	}
	{
		val, _ := json.DecodeURI(jx.DecodeStr("\"s3://foo/baz\""))

		s.URI.SetTo(val)
	}
	{
		val, _ := json.DecodeDate(jx.DecodeStr("\"2011-10-10\""))

		s.Birthday.SetTo(val)
	}
	{
		val, _ := json.DecodeDuration(jx.DecodeStr("\"5s\""))

		s.Rate.SetTo(val)
	}
	{
		val := string("foo@example.com")

		s.Email.SetTo(val)
	}
	{
		val := string("example.org")

		s.Hostname.SetTo(val)
	}
	{
		val := string("1-2")

		s.Format.SetTo(val)
	}
	{
		val, _ := jx.DecodeStr("\"aGVsbG8sIHdvcmxkIQ==\"").Base64()

		s.Base64 = val
	}
}
