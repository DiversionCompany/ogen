// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	custom0 "github.com/ogen-go/ogen/internal/integration/customformats/phonetype"
	custom1 "github.com/ogen-go/ogen/internal/integration/customformats/rgbatype"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// PhoneGetParams is parameters of GET /phone operation.
type PhoneGetParams struct {
	// Phone number.
	Phone custom0.Phone
	// Color.
	Color OptRgba
}

func unpackPhoneGetParams(packed middleware.Parameters) (params PhoneGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "phone",
			In:   "query",
		}
		params.Phone = packed[key].(custom0.Phone)
	}
	{
		key := middleware.ParameterKey{
			Name: "color",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Color = v.(OptRgba)
		}
	}
	return params
}

func decodePhoneGetParams(args [0]string, r *http.Request) (params PhoneGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: phone.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "phone",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := formatPhone().DecodeText(val)
				if err != nil {
					return err
				}

				params.Phone = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "phone",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: color.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "color",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotColorVal custom1.RGBA
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := formatRgba().DecodeText(val)
					if err != nil {
						return err
					}

					paramsDotColorVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Color.SetTo(paramsDotColorVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "color",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
