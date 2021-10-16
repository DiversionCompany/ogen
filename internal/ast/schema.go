package ast

import (
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/xerrors"
)

type SchemaKind = string

const (
	KindStruct    SchemaKind = "struct"
	KindAlias     SchemaKind = "alias"
	KindPrimitive SchemaKind = "primitive"
	KindArray     SchemaKind = "array"
	KindEnum      SchemaKind = "enum"
	KindPointer   SchemaKind = "pointer"
)

type JSON struct {
	Write string
	Read  string
}

type Schema struct {
	Kind        SchemaKind
	Name        string
	Description string
	Doc         string

	AliasTo    *Schema
	PointerTo  *Schema
	Primitive  string
	Item       *Schema
	EnumValues []interface{}
	Fields     []SchemaField

	Implements map[*Interface]struct{}

	// Numeric validation.
	MultipleOf       *uint64
	Minimum          *int64
	Maximum          *int64
	ExclusiveMinimum bool
	ExclusiveMaximum bool

	// String validation.
	// Pattern   string
	MaxLength *uint64
	MinLength *int64

	// Array validation.
	MaxItems *uint64
	MinItems *uint64
	// UniqueItems bool

	// Struct validation.
	// MaxProperties *uint64
	// MinProperties *uint64

	JSON *JSON
}

func (s *Schema) IsInteger() bool {
	switch s.Primitive {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		return true
	default:
		return false
	}
}

func (s *Schema) IsFloat() bool {
	switch s.Primitive {
	case "float32", "float64":
		return true
	default:
		return false
	}
}

func (s *Schema) IsNumeric() bool { return s.IsInteger() || s.IsFloat() }

func (s *Schema) NeedValidation() bool {
	return s.needValidation(map[*Schema]struct{}{})
}

func (s *Schema) needValidation(visited map[*Schema]struct{}) (result bool) {
	if s == nil {
		return false
	}

	if _, ok := visited[s]; ok {
		return false
	}

	visited[s] = struct{}{}

	switch s.Kind {
	case KindPrimitive:
		if s.IsNumeric() && (s.Minimum != nil || s.Maximum != nil || s.MultipleOf != nil) {
			return true
		}
		if s.Primitive == "string" && (s.MinLength != nil || s.MaxLength != nil) {
			return true
		}
		return false
	case KindEnum:
		return true
	case KindAlias:
		return s.AliasTo.needValidation(visited)
	case KindPointer:
		return s.PointerTo.needValidation(visited)
	case KindArray:
		if s.MinItems != nil || s.MaxItems != nil {
			return true
		}
		// Prevent infinite recursion.
		if s.Item == s {
			return false
		}
		return s.Item.needValidation(visited)
	case KindStruct:
		for _, f := range s.Fields {
			if f.Type.needValidation(visited) {
				return true
			}
		}
		return false
	default:
		panic("unreachable")
	}
}

type SchemaField struct {
	Name string
	Type *Schema
	Tag  string
}

func (s Schema) Type() string {
	switch s.Kind {
	case KindStruct:
		return s.Name
	case KindAlias:
		return s.Name
	case KindPrimitive:
		return s.Primitive
	case KindArray:
		return "[]" + s.Item.Type()
	case KindEnum:
		return s.Name
	case KindPointer:
		return "*" + s.PointerTo.Type()
	default:
		panic(fmt.Errorf("unexpected SchemaKind: %s", s.Kind))
	}
}

func (s Schema) Is(vs ...SchemaKind) bool {
	for _, v := range vs {
		if s.Kind == v {
			return true
		}
	}

	return false
}

func (s *Schema) Implement(iface *Interface) {
	if s.Is(KindPrimitive, KindArray, KindPointer) {
		panic("unreachable")
	}

	if s.Implements == nil {
		s.Implements = map[*Interface]struct{}{}
	}
	if iface.Implementations == nil {
		iface.Implementations = map[*Schema]struct{}{}
	}

	iface.Implementations[s] = struct{}{}
	s.Implements[iface] = struct{}{}
}

func (s *Schema) Unimplement(iface *Interface) {
	delete(iface.Implementations, s)
	delete(s.Implements, iface)
}

func (s *Schema) Methods() []string {
	ms := make(map[string]struct{})
	for iface := range s.Implements {
		for m := range iface.Methods {
			ms[m] = struct{}{}
		}
	}

	var result []string
	for m := range ms {
		result = append(result, m)
	}
	sort.Strings(result)
	return result
}

func Struct(name string) *Schema {
	return &Schema{
		Kind: KindStruct,
		Name: name,
	}
}

func Primitive(typ string) *Schema {
	return &Schema{
		Kind:      KindPrimitive,
		Primitive: typ,
	}
}

func Alias(name string, typ *Schema) *Schema {
	return &Schema{
		Kind:    KindAlias,
		Name:    name,
		AliasTo: typ,
	}
}

func Pointer(to *Schema) *Schema {
	return &Schema{
		Kind:      KindPointer,
		PointerTo: to,
	}
}

func Array(item *Schema) *Schema {
	return &Schema{
		Kind: KindArray,
		Item: item,
	}
}

func Enum(name, typ string, rawValues []json.RawMessage) (*Schema, error) {
	var (
		values []interface{}
		uniq   = map[interface{}]struct{}{}
	)
	for _, raw := range rawValues {
		val, err := parseJsonValue(typ, raw)
		if err != nil {
			if xerrors.Is(err, errNullValue) {
				continue
			}
			return nil, xerrors.Errorf("parse value '%s': %w", raw, err)
		}

		if _, found := uniq[val]; found {
			return nil, xerrors.Errorf("duplicate enum value: '%v'", val)
		}

		uniq[val] = struct{}{}
		values = append(values, val)
	}

	return &Schema{
		Kind:       KindEnum,
		Name:       name,
		Primitive:  typ,
		EnumValues: values,
	}, nil
}

func Iface(name string) *Interface {
	return &Interface{
		Name:            name,
		Methods:         map[string]struct{}{},
		Implementations: map[*Schema]struct{}{},
	}
}

func CreateRequestBody() *RequestBody {
	return &RequestBody{
		Contents: map[string]*Schema{},
	}
}

func CreateResponse() *Response {
	return &Response{
		Contents: map[string]*Schema{},
	}
}
