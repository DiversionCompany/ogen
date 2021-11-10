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

// Encode implements json.Marshaler.
func (s Data) Encode(e *jx.Encoder) {
	e.ObjStart()

	e.FieldStart("id")
	s.ID.Encode(e)

	e.FieldStart("description")
	s.Description.Encode(e)

	e.FieldStart("email")
	e.Str(s.Email)

	e.FieldStart("hostname")
	e.Str(s.Hostname)

	e.FieldStart("format")
	e.Str(s.Format)

	e.FieldStart("base64")
	e.Base64(s.Base64)
	e.ObjEnd()
}

// Decode decodes Data from json.
func (s *Data) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode Data to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := s.ID.Decode(d); err != nil {
				return err
			}
		case "description":
			if err := s.Description.Decode(d); err != nil {
				return err
			}
		case "email":
			v, err := d.Str()
			s.Email = string(v)
			if err != nil {
				return err
			}
		case "hostname":
			v, err := d.Str()
			s.Hostname = string(v)
			if err != nil {
				return err
			}
		case "format":
			v, err := d.Str()
			s.Format = string(v)
			if err != nil {
				return err
			}
		case "base64":
			v, err := d.Base64()
			s.Base64 = []byte(v)
			if err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes DataDescription as json.
func (s DataDescription) Encode(e *jx.Encoder) {
	switch s.Type {
	case DescriptionDetailedDataDescription:
		s.DescriptionDetailed.Encode(e)
	case DescriptionSimpleDataDescription:
		s.DescriptionSimple.Encode(e)
	}
}

// Decode decodes DataDescription from json.
func (s *DataDescription) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode DataDescription to nil`)
	}
	// Sum type fields.
	if d.Next() != jx.Object {
		return errors.Errorf("unexpected json type %q", d.Next())
	}
	var found bool
	if err := d.Capture(func(d *jx.Decoder) error {
		return d.ObjBytes(func(d *jx.Decoder, key []byte) error {
			if found {
				return d.Skip()
			}
			switch string(key) {
			case "name":
				found = true
				s.Type = DescriptionDetailedDataDescription
			case "count":
				found = true
				s.Type = DescriptionDetailedDataDescription
			case "description":
				found = true
				s.Type = DescriptionSimpleDataDescription
			}
			return d.Skip()
		})
	}); err != nil {
		return errors.Wrap(err, "capture")
	}
	if !found {
		return errors.New("unable to detect sum type variant")
	}
	switch s.Type {
	case DescriptionDetailedDataDescription:
		if err := s.DescriptionDetailed.Decode(d); err != nil {
			return err
		}
	case DescriptionSimpleDataDescription:
		if err := s.DescriptionSimple.Decode(d); err != nil {
			return err
		}
	default:
		return errors.Errorf("inferred invalid type: %s", s.Type)
	}
	return nil
}

// Encode implements json.Marshaler.
func (s DescriptionDetailed) Encode(e *jx.Encoder) {
	e.ObjStart()

	e.FieldStart("name")
	e.Str(s.Name)

	e.FieldStart("count")
	e.Int(s.Count)
	e.ObjEnd()
}

// Decode decodes DescriptionDetailed from json.
func (s *DescriptionDetailed) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode DescriptionDetailed to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			v, err := d.Str()
			s.Name = string(v)
			if err != nil {
				return err
			}
		case "count":
			v, err := d.Int()
			s.Count = int(v)
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
func (s DescriptionSimple) Encode(e *jx.Encoder) {
	e.ObjStart()

	e.FieldStart("description")
	e.Str(s.Description)
	e.ObjEnd()
}

// Decode decodes DescriptionSimple from json.
func (s *DescriptionSimple) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode DescriptionSimple to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "description":
			v, err := d.Str()
			s.Description = string(v)
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
func (s Error) Encode(e *jx.Encoder) {
	e.ObjStart()

	e.FieldStart("code")
	e.Int64(s.Code)

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
			v, err := d.Int64()
			s.Code = int64(v)
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

// Encode implements json.Marshaler.
func (s FoobarPutDef) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes FoobarPutDef from json.
func (s *FoobarPutDef) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode FoobarPutDef to nil`)
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
func (s FoobarPutDefStatusCode) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes FoobarPutDefStatusCode from json.
func (s *FoobarPutDefStatusCode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode FoobarPutDefStatusCode to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes ID as json.
func (s ID) Encode(e *jx.Encoder) {
	switch s.Type {
	case StringID:
		e.Str(s.String)
	case IntID:
		e.Int(s.Int)
	}
}

// Decode decodes ID from json.
func (s *ID) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode ID to nil`)
	}
	// Sum type primitive.
	switch t := d.Next(); t {
	case jx.String:
		v, err := d.Str()
		s.String = string(v)
		if err != nil {
			return err
		}
		s.Type = StringID
	case jx.Number:
		v, err := d.Int()
		s.Int = int(v)
		if err != nil {
			return err
		}
		s.Type = IntID
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// Encode encodes string as json.
func (o NilString) Encode(e *jx.Encoder) {
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *NilString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode NilString to nil`)
	}
	switch d.Next() {
	case jx.String:
		o.Null = false
		v, err := d.Str()
		if err != nil {
			return err
		}
		o.Value = string(v)
		return nil
	case jx.Null:
		if err := d.Null(); err != nil {
			return err
		}
		var v string
		o.Value = v
		o.Null = true
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading NilString`, d.Next())
	}
}

// Encode implements json.Marshaler.
func (s NotFound) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes NotFound from json.
func (s *NotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode NotFound to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes Data as json.
func (o OptData) Encode(e *jx.Encoder) {
	o.Value.Encode(e)
}

// Decode decodes Data from json.
func (o *OptData) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptData to nil`)
	}
	switch d.Next() {
	case jx.Object:
		o.Set = true
		if err := o.Value.Decode(d); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptData`, d.Next())
	}
}

// Encode encodes time.Duration as json.
func (o OptDuration) Encode(e *jx.Encoder) {
	json.EncodeDuration(e, o.Value)
}

// Decode decodes time.Duration from json.
func (o *OptDuration) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptDuration to nil`)
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		v, err := json.DecodeDuration(d)
		if err != nil {
			return err
		}
		o.Value = v
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptDuration`, d.Next())
	}
}

// Encode encodes float64 as json.
func (o OptFloat64) Encode(e *jx.Encoder) {
	e.Float64(float64(o.Value))
}

// Decode decodes float64 from json.
func (o *OptFloat64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptFloat64 to nil`)
	}
	switch d.Next() {
	case jx.Number:
		o.Set = true
		v, err := d.Float64()
		if err != nil {
			return err
		}
		o.Value = float64(v)
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptFloat64`, d.Next())
	}
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptInt to nil`)
	}
	switch d.Next() {
	case jx.Number:
		o.Set = true
		v, err := d.Int()
		if err != nil {
			return err
		}
		o.Value = int(v)
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptInt`, d.Next())
	}
}

// Encode encodes string as json.
func (o OptNilString) Encode(e *jx.Encoder) {
	if o.Null {
		e.Null()
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptNilString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptNilString to nil`)
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		o.Null = false
		v, err := d.Str()
		if err != nil {
			return err
		}
		o.Value = string(v)
		return nil
	case jx.Null:
		if err := d.Null(); err != nil {
			return err
		}
		var v string
		o.Value = v
		o.Set = true
		o.Null = true
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptNilString`, d.Next())
	}
}

// Encode encodes PetType as json.
func (o OptPetType) Encode(e *jx.Encoder) {
	e.Str(string(o.Value))
}

// Decode decodes PetType from json.
func (o *OptPetType) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptPetType to nil`)
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		v, err := d.Str()
		if err != nil {
			return err
		}
		o.Value = PetType(v)
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptPetType`, d.Next())
	}
}

// Encode encodes time.Time as json.
func (o OptTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptTime to nil`)
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		v, err := format(d)
		if err != nil {
			return err
		}
		o.Value = v
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptTime`, d.Next())
	}
}

// Encode encodes uuid.UUID as json.
func (o OptUUID) Encode(e *jx.Encoder) {
	json.EncodeUUID(e, o.Value)
}

// Decode decodes uuid.UUID from json.
func (o *OptUUID) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New(`invalid: unable to decode OptUUID to nil`)
	}
	switch d.Next() {
	case jx.String:
		o.Set = true
		v, err := json.DecodeUUID(d)
		if err != nil {
			return err
		}
		o.Value = v
		return nil
	default:
		return errors.Errorf(`unexpected type %q while reading OptUUID`, d.Next())
	}
}

// Encode implements json.Marshaler.
func (s Pet) Encode(e *jx.Encoder) {
	e.ObjStart()
	if s.Primary != nil {
		e.FieldStart("primary")
		s.Primary.Encode(e)
	}

	e.FieldStart("id")
	e.Int64(s.ID)

	e.FieldStart("unique_id")
	json.EncodeUUID(e, s.UniqueID)

	e.FieldStart("name")
	e.Str(s.Name)
	if s.Type.Set {
		e.FieldStart("type")
		s.Type.Encode(e)
	}

	e.FieldStart("kind")
	s.Kind.Encode(e)
	if s.Tag.Set {
		e.FieldStart("tag")
		s.Tag.Encode(e)
	}

	e.FieldStart("ip")
	json.EncodeIP(e, s.IP)

	e.FieldStart("ip_v4")
	json.EncodeIP(e, s.IPV4)

	e.FieldStart("ip_v6")
	json.EncodeIP(e, s.IPV6)

	e.FieldStart("uri")
	json.EncodeURI(e, s.URI)

	e.FieldStart("birthday")
	json.EncodeDate(e, s.Birthday)

	e.FieldStart("rate")
	json.EncodeDuration(e, s.Rate)

	e.FieldStart("nickname")
	s.Nickname.Encode(e)
	if s.NullStr.Set {
		e.FieldStart("nullStr")
		s.NullStr.Encode(e)
	}
	if s.Friends != nil {
		e.FieldStart("friends")
		e.ArrStart()
		for _, elem := range s.Friends {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	if s.Next.Set {
		e.FieldStart("next")
		s.Next.Encode(e)
	}
	if s.TestInteger1.Set {
		e.FieldStart("testInteger1")
		s.TestInteger1.Encode(e)
	}
	if s.TestFloat1.Set {
		e.FieldStart("testFloat1")
		s.TestFloat1.Encode(e)
	}
	if s.TestArray1 != nil {
		e.FieldStart("testArray1")
		e.ArrStart()
		for _, elem := range s.TestArray1 {
			e.ArrStart()
			for _, elem := range elem {
				e.Str(elem)
			}
			e.ArrEnd()
		}
		e.ArrEnd()
	}
	if s.TestDate.Set {
		e.FieldStart("testDate")
		s.TestDate.Encode(e, json.EncodeDate)
	}
	if s.TestDuration.Set {
		e.FieldStart("testDuration")
		s.TestDuration.Encode(e)
	}
	if s.TestTime.Set {
		e.FieldStart("testTime")
		s.TestTime.Encode(e, json.EncodeTime)
	}
	if s.TestDateTime.Set {
		e.FieldStart("testDateTime")
		s.TestDateTime.Encode(e, json.EncodeDateTime)
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
		case "primary":
			s.Primary = nil
			var elem Pet
			if err := elem.Decode(d); err != nil {
				return err
			}
			s.Primary = &elem
		case "id":
			v, err := d.Int64()
			s.ID = int64(v)
			if err != nil {
				return err
			}
		case "unique_id":
			v, err := json.DecodeUUID(d)
			s.UniqueID = v
			if err != nil {
				return err
			}
		case "name":
			v, err := d.Str()
			s.Name = string(v)
			if err != nil {
				return err
			}
		case "type":
			s.Type.Reset()
			if err := s.Type.Decode(d); err != nil {
				return err
			}
		case "kind":
			if err := s.Kind.Decode(d); err != nil {
				return err
			}
		case "tag":
			s.Tag.Reset()
			if err := s.Tag.Decode(d); err != nil {
				return err
			}
		case "ip":
			v, err := json.DecodeIP(d)
			s.IP = v
			if err != nil {
				return err
			}
		case "ip_v4":
			v, err := json.DecodeIP(d)
			s.IPV4 = v
			if err != nil {
				return err
			}
		case "ip_v6":
			v, err := json.DecodeIP(d)
			s.IPV6 = v
			if err != nil {
				return err
			}
		case "uri":
			v, err := json.DecodeURI(d)
			s.URI = v
			if err != nil {
				return err
			}
		case "birthday":
			v, err := json.DecodeDate(d)
			s.Birthday = v
			if err != nil {
				return err
			}
		case "rate":
			v, err := json.DecodeDuration(d)
			s.Rate = v
			if err != nil {
				return err
			}
		case "nickname":
			if err := s.Nickname.Decode(d); err != nil {
				return err
			}
		case "nullStr":
			s.NullStr.Reset()
			if err := s.NullStr.Decode(d); err != nil {
				return err
			}
		case "friends":
			s.Friends = nil
			if err := d.Arr(func(d *jx.Decoder) error {
				var elem Pet
				if err := elem.Decode(d); err != nil {
					return err
				}
				s.Friends = append(s.Friends, elem)
				return nil
			}); err != nil {
				return err
			}
		case "next":
			s.Next.Reset()
			if err := s.Next.Decode(d); err != nil {
				return err
			}
		case "testInteger1":
			s.TestInteger1.Reset()
			if err := s.TestInteger1.Decode(d); err != nil {
				return err
			}
		case "testFloat1":
			s.TestFloat1.Reset()
			if err := s.TestFloat1.Decode(d); err != nil {
				return err
			}
		case "testArray1":
			s.TestArray1 = nil
			if err := d.Arr(func(d *jx.Decoder) error {
				var elem []string
				elem = nil
				if err := d.Arr(func(d *jx.Decoder) error {
					var elemElem string
					v, err := d.Str()
					elemElem = string(v)
					if err != nil {
						return err
					}
					elem = append(elem, elemElem)
					return nil
				}); err != nil {
					return err
				}
				s.TestArray1 = append(s.TestArray1, elem)
				return nil
			}); err != nil {
				return err
			}
		case "testDate":
			s.TestDate.Reset()
			if err := s.TestDate.Decode(d, json.DecodeDate); err != nil {
				return err
			}
		case "testDuration":
			s.TestDuration.Reset()
			if err := s.TestDuration.Decode(d); err != nil {
				return err
			}
		case "testTime":
			s.TestTime.Reset()
			if err := s.TestTime.Decode(d, json.DecodeTime); err != nil {
				return err
			}
		case "testDateTime":
			s.TestDateTime.Reset()
			if err := s.TestDateTime.Decode(d, json.DecodeDateTime); err != nil {
				return err
			}
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode implements json.Marshaler.
func (s PetGetDef) Encode(e *jx.Encoder) {
	e.ObjStart()

	e.FieldStart("message")
	e.Str(s.Message)
	e.ObjEnd()
}

// Decode decodes PetGetDef from json.
func (s *PetGetDef) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetGetDef to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
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
func (s PetGetDefStatusCode) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes PetGetDefStatusCode from json.
func (s *PetGetDefStatusCode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetGetDefStatusCode to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}

// Encode encodes PetKind as json.
func (s PetKind) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes PetKind from json.
func (s *PetKind) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetKind to nil`)
	}
	v, err := d.Str()
	if err != nil {
		return err
	}
	*s = PetKind(v)
	return nil
}

// Encode encodes PetName as json.
func (s PetName) Encode(e *jx.Encoder) {
	unwrapped := string(s)
	e.Str(unwrapped)
}

// Decode decodes PetName from json.
func (s *PetName) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetName to nil`)
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = PetName(unwrapped)
	return nil
}

// Encode encodes PetType as json.
func (s PetType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes PetType from json.
func (s *PetType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetType to nil`)
	}
	v, err := d.Str()
	if err != nil {
		return err
	}
	*s = PetType(v)
	return nil
}

// Encode implements json.Marshaler.
func (s PetUpdateNameAliasPostDef) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes PetUpdateNameAliasPostDef from json.
func (s *PetUpdateNameAliasPostDef) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetUpdateNameAliasPostDef to nil`)
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
func (s PetUpdateNameAliasPostDefStatusCode) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes PetUpdateNameAliasPostDefStatusCode from json.
func (s *PetUpdateNameAliasPostDefStatusCode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetUpdateNameAliasPostDefStatusCode to nil`)
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
func (s PetUpdateNamePostDef) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes PetUpdateNamePostDef from json.
func (s *PetUpdateNamePostDef) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetUpdateNamePostDef to nil`)
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
func (s PetUpdateNamePostDefStatusCode) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes PetUpdateNamePostDefStatusCode from json.
func (s *PetUpdateNamePostDefStatusCode) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetUpdateNamePostDefStatusCode to nil`)
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
func (s PetUploadAvatarByIDOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	e.ObjEnd()
}

// Decode decodes PetUploadAvatarByIDOK from json.
func (s *PetUploadAvatarByIDOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New(`invalid: unable to decode PetUploadAvatarByIDOK to nil`)
	}
	return d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
		return nil
	})
}
