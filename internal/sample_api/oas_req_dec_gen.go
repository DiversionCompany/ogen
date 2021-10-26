// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

func decodeFoobarPostRequest(r *http.Request) (req Pet, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			// Struct or enum.
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := request.Validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetCreateRequest(r *http.Request) (req PetCreateReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			// Struct or enum.
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := request.Validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return &request, nil
	case "text/plain":
		var request PetCreateReqTextPlain
		_ = request
		return req, fmt.Errorf("text/plain decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}
