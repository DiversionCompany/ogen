// Code generated by ogen, DO NOT EDIT.

package api

import (
	"mime"
	"mime/multipart"
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeUserPassloginPostRequest(
	req OptUserPassloginPostReq,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	request := req.Value

	q := uri.NewQueryEncoder()
	{
		// Encode "passcode" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "passcode",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Passcode))
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
	ht.SetBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}

func encodeUserPostingPostRequest(
	req OptUserPostingPostReqForm,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	request := req.Value

	q := uri.NewQueryEncoder()
	{
		// Encode "captcha_type" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "captcha_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(string(request.CaptchaType)))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "board" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Board))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "thread" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "thread",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Thread.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "name" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Name.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "email" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "email",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Email.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "tags" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "tags",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Tags.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "subject" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "subject",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Subject.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "comment" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "comment",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Comment.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "icon" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "icon",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Icon.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "op_mark" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "op_mark",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.OpMark.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "file[]" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "file[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range request.File {
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
	body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	ht.SetBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}

func encodeUserReportPostRequest(
	req OptUserReportPostReq,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	request := req.Value

	q := uri.NewQueryEncoder()
	{
		// Encode "board" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "board",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Board))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "thread" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "thread",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(request.Thread))
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "post" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "post",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeArray(func(e uri.Encoder) error {
				for i, item := range request.Post {
					if err := func() error {
						return e.EncodeValue(conv.IntToString(item))
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
		// Encode "comment" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "comment",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(request.Comment))
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
	ht.SetBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}
