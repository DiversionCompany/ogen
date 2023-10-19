// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"mime"
	"mime/multipart"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeNullableStringsRequest(
	req NilString,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeObjectsWithConflictingArrayPropertyRequest(
	req *ObjectsWithConflictingArrayPropertyReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeObjectsWithConflictingPropertiesRequest(
	req *ObjectsWithConflictingPropertiesReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeReferencedAllofRequest(
	req ReferencedAllofReq,
	r *http.Request,
) error {
	switch req := req.(type) {
	case *Robot:
		const contentType = "application/json"
		e := new(jx.Encoder)
		{
			req.Encode(e)
		}
		encoded := e.Bytes()
		ht.SetBody(r, bytes.NewReader(encoded), contentType)
		return nil
	case *RobotMultipart:
		const contentType = "multipart/form-data"
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
				return errors.Wrap(err, "encode query")
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
				return errors.Wrap(err, "encode query")
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
				return errors.Wrap(err, "encode query")
			}
		}
		body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
			if err := q.WriteMultipart(w); err != nil {
				return errors.Wrap(err, "write multipart")
			}
			return nil
		})
		ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
		return nil
	default:
		return errors.Errorf("unexpected request type: %T", req)
	}
}

func encodeReferencedAllofOptionalRequest(
	req ReferencedAllofOptionalReq,
	r *http.Request,
) error {
	switch req := req.(type) {
	case *ReferencedAllofOptionalReqEmptyBody:
		// Empty body case.
		return nil
	case *Robot:
		const contentType = "application/json"
		e := new(jx.Encoder)
		{
			req.Encode(e)
		}
		encoded := e.Bytes()
		ht.SetBody(r, bytes.NewReader(encoded), contentType)
		return nil
	case *RobotMultipart:
		const contentType = "multipart/form-data"
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
				return errors.Wrap(err, "encode query")
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
				return errors.Wrap(err, "encode query")
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
				return errors.Wrap(err, "encode query")
			}
		}
		body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
			if err := q.WriteMultipart(w); err != nil {
				return errors.Wrap(err, "write multipart")
			}
			return nil
		})
		ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
		return nil
	default:
		return errors.Errorf("unexpected request type: %T", req)
	}
}

func encodeSimpleIntegerRequest(
	req int,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		e.Int(req)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeSimpleObjectsRequest(
	req *SimpleObjectsReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeStringsNotypeRequest(
	req NilString,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}
