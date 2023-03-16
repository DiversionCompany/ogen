// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
)

// ComplicatedParameterNameGetOK is response for ComplicatedParameterNameGet operation.
type ComplicatedParameterNameGetOK struct{}

type ContentParameters struct {
	Query  User `json:"query"`
	Path   User `json:"path"`
	Header User `json:"header"`
	Cookie User `json:"cookie"`
}

// GetQuery returns the value of Query.
func (s *ContentParameters) GetQuery() User {
	return s.Query
}

// GetPath returns the value of Path.
func (s *ContentParameters) GetPath() User {
	return s.Path
}

// GetHeader returns the value of Header.
func (s *ContentParameters) GetHeader() User {
	return s.Header
}

// GetCookie returns the value of Cookie.
func (s *ContentParameters) GetCookie() User {
	return s.Cookie
}

// SetQuery sets the value of Query.
func (s *ContentParameters) SetQuery(val User) {
	s.Query = val
}

// SetPath sets the value of Path.
func (s *ContentParameters) SetPath(val User) {
	s.Path = val
}

// SetHeader sets the value of Header.
func (s *ContentParameters) SetHeader(val User) {
	s.Header = val
}

// SetCookie sets the value of Cookie.
func (s *ContentParameters) SetCookie(val User) {
	s.Cookie = val
}

type ObjectQueryParameterOK struct {
	Style string         `json:"style"`
	Value OneLevelObject `json:"value"`
}

// GetStyle returns the value of Style.
func (s *ObjectQueryParameterOK) GetStyle() string {
	return s.Style
}

// GetValue returns the value of Value.
func (s *ObjectQueryParameterOK) GetValue() OneLevelObject {
	return s.Value
}

// SetStyle sets the value of Style.
func (s *ObjectQueryParameterOK) SetStyle(val string) {
	s.Style = val
}

// SetValue sets the value of Value.
func (s *ObjectQueryParameterOK) SetValue(val OneLevelObject) {
	s.Value = val
}

// Ref: #/components/schemas/OneLevelObject
type OneLevelObject struct {
	Min    int    `json:"min"`
	Max    int    `json:"max"`
	Filter string `json:"filter"`
}

// GetMin returns the value of Min.
func (s *OneLevelObject) GetMin() int {
	return s.Min
}

// GetMax returns the value of Max.
func (s *OneLevelObject) GetMax() int {
	return s.Max
}

// GetFilter returns the value of Filter.
func (s *OneLevelObject) GetFilter() string {
	return s.Filter
}

// SetMin sets the value of Min.
func (s *OneLevelObject) SetMin(val int) {
	s.Min = val
}

// SetMax sets the value of Max.
func (s *OneLevelObject) SetMax(val int) {
	s.Max = val
}

// SetFilter sets the value of Filter.
func (s *OneLevelObject) SetFilter(val string) {
	s.Filter = val
}

// NewOptOneLevelObject returns new OptOneLevelObject with value set to v.
func NewOptOneLevelObject(v OneLevelObject) OptOneLevelObject {
	return OptOneLevelObject{
		Value: v,
		Set:   true,
	}
}

// OptOneLevelObject is optional OneLevelObject.
type OptOneLevelObject struct {
	Value OneLevelObject
	Set   bool
}

// IsSet returns true if OptOneLevelObject was set.
func (o OptOneLevelObject) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptOneLevelObject) Reset() {
	var v OneLevelObject
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptOneLevelObject) SetTo(v OneLevelObject) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptOneLevelObject) Get() (v OneLevelObject, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptOneLevelObject) Or(d OneLevelObject) OneLevelObject {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// SameNameOK is response for SameName operation.
type SameNameOK struct{}

// Ref: #/components/schemas/User
type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Role     UserRole `json:"role"`
	Friends  []User   `json:"friends"`
}

// GetID returns the value of ID.
func (s *User) GetID() int {
	return s.ID
}

// GetUsername returns the value of Username.
func (s *User) GetUsername() string {
	return s.Username
}

// GetRole returns the value of Role.
func (s *User) GetRole() UserRole {
	return s.Role
}

// GetFriends returns the value of Friends.
func (s *User) GetFriends() []User {
	return s.Friends
}

// SetID sets the value of ID.
func (s *User) SetID(val int) {
	s.ID = val
}

// SetUsername sets the value of Username.
func (s *User) SetUsername(val string) {
	s.Username = val
}

// SetRole sets the value of Role.
func (s *User) SetRole(val UserRole) {
	s.Role = val
}

// SetFriends sets the value of Friends.
func (s *User) SetFriends(val []User) {
	s.Friends = val
}

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
	UserRoleBot   UserRole = "bot"
)

// MarshalText implements encoding.TextMarshaler.
func (s UserRole) MarshalText() ([]byte, error) {
	switch s {
	case UserRoleAdmin:
		return []byte(s), nil
	case UserRoleUser:
		return []byte(s), nil
	case UserRoleBot:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *UserRole) UnmarshalText(data []byte) error {
	switch UserRole(data) {
	case UserRoleAdmin:
		*s = UserRoleAdmin
		return nil
	case UserRoleUser:
		*s = UserRoleUser
		return nil
	case UserRoleBot:
		*s = UserRoleBot
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/Value
type Value struct {
	Value string `json:"value"`
}

// GetValue returns the value of Value.
func (s *Value) GetValue() string {
	return s.Value
}

// SetValue sets the value of Value.
func (s *Value) SetValue(val string) {
	s.Value = val
}
