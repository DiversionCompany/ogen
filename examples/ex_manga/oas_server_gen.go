// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GetBook implements getBook operation.
	//
	// Gets metadata of book.
	//
	// GET /api/gallery/{book_id}
	GetBook(ctx context.Context, params GetBookParams) (GetBookRes, error)
	// GetPageCoverImage implements getPageCoverImage operation.
	//
	// Gets page cover.
	//
	// GET /galleries/{media_id}/cover.{format}
	GetPageCoverImage(ctx context.Context, params GetPageCoverImageParams) (GetPageCoverImageRes, error)
	// GetPageImage implements getPageImage operation.
	//
	// Gets page.
	//
	// GET /galleries/{media_id}/{page}.{format}
	GetPageImage(ctx context.Context, params GetPageImageParams) (GetPageImageRes, error)
	// GetPageThumbnailImage implements getPageThumbnailImage operation.
	//
	// Gets page thumbnail.
	//
	// GET /galleries/{media_id}/{page}t.{format}
	GetPageThumbnailImage(ctx context.Context, params GetPageThumbnailImageParams) (GetPageThumbnailImageRes, error)
	// Search implements search operation.
	//
	// Search for comics.
	//
	// GET /api/galleries/search
	Search(ctx context.Context, params SearchParams) (SearchRes, error)
	// SearchByTagID implements searchByTagID operation.
	//
	// Search for comics by tag ID.
	//
	// GET /api/galleries/tagged
	SearchByTagID(ctx context.Context, params SearchByTagIDParams) (SearchByTagIDRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...Option) (*Server, error) {
	s, err := newConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
