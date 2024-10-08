// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"
)

func decodeComplicatedParameterNameGetResponse(resp *http.Response) (res *ComplicatedParameterNameGetOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &ComplicatedParameterNameGetOK{}, nil
	}
	err := validate.UnexpectedStatusCode(resp.StatusCode)
	if buf, bodyErr := io.ReadAll(resp.Body); bodyErr == nil {
		err = errors.Wrapf(err, "request failed: %s", string(buf))
	}
	return res, err
}

func decodeContentParametersResponse(resp *http.Response) (res *ContentParameters, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response ContentParameters
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	err := validate.UnexpectedStatusCode(resp.StatusCode)
	if buf, bodyErr := io.ReadAll(resp.Body); bodyErr == nil {
		err = errors.Wrapf(err, "request failed: %s", string(buf))
	}
	return res, err
}

func decodeCookieParameterResponse(resp *http.Response) (res *Value, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response Value
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	err := validate.UnexpectedStatusCode(resp.StatusCode)
	if buf, bodyErr := io.ReadAll(resp.Body); bodyErr == nil {
		err = errors.Wrapf(err, "request failed: %s", string(buf))
	}
	return res, err
}

func decodeHeaderParameterResponse(resp *http.Response) (res *Value, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response Value
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	err := validate.UnexpectedStatusCode(resp.StatusCode)
	if buf, bodyErr := io.ReadAll(resp.Body); bodyErr == nil {
		err = errors.Wrapf(err, "request failed: %s", string(buf))
	}
	return res, err
}

func decodeObjectCookieParameterResponse(resp *http.Response) (res *OneLevelObject, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response OneLevelObject
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	err := validate.UnexpectedStatusCode(resp.StatusCode)
	if buf, bodyErr := io.ReadAll(resp.Body); bodyErr == nil {
		err = errors.Wrapf(err, "request failed: %s", string(buf))
	}
	return res, err
}

func decodeObjectQueryParameterResponse(resp *http.Response) (res *ObjectQueryParameterOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response ObjectQueryParameterOK
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	err := validate.UnexpectedStatusCode(resp.StatusCode)
	if buf, bodyErr := io.ReadAll(resp.Body); bodyErr == nil {
		err = errors.Wrapf(err, "request failed: %s", string(buf))
	}
	return res, err
}

func decodePathParameterResponse(resp *http.Response) (res *Value, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response Value
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	err := validate.UnexpectedStatusCode(resp.StatusCode)
	if buf, bodyErr := io.ReadAll(resp.Body); bodyErr == nil {
		err = errors.Wrapf(err, "request failed: %s", string(buf))
	}
	return res, err
}

func decodeSameNameResponse(resp *http.Response) (res *SameNameOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &SameNameOK{}, nil
	}
	err := validate.UnexpectedStatusCode(resp.StatusCode)
	if buf, bodyErr := io.ReadAll(resp.Body); bodyErr == nil {
		err = errors.Wrapf(err, "request failed: %s", string(buf))
	}
	return res, err
}

func decodeSimilarNamesResponse(resp *http.Response) (res *SimilarNamesOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &SimilarNamesOK{}, nil
	}
	err := validate.UnexpectedStatusCode(resp.StatusCode)
	if buf, bodyErr := io.ReadAll(resp.Body); bodyErr == nil {
		err = errors.Wrapf(err, "request failed: %s", string(buf))
	}
	return res, err
}
