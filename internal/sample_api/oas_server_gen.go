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

// Server handles operations described by OpenAPI v3 specification.
type Server interface {
	// FoobarGet implements foobarGet operation.
	FoobarGet(ctx context.Context, params FoobarGetParams) (FoobarGetResponse, error)
	// FoobarPut implements  operation.
	FoobarPut(ctx context.Context) (FoobarPutResponseDefaultStatusCode, error)
	// FoobarPost implements foobarPost operation.
	FoobarPost(ctx context.Context) (FoobarPostResponse, error)
	// PetNameByID implements petNameByID operation.
	PetNameByID(ctx context.Context, params PetNameByIDParams) (string, error)
	// PetFriendsNamesByID implements petFriendsNamesByID operation.
	PetFriendsNamesByID(ctx context.Context, params PetFriendsNamesByIDParams) ([]string, error)
	// PetGet implements petGet operation.
	PetGet(ctx context.Context, params PetGetParams) (PetGetResponse, error)
	// PetCreate implements petCreate operation.
	PetCreate(ctx context.Context) (FoobarGetResponseOKApplicationJSON, error)
	// PetGetByName implements petGetByName operation.
	PetGetByName(ctx context.Context, params PetGetByNameParams) (FoobarGetResponseOKApplicationJSON, error)
}
