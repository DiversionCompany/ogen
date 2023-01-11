// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// APIKey provides api_key security value.
	APIKey(ctx context.Context, operationName string) (APIKey, error)
}

func (s *Client) securityAPIKey(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.APIKey(ctx, operationName)
	if err != nil {
		if errors.Is(err, ogenerrors.ErrSkipClientSecurity) {
			return ogenerrors.ErrSkipClientSecurity
		}
		return errors.Wrap(err, "security source \"APIKey\"")
	}
	req.Header.Set("api_key", t.APIKey)
	return nil
}
