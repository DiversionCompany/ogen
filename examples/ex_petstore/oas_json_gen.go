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

// Encode implements json.Marshaler.
func (s CreatePetsCreated) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes CreatePetsCreated from json.
func (s *CreatePetsCreated) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode CreatePetsCreated to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s Error) Encode(e *jx.Encoder) {
	e.ObjStart()

	e.FieldStart("code")
	e.Int32(s.Code)

	e.FieldStart("message")
	e.Str(s.Message)
	e.ObjEnd()
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode Error to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			v, err := d.Int32()
			s.Code = int32(v)
			if err != nil {
				return err
			}
		case "message":
			v, err := d.Str()
			s.Message = string(v)
			if err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s ErrorStatusCode) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes ErrorStatusCode from json.
func (s *ErrorStatusCode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode ErrorStatusCode to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes int32 as json.
func (o OptInt32) Encode(e *jx.Encoder) {
	e.Int32(int32(o.Value))
}

// Decode decodes int32 from json.
func (o *OptInt32) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptInt32 to nil`)
	}
	switch d.Next() {
	case jx.Number:
		o.Set = true
		v, err := d.Int32()
		if err != nil {
			return err
		}
		o.Value = int32(v)
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptInt32`, d.Next())
	}
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptString to nil`)
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		v, err := d.Str()
		if err != nil {
			return err
		}
		o.Value = string(v)
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptString`, d.Next())
	}
}

// Encode implements json.Marshaler.
func (s Pet) Encode(e *jx.Encoder) {
	e.ObjStart()

	e.FieldStart("id")
	e.Int64(s.ID)

	e.FieldStart("name")
	e.Str(s.Name)
	if s.Tag.Set {
		e.FieldStart("tag")
		s.Tag.Encode(e)
	}
	e.ObjEnd()
}

// Decode decodes Pet from json.
func (s *Pet) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode Pet to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			v, err := d.Int64()
			s.ID = int64(v)
			if err != nil {
				return err
			}
		case "name":
			v, err := d.Str()
			s.Name = string(v)
			if err != nil {
				return err
			}
		case "tag":
			s.Tag.Reset()
			if err := s.Tag.Decode(d); err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes Pets as json.
func (s Pets) Encode(e *jx.Encoder) {
	unwrapped := []Pet(s)
	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes Pets from json.
func (s *Pets) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode Pets to nil`)
	}
	var unwrapped []Pet
	if err := func() error {
		unwrapped = nil
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem Pet
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
	*s = Pets(unwrapped)
	return nil
}
