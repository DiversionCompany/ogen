// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"net/http"

	"github.com/go-faster/jx"

	ht "github.com/ogen-go/ogen/http"
)

func encodeDefaultTestRequest(
	req *DefaultTest,
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

func encodeFoobarPostRequest(
	req OptPet,
	r *http.Request,
) error {
	const contentType = "application/json"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := new(jx.Encoder)
	{
		if req.Set {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeOneofBugRequest(
	req *OneOfBugs,
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

func encodePetCreateRequest(
	req OptPet,
	r *http.Request,
) error {
	const contentType = "application/json"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := new(jx.Encoder)
	{
		if req.Set {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePetUpdateNameAliasPostRequest(
	req OptPetName,
	r *http.Request,
) error {
	const contentType = "application/json"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := new(jx.Encoder)
	{
		if req.Set {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePetUpdateNamePostRequest(
	req OptString,
	r *http.Request,
) error {
	const contentType = "application/json"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := new(jx.Encoder)
	{
		if req.Set {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePetUploadAvatarByIDRequest(
	req PetUploadAvatarByIDReq,
	r *http.Request,
) error {
	const contentType = "application/octet-stream"
	body := req.Data
	ht.SetBody(r, body, contentType)
	return nil
}

func encodeTestFloatValidationRequest(
	req *TestFloatValidation,
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
