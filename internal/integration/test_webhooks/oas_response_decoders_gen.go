// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

func decodePublishEventResponse(resp *http.Response) (res *Event, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Event
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "decode \"application/json\"")
			}
			if err := d.Skip(); err != io.EOF {
				return res, errors.New("unexpected trailing data")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	// Convenient error response.
	defRes, err := func() (res *ErrorStatusCode, err error) {
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "decode \"application/json\"")
			}
			if err := d.Skip(); err != io.EOF {
				return res, errors.New("unexpected trailing data")
			}
			return &ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}()
	if err != nil {
		return res, errors.Wrap(err, "default")
	}
	return res, errors.Wrap(defRes, "error")
}

func decodeStatusWebhookResponse(resp *http.Response) (res *StatusWebhookOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response StatusWebhookOK
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "decode \"application/json\"")
			}
			if err := d.Skip(); err != io.EOF {
				return res, errors.New("unexpected trailing data")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeUpdateDeleteResponse(resp *http.Response) (res UpdateDeleteRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &UpdateDeleteOK{}, nil
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
		var response Error
		if err := func() error {
			if err := response.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := d.Skip(); err != io.EOF {
			return res, errors.New("unexpected trailing data")
		}
		return &ErrorStatusCode{
			StatusCode: resp.StatusCode,
			Response:   response,
		}, nil
	default:
		return res, validate.InvalidContentType(ct)
	}
}

func decodeUpdateWebhookResponse(resp *http.Response) (res UpdateWebhookRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response WebhookResponse
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "decode \"application/json\"")
			}
			if err := d.Skip(); err != io.EOF {
				return res, errors.New("unexpected trailing data")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	// Default response.
	ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return res, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}

		d := jx.DecodeBytes(b)
		var response Error
		if err := func() error {
			if err := response.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := d.Skip(); err != io.EOF {
			return res, errors.New("unexpected trailing data")
		}
		return &ErrorStatusCode{
			StatusCode: resp.StatusCode,
			Response:   response,
		}, nil
	default:
		return res, validate.InvalidContentType(ct)
	}
}
