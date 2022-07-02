// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode encodes AllRequestBodiesApplicationJSON as json.
func (s AllRequestBodiesApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := SimpleObject(s)

	unwrapped.Encode(e)
}

// Decode decodes AllRequestBodiesApplicationJSON from json.
func (s *AllRequestBodiesApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AllRequestBodiesApplicationJSON to nil")
	}
	var unwrapped SimpleObject
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AllRequestBodiesApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AllRequestBodiesApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AllRequestBodiesApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AllRequestBodiesApplicationXWwwFormUrlencoded as json.
func (s AllRequestBodiesApplicationXWwwFormUrlencoded) Encode(e *jx.Encoder) {
	unwrapped := SimpleObject(s)

	unwrapped.Encode(e)
}

// Decode decodes AllRequestBodiesApplicationXWwwFormUrlencoded from json.
func (s *AllRequestBodiesApplicationXWwwFormUrlencoded) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AllRequestBodiesApplicationXWwwFormUrlencoded to nil")
	}
	var unwrapped SimpleObject
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AllRequestBodiesApplicationXWwwFormUrlencoded(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AllRequestBodiesApplicationXWwwFormUrlencoded) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AllRequestBodiesApplicationXWwwFormUrlencoded) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AllRequestBodiesMultipartFormData as json.
func (s AllRequestBodiesMultipartFormData) Encode(e *jx.Encoder) {
	unwrapped := SimpleObject(s)

	unwrapped.Encode(e)
}

// Decode decodes AllRequestBodiesMultipartFormData from json.
func (s *AllRequestBodiesMultipartFormData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AllRequestBodiesMultipartFormData to nil")
	}
	var unwrapped SimpleObject
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AllRequestBodiesMultipartFormData(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AllRequestBodiesMultipartFormData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AllRequestBodiesMultipartFormData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AllRequestBodiesOptionalApplicationJSON as json.
func (s AllRequestBodiesOptionalApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := OptSimpleObject(s)
	if unwrapped.Set {
		unwrapped.Encode(e)
	}
}

// Decode decodes AllRequestBodiesOptionalApplicationJSON from json.
func (s *AllRequestBodiesOptionalApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AllRequestBodiesOptionalApplicationJSON to nil")
	}
	var unwrapped OptSimpleObject
	if err := func() error {
		unwrapped.Reset()
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AllRequestBodiesOptionalApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AllRequestBodiesOptionalApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AllRequestBodiesOptionalApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AllRequestBodiesOptionalApplicationXWwwFormUrlencoded as json.
func (s AllRequestBodiesOptionalApplicationXWwwFormUrlencoded) Encode(e *jx.Encoder) {
	unwrapped := OptSimpleObject(s)
	if unwrapped.Set {
		unwrapped.Encode(e)
	}
}

// Decode decodes AllRequestBodiesOptionalApplicationXWwwFormUrlencoded from json.
func (s *AllRequestBodiesOptionalApplicationXWwwFormUrlencoded) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AllRequestBodiesOptionalApplicationXWwwFormUrlencoded to nil")
	}
	var unwrapped OptSimpleObject
	if err := func() error {
		unwrapped.Reset()
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AllRequestBodiesOptionalApplicationXWwwFormUrlencoded(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AllRequestBodiesOptionalApplicationXWwwFormUrlencoded) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AllRequestBodiesOptionalApplicationXWwwFormUrlencoded) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AllRequestBodiesOptionalMultipartFormData as json.
func (s AllRequestBodiesOptionalMultipartFormData) Encode(e *jx.Encoder) {
	unwrapped := OptSimpleObject(s)
	if unwrapped.Set {
		unwrapped.Encode(e)
	}
}

// Decode decodes AllRequestBodiesOptionalMultipartFormData from json.
func (s *AllRequestBodiesOptionalMultipartFormData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AllRequestBodiesOptionalMultipartFormData to nil")
	}
	var unwrapped OptSimpleObject
	if err := func() error {
		unwrapped.Reset()
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AllRequestBodiesOptionalMultipartFormData(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AllRequestBodiesOptionalMultipartFormData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AllRequestBodiesOptionalMultipartFormData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SimpleObject as json.
func (o OptSimpleObject) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes SimpleObject from json.
func (o *OptSimpleObject) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSimpleObject to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSimpleObject) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSimpleObject) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s SimpleObject) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s SimpleObject) encodeFields(e *jx.Encoder) {
	{

		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		if s.Age.Set {
			e.FieldStart("age")
			s.Age.Encode(e)
		}
	}
}

var jsonFieldsNameOfSimpleObject = [2]string{
	0: "name",
	1: "age",
}

// Decode decodes SimpleObject from json.
func (s *SimpleObject) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SimpleObject to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "age":
			if err := func() error {
				s.Age.Reset()
				if err := s.Age.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SimpleObject")
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
				if fieldIdx < len(jsonFieldsNameOfSimpleObject) {
					name = jsonFieldsNameOfSimpleObject[fieldIdx]
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

// MarshalJSON implements stdjson.Marshaler.
func (s SimpleObject) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SimpleObject) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
