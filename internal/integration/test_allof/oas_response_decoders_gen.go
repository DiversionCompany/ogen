// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/ogen-go/ogen/validate"
)

func decodeNullableStringsResponse(resp *http.Response) (res *NullableStringsOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &NullableStringsOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeObjectsWithConflictingArrayPropertyResponse(resp *http.Response) (res *ObjectsWithConflictingArrayPropertyOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &ObjectsWithConflictingArrayPropertyOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeObjectsWithConflictingPropertiesResponse(resp *http.Response) (res *ObjectsWithConflictingPropertiesOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &ObjectsWithConflictingPropertiesOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeReferencedAllofResponse(resp *http.Response) (res *ReferencedAllofOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &ReferencedAllofOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeReferencedAllofOptionalResponse(resp *http.Response) (res *ReferencedAllofOptionalOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &ReferencedAllofOptionalOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeSimpleIntegerResponse(resp *http.Response) (res *SimpleIntegerOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &SimpleIntegerOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeSimpleObjectsResponse(resp *http.Response) (res *SimpleObjectsOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &SimpleObjectsOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeStringsNotypeResponse(resp *http.Response) (res *StringsNotypeOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &StringsNotypeOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
