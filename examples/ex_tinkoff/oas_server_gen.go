// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
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
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// MarketBondsGet implements  operation.
	//
	// GET /market/bonds
	MarketBondsGet(ctx context.Context) (MarketBondsGetRes, error)
	// MarketCandlesGet implements  operation.
	//
	// GET /market/candles
	MarketCandlesGet(ctx context.Context, params MarketCandlesGetParams) (MarketCandlesGetRes, error)
	// MarketCurrenciesGet implements  operation.
	//
	// GET /market/currencies
	MarketCurrenciesGet(ctx context.Context) (MarketCurrenciesGetRes, error)
	// MarketEtfsGet implements  operation.
	//
	// GET /market/etfs
	MarketEtfsGet(ctx context.Context) (MarketEtfsGetRes, error)
	// MarketOrderbookGet implements  operation.
	//
	// GET /market/orderbook
	MarketOrderbookGet(ctx context.Context, params MarketOrderbookGetParams) (MarketOrderbookGetRes, error)
	// MarketSearchByFigiGet implements  operation.
	//
	// GET /market/search/by-figi
	MarketSearchByFigiGet(ctx context.Context, params MarketSearchByFigiGetParams) (MarketSearchByFigiGetRes, error)
	// MarketSearchByTickerGet implements  operation.
	//
	// GET /market/search/by-ticker
	MarketSearchByTickerGet(ctx context.Context, params MarketSearchByTickerGetParams) (MarketSearchByTickerGetRes, error)
	// MarketStocksGet implements  operation.
	//
	// GET /market/stocks
	MarketStocksGet(ctx context.Context) (MarketStocksGetRes, error)
	// OperationsGet implements  operation.
	//
	// GET /operations
	OperationsGet(ctx context.Context, params OperationsGetParams) (OperationsGetRes, error)
	// OrdersCancelPost implements  operation.
	//
	// POST /orders/cancel
	OrdersCancelPost(ctx context.Context, params OrdersCancelPostParams) (OrdersCancelPostRes, error)
	// OrdersGet implements  operation.
	//
	// GET /orders
	OrdersGet(ctx context.Context, params OrdersGetParams) (OrdersGetRes, error)
	// OrdersLimitOrderPost implements  operation.
	//
	// POST /orders/limit-order
	OrdersLimitOrderPost(ctx context.Context, req LimitOrderRequest, params OrdersLimitOrderPostParams) (OrdersLimitOrderPostRes, error)
	// OrdersMarketOrderPost implements  operation.
	//
	// POST /orders/market-order
	OrdersMarketOrderPost(ctx context.Context, req MarketOrderRequest, params OrdersMarketOrderPostParams) (OrdersMarketOrderPostRes, error)
	// PortfolioCurrenciesGet implements  operation.
	//
	// GET /portfolio/currencies
	PortfolioCurrenciesGet(ctx context.Context, params PortfolioCurrenciesGetParams) (PortfolioCurrenciesGetRes, error)
	// PortfolioGet implements  operation.
	//
	// GET /portfolio
	PortfolioGet(ctx context.Context, params PortfolioGetParams) (PortfolioGetRes, error)
	// SandboxClearPost implements  operation.
	//
	// Удаление всех позиций клиента.
	//
	// POST /sandbox/clear
	SandboxClearPost(ctx context.Context, params SandboxClearPostParams) (SandboxClearPostRes, error)
	// SandboxCurrenciesBalancePost implements  operation.
	//
	// POST /sandbox/currencies/balance
	SandboxCurrenciesBalancePost(ctx context.Context, req SandboxSetCurrencyBalanceRequest, params SandboxCurrenciesBalancePostParams) (SandboxCurrenciesBalancePostRes, error)
	// SandboxPositionsBalancePost implements  operation.
	//
	// POST /sandbox/positions/balance
	SandboxPositionsBalancePost(ctx context.Context, req SandboxSetPositionBalanceRequest, params SandboxPositionsBalancePostParams) (SandboxPositionsBalancePostRes, error)
	// SandboxRegisterPost implements  operation.
	//
	// Создание счета и валютных позиций для клиента.
	//
	// POST /sandbox/register
	SandboxRegisterPost(ctx context.Context, req OptSandboxRegisterRequest) (SandboxRegisterPostRes, error)
	// SandboxRemovePost implements  operation.
	//
	// Удаление счета клиента.
	//
	// POST /sandbox/remove
	SandboxRemovePost(ctx context.Context, params SandboxRemovePostParams) (SandboxRemovePostRes, error)
	// UserAccountsGet implements  operation.
	//
	// GET /user/accounts
	UserAccountsGet(ctx context.Context) (UserAccountsGetRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config
}

func NewServer(h Handler, opts ...Option) *Server {
	srv := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	return srv
}
