// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeNullableStringsRequestJSON(
	req NilString,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeObjectsWithConflictingArrayPropertyRequestJSON(
	req ObjectsWithConflictingArrayPropertyReq,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeObjectsWithConflictingPropertiesRequestJSON(
	req ObjectsWithConflictingPropertiesReq,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeReferencedAllofRequestJSON(
	req ReferencedAllofApplicationJSON,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeReferencedAllofRequest(
	req ReferencedAllofMultipartFormData,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	contentType string,
	rerr error,
) {
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "state" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "state",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(request.State)))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "id" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.UUIDToString(request.ID))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "location" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "location",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return request.Location.EncodeURI(e)
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	getBody, contentType := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	return getBody, contentType, nil
}
func encodeReferencedAllofOptionalRequestJSON(
	req ReferencedAllofOptionalApplicationJSON,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	if !req.Set {
		// Return nil callback if value is not set.
		return
	}
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeReferencedAllofOptionalRequest(
	req ReferencedAllofOptionalMultipartFormData,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	contentType string,
	rerr error,
) {
	if !req.Set {
		// Return nil callback if value is not set.
		return
	}
	request := req.Value

	q := uri.NewQueryEncoder()
	{
		// Encode "state" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "state",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(request.State)))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "id" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.UUIDToString(request.ID))
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "location" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "location",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return request.Location.EncodeURI(e)
		}); err != nil {
			return data, "", errors.Wrap(err, "encode query")
		}
	}
	getBody, contentType := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	return getBody, contentType, nil
}
func encodeSimpleIntegerRequestJSON(
	req int,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	e.Int(req)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
func encodeSimpleObjectsRequestJSON(
	req SimpleObjectsReq,
	span trace.Span,
) (
	data func() (io.ReadCloser, error),
	rerr error,
) {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	return func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(encoded)), nil
	}, nil
}
