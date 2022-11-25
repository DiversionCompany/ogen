// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// PublishEvent implements publishEvent operation.
	//
	// POST /event
	PublishEvent(ctx context.Context, req OptEvent) (*Event, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}

// WebhookHandler handles webhooks described by OpenAPI v3 specification.
type WebhookHandler interface {
	// StatusWebhook implements statusWebhook operation.
	//
	StatusWebhook(ctx context.Context) (*StatusWebhookOK, error)
	// UpdateDelete implements DELETE update operation.
	//
	UpdateDelete(ctx context.Context) (UpdateDeleteRes, error)
	// UpdateWebhook implements updateWebhook operation.
	//
	UpdateWebhook(ctx context.Context, req OptEvent, params UpdateWebhookParams) (UpdateWebhookRes, error)
}

// WebhookServer implements http server based on OpenAPI v3 specification and
// calls WebhookHandler to handle requests.
type WebhookServer struct {
	h WebhookHandler
	baseServer
}

// NewWebhookServer creates new WebhookServer.
func NewWebhookServer(h WebhookHandler, opts ...ServerOption) (*WebhookServer, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &WebhookServer{
		h:          h,
		baseServer: s,
	}, nil
}
