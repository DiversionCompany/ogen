// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// HealthzGet implements GET /healthz operation.
//
// GET /healthz
func (UnimplementedHandler) HealthzGet(ctx context.Context, params HealthzGetParams) (r *HealthzGetOK, _ error) {
	return r, ht.ErrNotImplemented
}

// SameName implements sameName operation.
//
// Parameter with different location, but the same name.
//
// GET /same_name/{path}
func (UnimplementedHandler) SameName(ctx context.Context, params SameNameParams) (r *SameNameOK, _ error) {
	return r, ht.ErrNotImplemented
}
