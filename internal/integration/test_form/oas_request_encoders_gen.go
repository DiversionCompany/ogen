// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeTestFormURLEncodedRequest(
	req *TestForm,
	r *http.Request,
) error {
	const contentType = "application/x-www-form-urlencoded"
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "id" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.ID.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "uuid" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "uuid",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.UUID.Get(); ok {
				return e.EncodeValue(conv.UUIDToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "description" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "description",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Description))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "array" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "array",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range request.Array {
					if err := func() error {
						return e.EncodeValue(conv.StringToString(item))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "object" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "object",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Object.Get(); ok {
				return val.EncodeURI(e)
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "deepObject" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "deepObject",
			Style:   uri.QueryStyleDeepObject,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.DeepObject.Get(); ok {
				return val.EncodeURI(e)
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	encoded := q.Values().Encode()
	ht.SetBody(r, strings.NewReader(encoded), contentType)
	return nil
}

func encodeTestMultipartRequest(
	req *TestForm,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "id" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.ID.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "uuid" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "uuid",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.UUID.Get(); ok {
				return e.EncodeValue(conv.UUIDToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "description" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "description",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Description))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "array" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "array",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range request.Array {
					if err := func() error {
						return e.EncodeValue(conv.StringToString(item))
					}(); err != nil {
						return errors.Wrapf(err, "[%d]", i)
					}
				}
				return nil
			})
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "object" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "object",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Object.Get(); ok {
				return val.EncodeURI(e)
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "deepObject" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "deepObject",
			Style:   uri.QueryStyleDeepObject,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.DeepObject.Get(); ok {
				return val.EncodeURI(e)
			}
			return nil
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
}

func encodeTestMultipartUploadRequest(
	req *TestMultipartUploadReqForm,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	request := req

	q := uri.NewQueryEncoder()
	{
		// Encode "orderId" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "orderId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.OrderId.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "userId" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "userId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.UserId.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := request.File.WriteMultipart("file", w); err != nil {
			return errors.Wrap(err, "write \"file\"")
		}
		if val, ok := request.OptionalFile.Get(); ok {
			if err := val.WriteMultipart("optional_file", w); err != nil {
				return errors.Wrap(err, "write \"optional_file\"")
			}
		}
		if err := func() error {
			for idx, val := range request.Files {
				if err := val.WriteMultipart("files", w); err != nil {
					return errors.Wrapf(err, "file [%d]", idx)
				}
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "write \"files\"")
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}

func encodeTestShareFormSchemaRequest(
	req TestShareFormSchemaReq,
	r *http.Request,
) error {
	switch req := req.(type) {
	case *SharedRequest:
		const contentType = "application/json"
		e := jx.GetEncoder()
		{
			req.Encode(e)
		}
		encoded := e.Bytes()
		ht.SetBody(r, bytes.NewReader(encoded), contentType)
		return nil
	case *SharedRequestForm:
		const contentType = "multipart/form-data"
		request := req

		q := uri.NewQueryEncoder()
		{
			// Encode "filename" form field.
			cfg := uri.QueryParameterEncodingConfig{
				Name:    "filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
				if val, ok := request.Filename.Get(); ok {
					return e.EncodeValue(conv.StringToString(val))
				}
				return nil
			}); err != nil {
				return errors.Wrap(err, "encode query")
			}
		}
		body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
			if val, ok := request.File.Get(); ok {
				if err := val.WriteMultipart("file", w); err != nil {
					return errors.Wrap(err, "write \"file\"")
				}
			}
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
