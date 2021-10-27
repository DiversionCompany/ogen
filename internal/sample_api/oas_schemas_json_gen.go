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
	more.More()
	j.WriteObjectField("id")
	s.ID.WriteJSON(j)
	j.WriteObjectEnd()
}

// ReadJSON reads Data from json stream.
func (s *Data) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "id":
			if err := func() error {
				if err := s.ID.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarGetResNotFound) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads FoobarGetResNotFound from json stream.
func (s *FoobarGetResNotFound) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPostDef) WriteJSON(j *json.Stream) {
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

// ReadJSON reads FoobarPostDef from json stream.
func (s *FoobarPostDef) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "code":
			if err := func() error {
				s.Code = int64(i.ReadInt64())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "message":
			if err := func() error {
				s.Message = string(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPostDefStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads FoobarPostDefStatusCode from json stream.
func (s *FoobarPostDefStatusCode) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPostResNotFound) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads FoobarPostResNotFound from json stream.
func (s *FoobarPostResNotFound) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPutDef) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads FoobarPutDef from json stream.
func (s *FoobarPutDef) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FoobarPutDefStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads FoobarPutDefStatusCode from json stream.
func (s *FoobarPutDefStatusCode) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s ID) WriteJSON(j *json.Stream) {
	switch s.Type {
	case StringID:
		j.WriteString(s.String)
	case IntID:
		j.WriteInt(s.Int)
	}
}

// ReadJSON reads value from json stream.
func (s *ID) ReadJSON(i *json.Iterator) error {
	switch t := i.WhatIsNext(); t {
	case json.StringValue:
		if err := func() error {
			s.String = string(i.ReadString())
			return i.Error
		}(); err != nil {
			return err
		}
		s.Type = StringID
	case json.NumberValue:
		if err := func() error {
			s.Int = int(i.ReadInt())
			return i.Error
		}(); err != nil {
			return err
		}
		s.Type = IntID
	default:
		return fmt.Errorf("unexpected json type %d", t)
	}
	return nil
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
}

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
	if s.Primary != nil {
		more.More()
		j.WriteObjectField("primary")
		s.Primary.WriteJSON(j)
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
	j.WriteObjectField("unique_id")
	json.WriteUUID(j, s.UniqueID)
	more.More()
	j.WriteObjectField("uri")
	json.WriteURI(j, s.URI)
	j.WriteObjectEnd()
}

// ReadJSON reads Pet from json stream.
func (s *Pet) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "birthday":
			if err := func() error {
				v, err := json.ReadDate(i)
				s.Birthday = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "friends":
			if err := func() error {
				s.Friends = nil
				var retErr error
				i.ReadArrayCB(func(i *json.Iterator) bool {
					var elem Pet
					if err := func() error {
						if err := elem.ReadJSON(i); err != nil {
							return err
						}
						return i.Error
					}(); err != nil {
						retErr = err
						return false
					}
					s.Friends = append(s.Friends, elem)
					return true
				})
				if retErr != nil {
					return retErr
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "id":
			if err := func() error {
				s.ID = int64(i.ReadInt64())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "ip":
			if err := func() error {
				v, err := json.ReadIP(i)
				s.IP = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "ip_v4":
			if err := func() error {
				v, err := json.ReadIP(i)
				s.IPV4 = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "ip_v6":
			if err := func() error {
				v, err := json.ReadIP(i)
				s.IPV6 = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "kind":
			if err := func() error {
				if err := s.Kind.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "name":
			if err := func() error {
				s.Name = string(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "next":
			if err := func() error {
				s.Next.Reset()
				if err := s.Next.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "nickname":
			if err := func() error {
				if err := s.Nickname.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "nullStr":
			if err := func() error {
				s.NullStr.Reset()
				if err := s.NullStr.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "primary":
			if err := func() error {
				s.Primary = nil
				var elem Pet
				if err := func() error {
					if err := elem.ReadJSON(i); err != nil {
						return err
					}
					return i.Error
				}(); err != nil {
					return err
				}
				s.Primary = &elem
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "rate":
			if err := func() error {
				v, err := json.ReadDuration(i)
				s.Rate = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "tag":
			if err := func() error {
				s.Tag.Reset()
				if err := s.Tag.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testArray1":
			if err := func() error {
				s.TestArray1 = nil
				var retErr error
				i.ReadArrayCB(func(i *json.Iterator) bool {
					var elem []string
					if err := func() error {
						elem = nil
						var retErr error
						i.ReadArrayCB(func(i *json.Iterator) bool {
							var elemElem string
							if err := func() error {
								elemElem = string(i.ReadString())
								return i.Error
							}(); err != nil {
								retErr = err
								return false
							}
							elem = append(elem, elemElem)
							return true
						})
						if retErr != nil {
							return retErr
						}
						return i.Error
					}(); err != nil {
						retErr = err
						return false
					}
					s.TestArray1 = append(s.TestArray1, elem)
					return true
				})
				if retErr != nil {
					return retErr
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testDate":
			if err := func() error {
				s.TestDate.Reset()
				if err := s.TestDate.ReadJSON(i, json.ReadDate); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testDateTime":
			if err := func() error {
				s.TestDateTime.Reset()
				if err := s.TestDateTime.ReadJSON(i, json.ReadDateTime); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testDuration":
			if err := func() error {
				s.TestDuration.Reset()
				if err := s.TestDuration.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testFloat1":
			if err := func() error {
				s.TestFloat1.Reset()
				if err := s.TestFloat1.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testInteger1":
			if err := func() error {
				s.TestInteger1.Reset()
				if err := s.TestInteger1.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "testTime":
			if err := func() error {
				s.TestTime.Reset()
				if err := s.TestTime.ReadJSON(i, json.ReadTime); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.ReadJSON(i); err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "unique_id":
			if err := func() error {
				v, err := json.ReadUUID(i)
				s.UniqueID = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		case "uri":
			if err := func() error {
				v, err := json.ReadURI(i)
				s.URI = v
				if err != nil {
					return err
				}
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetCreateReqTextPlain) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads PetCreateReqTextPlain from json stream.
func (s *PetCreateReqTextPlain) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetDef) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	j.WriteObjectField("message")
	j.WriteString(s.Message)
	j.WriteObjectEnd()
}

// ReadJSON reads PetGetDef from json stream.
func (s *PetGetDef) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "message":
			if err := func() error {
				s.Message = string(i.ReadString())
				return i.Error
			}(); err != nil {
				retErr = err
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetGetDefStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads PetGetDefStatusCode from json stream.
func (s *PetGetDefStatusCode) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
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

func (PetName) WriteJSON(j *json.Stream)        {}
func (PetName) ReadJSON(i *json.Iterator) error { return nil }

// WriteJSON implements json.Marshaler.
func (s PetType) WriteJSON(j *json.Stream) {
	j.WriteString(string(s))
}

// ReadJSON reads PetType from json stream.
func (s *PetType) ReadJSON(i *json.Iterator) error {
	*s = PetType(i.ReadString())
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetUpdateNameAliasPostDef) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads PetUpdateNameAliasPostDef from json stream.
func (s *PetUpdateNameAliasPostDef) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetUpdateNameAliasPostDefStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads PetUpdateNameAliasPostDefStatusCode from json stream.
func (s *PetUpdateNameAliasPostDefStatusCode) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetUpdateNamePostDef) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads PetUpdateNamePostDef from json stream.
func (s *PetUpdateNamePostDef) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PetUpdateNamePostDefStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	more := json.NewMore(j)
	defer more.Reset()
	j.WriteObjectEnd()
}

// ReadJSON reads PetUpdateNamePostDefStatusCode from json stream.
func (s *PetUpdateNamePostDefStatusCode) ReadJSON(i *json.Iterator) error {
	var retErr error
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	if retErr != nil {
		return retErr
	}
	return i.Error
}
