// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
)

// WriteJSON implements json.Marshaler.
func (s Error) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("code")
	j.WriteInt64(s.Code)
	field.Write("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes Error json value to io.Writer.
func (s Error) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Error json value from io.Reader.
func (s *Error) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Error from json stream.
func (s *Error) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "code":
			s.Code = i.ReadInt64()
			return i.Error == nil
		case "message":
			s.Message = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s ErrorStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes ErrorStatusCode json value to io.Writer.
func (s ErrorStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads ErrorStatusCode json value from io.Reader.
func (s *ErrorStatusCode) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads ErrorStatusCode from json stream.
func (s *ErrorStatusCode) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPutDefault) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes FoobarPutDefault json value to io.Writer.
func (s FoobarPutDefault) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads FoobarPutDefault json value from io.Reader.
func (s *FoobarPutDefault) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads FoobarPutDefault from json stream.
func (s *FoobarPutDefault) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON writes json value of PetType to json stream.
func (o OptionalPetType) WriteJSON(j *json.Stream) {
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of PetType from json iterator.
func (o *OptionalPetType) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Value = PetType(i.ReadString())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalPetType", i.WhatIsNext())
	}
	return nil
}

// WriteJSON implements json.Marshaler.
func (s Pet) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("birthday")
	json.WriteDate(j, s.Birthday)
	// Unsupported kind "pointer" for field "Friends".
	field.Write("id")
	j.WriteInt64(s.ID)
	field.Write("ip")
	json.WriteIP(j, s.IP)
	field.Write("ip_v4")
	json.WriteIP(j, s.IPV4)
	field.Write("ip_v6")
	json.WriteIP(j, s.IPV6)
	field.Write("name")
	j.WriteString(s.Name)
	field.Write("nickname")
	s.Nickname.WriteJSON(j)
	if s.NullStr.Set {
		field.Write("nullStr")
		s.NullStr.WriteJSON(j)
	}
	field.Write("rate")
	json.WriteDuration(j, s.Rate)
	if s.Tag.Set {
		field.Write("tag")
		s.Tag.WriteJSON(j)
	}
	// Unsupported kind "pointer" for field "TestArray1".
	if s.TestDate.Set {
		field.Write("testDate")
		s.TestDate.WriteJSON(j, json.WriteDate)
	}
	if s.TestDateTime.Set {
		field.Write("testDateTime")
		s.TestDateTime.WriteJSON(j, json.WriteDateTime)
	}
	if s.TestDuration.Set {
		field.Write("testDuration")
		s.TestDuration.WriteJSON(j)
	}
	if s.TestFloat1.Set {
		field.Write("testFloat1")
		s.TestFloat1.WriteJSON(j)
	}
	if s.TestInteger1.Set {
		field.Write("testInteger1")
		s.TestInteger1.WriteJSON(j)
	}
	if s.TestTime.Set {
		field.Write("testTime")
		s.TestTime.WriteJSON(j, json.WriteTime)
	}
	// Unsupported kind "generic" for field "Type".
	field.Write("uri")
	json.WriteURI(j, s.URI)
	field.Write("unique_id")
	json.WriteUUID(j, s.UniqueID)
	j.WriteObjectEnd()
}

// WriteJSONTo writes Pet json value to io.Writer.
func (s Pet) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Pet json value from io.Reader.
func (s *Pet) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Pet from json stream.
func (s *Pet) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "birthday":
			v, err := json.ReadDate(i)
			if err != nil {
				i.ReportError("Field Birthday", err.Error())
				return false
			}
			s.Birthday = v
			return true
		case "friends":
			// Unsupported kind "pointer" for field "Friends".
			i.Skip()
			return true
		case "id":
			s.ID = i.ReadInt64()
			return i.Error == nil
		case "ip":
			v, err := json.ReadIP(i)
			if err != nil {
				i.ReportError("Field IP", err.Error())
				return false
			}
			s.IP = v
			return true
		case "ip_v4":
			v, err := json.ReadIP(i)
			if err != nil {
				i.ReportError("Field IPV4", err.Error())
				return false
			}
			s.IPV4 = v
			return true
		case "ip_v6":
			v, err := json.ReadIP(i)
			if err != nil {
				i.ReportError("Field IPV6", err.Error())
				return false
			}
			s.IPV6 = v
			return true
		case "name":
			s.Name = i.ReadString()
			return i.Error == nil
		case "nickname":
			if err := s.Nickname.ReadJSON(i); err != nil {
				i.ReportError("Field Nickname", err.Error())
				return false
			}
			return true
		case "nullStr":
			s.NullStr.Reset()
			if err := s.NullStr.ReadJSON(i); err != nil {
				i.ReportError("Field NullStr", err.Error())
				return false
			}
			return true
		case "rate":
			v, err := json.ReadDuration(i)
			if err != nil {
				i.ReportError("Field Rate", err.Error())
				return false
			}
			s.Rate = v
			return true
		case "tag":
			s.Tag.Reset()
			if err := s.Tag.ReadJSON(i); err != nil {
				i.ReportError("Field Tag", err.Error())
				return false
			}
			return true
		case "testArray1":
			// Unsupported kind "pointer" for field "TestArray1".
			i.Skip()
			return true
		case "testDate":
			s.TestDate.Reset()
			if err := s.TestDate.ReadJSON(i, json.ReadDate); err != nil {
				i.ReportError("Field TestDate", err.Error())
				return false
			}
			return true
		case "testDateTime":
			s.TestDateTime.Reset()
			if err := s.TestDateTime.ReadJSON(i, json.ReadDateTime); err != nil {
				i.ReportError("Field TestDateTime", err.Error())
				return false
			}
			return true
		case "testDuration":
			s.TestDuration.Reset()
			if err := s.TestDuration.ReadJSON(i); err != nil {
				i.ReportError("Field TestDuration", err.Error())
				return false
			}
			return true
		case "testFloat1":
			s.TestFloat1.Reset()
			if err := s.TestFloat1.ReadJSON(i); err != nil {
				i.ReportError("Field TestFloat1", err.Error())
				return false
			}
			return true
		case "testInteger1":
			s.TestInteger1.Reset()
			if err := s.TestInteger1.ReadJSON(i); err != nil {
				i.ReportError("Field TestInteger1", err.Error())
				return false
			}
			return true
		case "testTime":
			s.TestTime.Reset()
			if err := s.TestTime.ReadJSON(i, json.ReadTime); err != nil {
				i.ReportError("Field TestTime", err.Error())
				return false
			}
			return true
		case "type":
			// Unsupported kind "generic" for field "Type".
			i.Skip()
			return true
		case "uri":
			v, err := json.ReadURI(i)
			if err != nil {
				i.ReportError("Field URI", err.Error())
				return false
			}
			s.URI = v
			return true
		case "unique_id":
			v, err := json.ReadUUID(i)
			if err != nil {
				i.ReportError("Field UniqueID", err.Error())
				return false
			}
			s.UniqueID = v
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetDefault) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetDefault json value to io.Writer.
func (s PetGetDefault) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PetGetDefault json value from io.Reader.
func (s *PetGetDefault) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PetGetDefault from json stream.
func (s *PetGetDefault) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "message":
			s.Message = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetDefaultStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetDefaultStatusCode json value to io.Writer.
func (s PetGetDefaultStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PetGetDefaultStatusCode json value from io.Reader.
func (s *PetGetDefaultStatusCode) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PetGetDefaultStatusCode from json stream.
func (s *PetGetDefaultStatusCode) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}
