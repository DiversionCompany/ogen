// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
	"github.com/ogen-go/ogen/types"
	"github.com/ogen-go/ogen/uri"
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
	_ = math.Mod
	_ = types.Date{}
	_ = jsoniter.Config{}
)

func NewFoobarGetHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodeFoobarGetParams(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.FoobarGet(r.Context(), params)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeFoobarGetResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewFoobarPutHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := s.FoobarPut(r.Context())
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeFoobarPutResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewFoobarPostHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := decodeFoobarPostRequest(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.FoobarPost(r.Context(), request)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodeFoobarPostResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewPetGetHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodePetGetParams(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PetGet(r.Context(), params)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePetGetResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewPetCreateHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := decodePetCreateRequest(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PetCreate(r.Context(), request)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePetCreateResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func NewPetGetByNameHandler(s Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := decodePetGetByNameParams(r)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		response, err := s.PetGetByName(r.Context(), params)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		if err := encodePetGetByNameResponse(response, w); err != nil {
			_ = err
			return
		}
	}
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}
