// Code generated by ogen, DO NOT EDIT.

package techempower

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
	_ = codes.Unset
)

// Encode implements json.Marshaler.
func (s HelloWorld) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"message\"" + ":")
		e.Str(s.Message)
	}
	e.ObjEnd()
}

var jsonFieldsNameOfHelloWorld = [1]string{
	0: "message",
}

// Decode decodes HelloWorld from json.
func (s *HelloWorld) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode HelloWorld to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode HelloWorld")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfHelloWorld) {
					name = jsonFieldsNameOfHelloWorld[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode implements json.Marshaler.
func (s WorldObject) Encode(e *jx.Writer) {
	e.ObjStart()
	var (
		first = true
		_     = first
	)
	{
		if !first {
			e.Comma()
		}
		first = false

		e.RawStr("\"id\"" + ":")
		e.Int64(s.ID)
	}
	{
		e.Comma()

		e.RawStr("\"randomNumber\"" + ":")
		e.Int64(s.RandomNumber)
	}
	e.ObjEnd()
}

var jsonFieldsNameOfWorldObject = [2]string{
	0: "id",
	1: "randomNumber",
}

// Decode decodes WorldObject from json.
func (s *WorldObject) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode WorldObject to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0

			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "randomNumber":
			requiredBitSet[0] |= 1 << 1

			if err := func() error {
				v, err := d.Int64()
				s.RandomNumber = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"randomNumber\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode WorldObject")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfWorldObject) {
					name = jsonFieldsNameOfWorldObject[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// Encode encodes WorldObjects as json.
func (s WorldObjects) Encode(e *jx.Writer) {
	unwrapped := []WorldObject(s)
	e.ArrStart()
	if len(unwrapped) >= 1 {
		// Encode first element without comma.
		{
			elem := unwrapped[0]
			elem.Encode(e)
		}
		for _, elem := range unwrapped[1:] {
			e.Comma()
			elem.Encode(e)
		}
	}
	e.ArrEnd()
}

// Decode decodes WorldObjects from json.
func (s *WorldObjects) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode WorldObjects to nil")
	}
	var unwrapped []WorldObject
	if err := func() error {
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
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = WorldObjects(unwrapped)
	return nil
}
