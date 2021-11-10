// Code generated by ogen, DO NOT EDIT.

package techempower

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
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
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
	_ = jx.Null
	_ = sync.Pool{}
)

func decodeCachingResponse(resp *http.Response, span trace.Span) (res WorldObjects, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response WorldObjects
			if err := func() error {
				{
					var unwrapped []WorldObject
					unwrapped = nil
					if err := d.Arr(func(d *jx.Decoder) error {
						var elem WorldObject
						if err := elem.Decode(d); err != nil {
							return err
						}
						unwrapped = append(unwrapped, elem)
						return nil
					}); err != nil {
						return err
					}
					response = WorldObjects(unwrapped)
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeDBResponse(resp *http.Response, span trace.Span) (res WorldObject, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response WorldObject
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeJSONResponse(resp *http.Response, span trace.Span) (res HelloWorld, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response HelloWorld
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeQueriesResponse(resp *http.Response, span trace.Span) (res WorldObjects, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response WorldObjects
			if err := func() error {
				{
					var unwrapped []WorldObject
					unwrapped = nil
					if err := d.Arr(func(d *jx.Decoder) error {
						var elem WorldObject
						if err := elem.Decode(d); err != nil {
							return err
						}
						unwrapped = append(unwrapped, elem)
						return nil
					}); err != nil {
						return err
					}
					response = WorldObjects(unwrapped)
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}

func decodeUpdatesResponse(resp *http.Response, span trace.Span) (res WorldObjects, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response WorldObjects
			if err := func() error {
				{
					var unwrapped []WorldObject
					unwrapped = nil
					if err := d.Arr(func(d *jx.Decoder) error {
						var elem WorldObject
						if err := elem.Decode(d); err != nil {
							return err
						}
						unwrapped = append(unwrapped, elem)
						return nil
					}); err != nil {
						return err
					}
					response = WorldObjects(unwrapped)
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return response, nil
		default:
			return res, errors.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		return res, errors.Errorf("unexpected statusCode: %d", resp.StatusCode)
	}
}
