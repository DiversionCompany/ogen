// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s Pets) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s *PetsHeaders) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Response.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "Response",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
