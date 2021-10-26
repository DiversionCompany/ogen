// Code generated by ogen, DO NOT EDIT.

package techempower

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
	// DB implements DB operation.
	DB(ctx context.Context) (DBResponseOKApplicationJSON, error)
	// Queries implements Queries operation.
	Queries(ctx context.Context, params QueriesParams) (QueriesResponseOKApplicationJSON, error)
	// Updates implements Updates operation.
	Updates(ctx context.Context, params UpdatesParams) (QueriesResponseOKApplicationJSON, error)
	// Caching implements Caching operation.
	Caching(ctx context.Context, params CachingParams) (QueriesResponseOKApplicationJSON, error)
	// JSON implements json operation.
	JSON(ctx context.Context) (JSONResponseOKApplicationJSON, error)
}
