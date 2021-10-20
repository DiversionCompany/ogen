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
func (s Data) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	if s.Description.Set {
		more.More()
		j.WriteObjectField("description")
		s.Description.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes Data json value to io.Writer.
func (s Data) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Data json value from io.Reader.
func (s *Data) ReadJSONFrom(r io.Reader) error {
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

// ReadJSON reads Data from json stream.
func (s *Data) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "description":
			s.Description.Reset()
			if err := s.Description.ReadJSON(i); err != nil {
				i.ReportError("Field Description", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Error) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	j.WriteObjectField("code")
	j.WriteInt64(s.Code)
	more.More()
	j.WriteObjectField("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes Error json value to io.Writer.
func (s Error) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
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
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes ErrorStatusCode json value to io.Writer.
func (s ErrorStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
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
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes FoobarPutDefault json value to io.Writer.
func (s FoobarPutDefault) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
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

// WriteJSON writes json value of string to json stream.
func (o NilString) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *NilString) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Null = false
		o.Value = string(i.ReadString())
		return i.Error
	case json.NilValue:
		var v string
		o.Value = v
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilString", i.WhatIsNext())
	}
	return nil
}

func (NotFound) WriteJSON(j *json.Stream)        {}
func (NotFound) ReadJSON(i *json.Iterator) error { return nil }
func (NotFound) ReadJSONFrom(r io.Reader) error  { return nil }
func (NotFound) WriteJSONTo(w io.Writer) error   { return nil }

// WriteJSON writes json value of Data to json stream.
func (o OptData) WriteJSON(j *json.Stream) {
	o.Value.WriteJSON(j)
}

// ReadJSON reads json value of Data from json iterator.
func (o *OptData) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.ObjectValue:
		o.Set = true
		if err := o.Value.ReadJSON(i); err != nil {
			return err
		}
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptData", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of time.Duration to json stream.
func (o OptDuration) WriteJSON(j *json.Stream) {
	json.WriteDuration(j, o.Value)
}

// ReadJSON reads json value of time.Duration from json iterator.
func (o *OptDuration) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		v, err := json.ReadDuration(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptDuration", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of float64 to json stream.
func (o OptFloat64) WriteJSON(j *json.Stream) {
	j.WriteFloat64(float64(o.Value))
}

// ReadJSON reads json value of float64 from json iterator.
func (o *OptFloat64) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Value = float64(i.ReadFloat64())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptFloat64", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of int to json stream.
func (o OptInt) WriteJSON(j *json.Stream) {
	j.WriteInt(int(o.Value))
}

// ReadJSON reads json value of int from json iterator.
func (o *OptInt) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.NumberValue:
		o.Set = true
		o.Value = int(i.ReadInt())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptInt", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of string to json stream.
func (o OptNilString) WriteJSON(j *json.Stream) {
	if o.Null {
		j.WriteNil()
		return
	}
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *OptNilString) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Null = false
		o.Value = string(i.ReadString())
		return i.Error
	case json.NilValue:
		var v string
		o.Value = v
		o.Set = true
		o.Null = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptNilString", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of PetType to json stream.
func (o OptPetType) WriteJSON(j *json.Stream) {
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of PetType from json iterator.
func (o *OptPetType) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Value = PetType(i.ReadString())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptPetType", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of string to json stream.
func (o OptString) WriteJSON(j *json.Stream) {
	j.WriteString(string(o.Value))
}

// ReadJSON reads json value of string from json iterator.
func (o *OptString) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		o.Value = string(i.ReadString())
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptString", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of time.Time to json stream.
func (o OptTime) WriteJSON(j *json.Stream, format func(*json.Stream, time.Time)) {
	format(j, o.Value)
}

// ReadJSON reads json value of time.Time from json iterator.
func (o *OptTime) ReadJSON(i *json.Iterator, format func(*json.Iterator) (time.Time, error)) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		v, err := format(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptTime", i.WhatIsNext())
	}
	return nil
}

// WriteJSON writes json value of uuid.UUID to json stream.
func (o OptUUID) WriteJSON(j *json.Stream) {
	json.WriteUUID(j, o.Value)
}

// ReadJSON reads json value of uuid.UUID from json iterator.
func (o *OptUUID) ReadJSON(i *json.Iterator) error {
	switch i.WhatIsNext() {
	case json.StringValue:
		o.Set = true
		v, err := json.ReadUUID(i)
		if err != nil {
			return err
		}
		o.Value = v
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptUUID", i.WhatIsNext())
	}
	return nil
}

// WriteJSON implements json.Marshaler.
func (s Pet) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	j.WriteObjectField("birthday")
	json.WriteDate(j, s.Birthday)
	if s.Friends != nil {
		more.More()
		j.WriteObjectField("friends")
		more.Down()
		j.WriteArrayStart()
		for _, elem := range s.Friends {
			more.More()
			elem.WriteJSON(j)
		}
		j.WriteArrayEnd()
		more.Up()
	}
	more.More()
	j.WriteObjectField("id")
	j.WriteInt64(s.ID)
	more.More()
	j.WriteObjectField("ip")
	json.WriteIP(j, s.IP)
	more.More()
	j.WriteObjectField("ip_v4")
	json.WriteIP(j, s.IPV4)
	more.More()
	j.WriteObjectField("ip_v6")
	json.WriteIP(j, s.IPV6)
	more.More()
	j.WriteObjectField("kind")
	s.Kind.WriteJSON(j)
	more.More()
	j.WriteObjectField("name")
	j.WriteString(s.Name)
	if s.Next.Set {
		more.More()
		j.WriteObjectField("next")
		s.Next.WriteJSON(j)
	}
	more.More()
	j.WriteObjectField("nickname")
	s.Nickname.WriteJSON(j)
	if s.NullStr.Set {
		more.More()
		j.WriteObjectField("nullStr")
		s.NullStr.WriteJSON(j)
	}
	more.More()
	j.WriteObjectField("rate")
	json.WriteDuration(j, s.Rate)
	if s.Tag.Set {
		more.More()
		j.WriteObjectField("tag")
		s.Tag.WriteJSON(j)
	}
	if s.TestArray1 != nil {
		more.More()
		j.WriteObjectField("testArray1")
		more.Down()
		j.WriteArrayStart()
		for _, elem := range s.TestArray1 {
			more.More()
			more.Down()
			j.WriteArrayStart()
			for _, elem := range elem {
				more.More()
				j.WriteString(elem)
			}
			j.WriteArrayEnd()
			more.Up()
		}
		j.WriteArrayEnd()
		more.Up()
	}
	if s.TestDate.Set {
		more.More()
		j.WriteObjectField("testDate")
		s.TestDate.WriteJSON(j, json.WriteDate)
	}
	if s.TestDateTime.Set {
		more.More()
		j.WriteObjectField("testDateTime")
		s.TestDateTime.WriteJSON(j, json.WriteDateTime)
	}
	if s.TestDuration.Set {
		more.More()
		j.WriteObjectField("testDuration")
		s.TestDuration.WriteJSON(j)
	}
	if s.TestFloat1.Set {
		more.More()
		j.WriteObjectField("testFloat1")
		s.TestFloat1.WriteJSON(j)
	}
	if s.TestInteger1.Set {
		more.More()
		j.WriteObjectField("testInteger1")
		s.TestInteger1.WriteJSON(j)
	}
	if s.TestTime.Set {
		more.More()
		j.WriteObjectField("testTime")
		s.TestTime.WriteJSON(j, json.WriteTime)
	}
	if s.Type.Set {
		more.More()
		j.WriteObjectField("type")
		s.Type.WriteJSON(j)
	}
	more.More()
	j.WriteObjectField("uri")
	json.WriteURI(j, s.URI)
	more.More()
	j.WriteObjectField("unique_id")
	json.WriteUUID(j, s.UniqueID)
	j.WriteObjectEnd()
}

// WriteJSONTo writes Pet json value to io.Writer.
func (s Pet) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
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
			// Unsupported kind "array" for field "Friends".
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
		case "kind":
			if err := s.Kind.ReadJSON(i); err != nil {
				i.ReportError("Field Kind", err.Error())
				return false
			}
			return true
		case "name":
			s.Name = i.ReadString()
			return i.Error == nil
		case "next":
			s.Next.Reset()
			if err := s.Next.ReadJSON(i); err != nil {
				i.ReportError("Field Next", err.Error())
				return false
			}
			return true
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
			// Unsupported kind "array" for field "TestArray1".
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
			s.Type.Reset()
			if err := s.Type.ReadJSON(i); err != nil {
				i.ReportError("Field Type", err.Error())
				return false
			}
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

func (PetCreateTextPlainRequest) WriteJSON(j *json.Stream)        {}
func (PetCreateTextPlainRequest) ReadJSON(i *json.Iterator) error { return nil }
func (PetCreateTextPlainRequest) ReadJSONFrom(r io.Reader) error  { return nil }
func (PetCreateTextPlainRequest) WriteJSONTo(w io.Writer) error   { return nil }

// WriteJSON implements json.Marshaler.
func (s PetGetDefault) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	j.WriteObjectField("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetDefault json value to io.Writer.
func (s PetGetDefault) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
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
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes PetGetDefaultStatusCode json value to io.Writer.
func (s PetGetDefaultStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.GetStream(w)
	defer json.PutStream(j)
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

// WriteJSON implements json.Marshaler.
func (s PetKind) WriteJSON(j *json.Stream) {
	j.WriteString(string(s))
}

// ReadJSON reads PetKind from json stream.
func (s *PetKind) ReadJSON(i *json.Iterator) error {
	*s = PetKind(i.ReadString())
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetType) WriteJSON(j *json.Stream) {
	j.WriteString(string(s))
}

// ReadJSON reads PetType from json stream.
func (s *PetType) ReadJSON(i *json.Iterator) error {
	*s = PetType(i.ReadString())
	return i.Error
}
