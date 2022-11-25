// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/base64"
	"io"
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

func encodeAllRequestBodiesRequest(
	req AllRequestBodiesReq,
	r *http.Request,
) error {
	switch req := req.(type) {
	case *AllRequestBodiesApplicationJSON:
		const contentType = "application/json"
		e := jx.GetEncoder()
		{
			req.Encode(e)
		}
		encoded := e.Bytes()
		ht.SetBody(r, bytes.NewReader(encoded), contentType)
		return nil
	case *AllRequestBodiesReqApplicationOctetStream:
		const contentType = "application/octet-stream"
		body := req
		ht.SetBody(r, body, contentType)
		return nil
	case *AllRequestBodiesApplicationXWwwFormUrlencoded:
		const contentType = "application/x-www-form-urlencoded"
		request := req

		q := uri.NewQueryEncoder()
		{
			// Encode "name" form field.
			cfg := uri.QueryParameterEncodingConfig{
				Name:    "name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(request.Name))
			}); err != nil {
				return errors.Wrap(err, "encode query")
			}
		}
		{
			// Encode "age" form field.
			cfg := uri.QueryParameterEncodingConfig{
				Name:    "age",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
				if val, ok := request.Age.Get(); ok {
					return e.EncodeValue(conv.IntToString(val))
				}
				return nil
			}); err != nil {
				return errors.Wrap(err, "encode query")
			}
		}
		encoded := q.Values().Encode()
		ht.SetBody(r, strings.NewReader(encoded), contentType)
		return nil
	case *AllRequestBodiesMultipartFormData:
		const contentType = "multipart/form-data"
		request := req

		q := uri.NewQueryEncoder()
		{
			// Encode "name" form field.
			cfg := uri.QueryParameterEncodingConfig{
				Name:    "name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(request.Name))
			}); err != nil {
				return errors.Wrap(err, "encode query")
			}
		}
		{
			// Encode "age" form field.
			cfg := uri.QueryParameterEncodingConfig{
				Name:    "age",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
				if val, ok := request.Age.Get(); ok {
					return e.EncodeValue(conv.IntToString(val))
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
	case *AllRequestBodiesReqTextPlain:
		const contentType = "text/plain"
		body := req
		ht.SetBody(r, body, contentType)
		return nil
	default:
		return errors.Errorf("unexpected request type: %T", req)
	}
}

func encodeAllRequestBodiesOptionalRequest(
	req AllRequestBodiesOptionalReq,
	r *http.Request,
) error {
	switch req := req.(type) {
	case *AllRequestBodiesOptionalReqEmptyBody:
		// Empty body case.
		return nil
	case *AllRequestBodiesOptionalApplicationJSON:
		const contentType = "application/json"
		e := jx.GetEncoder()
		{
			req.Encode(e)
		}
		encoded := e.Bytes()
		ht.SetBody(r, bytes.NewReader(encoded), contentType)
		return nil
	case *AllRequestBodiesOptionalReqApplicationOctetStream:
		const contentType = "application/octet-stream"
		body := req
		ht.SetBody(r, body, contentType)
		return nil
	case *AllRequestBodiesOptionalApplicationXWwwFormUrlencoded:
		const contentType = "application/x-www-form-urlencoded"
		request := req

		q := uri.NewQueryEncoder()
		{
			// Encode "name" form field.
			cfg := uri.QueryParameterEncodingConfig{
				Name:    "name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(request.Name))
			}); err != nil {
				return errors.Wrap(err, "encode query")
			}
		}
		{
			// Encode "age" form field.
			cfg := uri.QueryParameterEncodingConfig{
				Name:    "age",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
				if val, ok := request.Age.Get(); ok {
					return e.EncodeValue(conv.IntToString(val))
				}
				return nil
			}); err != nil {
				return errors.Wrap(err, "encode query")
			}
		}
		encoded := q.Values().Encode()
		ht.SetBody(r, strings.NewReader(encoded), contentType)
		return nil
	case *AllRequestBodiesOptionalMultipartFormData:
		const contentType = "multipart/form-data"
		request := req

		q := uri.NewQueryEncoder()
		{
			// Encode "name" form field.
			cfg := uri.QueryParameterEncodingConfig{
				Name:    "name",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
				return e.EncodeValue(conv.StringToString(request.Name))
			}); err != nil {
				return errors.Wrap(err, "encode query")
			}
		}
		{
			// Encode "age" form field.
			cfg := uri.QueryParameterEncodingConfig{
				Name:    "age",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
				if val, ok := request.Age.Get(); ok {
					return e.EncodeValue(conv.IntToString(val))
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
	case *AllRequestBodiesOptionalReqTextPlain:
		const contentType = "text/plain"
		body := req
		ht.SetBody(r, body, contentType)
		return nil
	default:
		return errors.Errorf("unexpected request type: %T", req)
	}
}

func encodeBase64RequestRequest(
	req Base64RequestReq,
	r *http.Request,
) error {
	const contentType = "text/plain"
	body := ht.CreateBodyWriter(func(w io.Writer) error {
		writer := base64.NewEncoder(base64.StdEncoding, w)
		defer writer.Close()

		_, err := io.Copy(writer, req)
		return err
	})
	ht.SetCloserBody(r, body, contentType)
	return nil
}

func encodeMaskContentTypeRequest(
	req *MaskContentTypeReqWithContentType,
	r *http.Request,
) error {
	contentType := req.ContentType
	if contentType != "" && !ht.MatchContentType("application/*", contentType) {
		return errors.Errorf("%q does not match mask %q", contentType, "application/*")
	}
	{
		req := req.Content
		body := req
		ht.SetBody(r, body, contentType)
		return nil
	}
}

func encodeMaskContentTypeOptionalRequest(
	req *MaskContentTypeOptionalReqWithContentType,
	r *http.Request,
) error {
	contentType := req.ContentType
	if contentType != "" && !ht.MatchContentType("application/*", contentType) {
		return errors.Errorf("%q does not match mask %q", contentType, "application/*")
	}
	{
		req := req.Content
		body := req
		ht.SetBody(r, body, contentType)
		return nil
	}
}
