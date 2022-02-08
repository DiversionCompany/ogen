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

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleOperationsGetRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'm': // Prefix: "market/"
				if l := len("market/"); len(elem) >= l && elem[0:l] == "market/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleMarketCandlesGetRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case 'b': // Prefix: "bonds"
					if l := len("bonds"); len(elem) >= l && elem[0:l] == "bonds" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: MarketBondsGet
						s.handleMarketBondsGetRequest([0]string{}, w, r)

						return
					}
				case 'c': // Prefix: "c"
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleMarketCurrenciesGetRequest([0]string{}, w, r)

						return
					}
					switch elem[0] {
					case 'a': // Prefix: "andles"
						if l := len("andles"); len(elem) >= l && elem[0:l] == "andles" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: MarketCandlesGet
							s.handleMarketCandlesGetRequest([0]string{}, w, r)

							return
						}
					case 'u': // Prefix: "urrencies"
						if l := len("urrencies"); len(elem) >= l && elem[0:l] == "urrencies" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: MarketCurrenciesGet
							s.handleMarketCurrenciesGetRequest([0]string{}, w, r)

							return
						}
					}
				case 'e': // Prefix: "etfs"
					if l := len("etfs"); len(elem) >= l && elem[0:l] == "etfs" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: MarketEtfsGet
						s.handleMarketEtfsGetRequest([0]string{}, w, r)

						return
					}
				case 'o': // Prefix: "orderbook"
					if l := len("orderbook"); len(elem) >= l && elem[0:l] == "orderbook" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: MarketOrderbookGet
						s.handleMarketOrderbookGetRequest([0]string{}, w, r)

						return
					}
				case 's': // Prefix: "s"
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleMarketStocksGetRequest([0]string{}, w, r)

						return
					}
					switch elem[0] {
					case 'e': // Prefix: "earch/by-"
						if l := len("earch/by-"); len(elem) >= l && elem[0:l] == "earch/by-" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleMarketSearchByTickerGetRequest([0]string{}, w, r)

							return
						}
						switch elem[0] {
						case 'f': // Prefix: "figi"
							if l := len("figi"); len(elem) >= l && elem[0:l] == "figi" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: MarketSearchByFigiGet
								s.handleMarketSearchByFigiGetRequest([0]string{}, w, r)

								return
							}
						case 't': // Prefix: "ticker"
							if l := len("ticker"); len(elem) >= l && elem[0:l] == "ticker" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: MarketSearchByTickerGet
								s.handleMarketSearchByTickerGetRequest([0]string{}, w, r)

								return
							}
						}
					case 't': // Prefix: "tocks"
						if l := len("tocks"); len(elem) >= l && elem[0:l] == "tocks" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: MarketStocksGet
							s.handleMarketStocksGetRequest([0]string{}, w, r)

							return
						}
					}
				}
			case 'o': // Prefix: "o"
				if l := len("o"); len(elem) >= l && elem[0:l] == "o" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleOrdersGetRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case 'p': // Prefix: "perations"
					if l := len("perations"); len(elem) >= l && elem[0:l] == "perations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OperationsGet
						s.handleOperationsGetRequest([0]string{}, w, r)

						return
					}
				case 'r': // Prefix: "rders"
					if l := len("rders"); len(elem) >= l && elem[0:l] == "rders" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OrdersGet
						s.handleOrdersGetRequest([0]string{}, w, r)

						return
					}
				}
			case 'p': // Prefix: "portfolio"
				if l := len("portfolio"); len(elem) >= l && elem[0:l] == "portfolio" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handlePortfolioGetRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/currencies"
					if l := len("/currencies"); len(elem) >= l && elem[0:l] == "/currencies" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: PortfolioCurrenciesGet
						s.handlePortfolioCurrenciesGetRequest([0]string{}, w, r)

						return
					}
				}
			case 'u': // Prefix: "user/accounts"
				if l := len("user/accounts"); len(elem) >= l && elem[0:l] == "user/accounts" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: UserAccountsGet
					s.handleUserAccountsGetRequest([0]string{}, w, r)

					return
				}
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleSandboxClearPostRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'o': // Prefix: "orders/"
				if l := len("orders/"); len(elem) >= l && elem[0:l] == "orders/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleOrdersLimitOrderPostRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case 'c': // Prefix: "cancel"
					if l := len("cancel"); len(elem) >= l && elem[0:l] == "cancel" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OrdersCancelPost
						s.handleOrdersCancelPostRequest([0]string{}, w, r)

						return
					}
				case 'l': // Prefix: "limit-order"
					if l := len("limit-order"); len(elem) >= l && elem[0:l] == "limit-order" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OrdersLimitOrderPost
						s.handleOrdersLimitOrderPostRequest([0]string{}, w, r)

						return
					}
				case 'm': // Prefix: "market-order"
					if l := len("market-order"); len(elem) >= l && elem[0:l] == "market-order" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OrdersMarketOrderPost
						s.handleOrdersMarketOrderPostRequest([0]string{}, w, r)

						return
					}
				}
			case 's': // Prefix: "sandbox/"
				if l := len("sandbox/"); len(elem) >= l && elem[0:l] == "sandbox/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleSandboxPositionsBalancePostRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case 'c': // Prefix: "c"
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleSandboxCurrenciesBalancePostRequest([0]string{}, w, r)

						return
					}
					switch elem[0] {
					case 'l': // Prefix: "lear"
						if l := len("lear"); len(elem) >= l && elem[0:l] == "lear" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SandboxClearPost
							s.handleSandboxClearPostRequest([0]string{}, w, r)

							return
						}
					case 'u': // Prefix: "urrencies/balance"
						if l := len("urrencies/balance"); len(elem) >= l && elem[0:l] == "urrencies/balance" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SandboxCurrenciesBalancePost
							s.handleSandboxCurrenciesBalancePostRequest([0]string{}, w, r)

							return
						}
					}
				case 'p': // Prefix: "positions/balance"
					if l := len("positions/balance"); len(elem) >= l && elem[0:l] == "positions/balance" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: SandboxPositionsBalancePost
						s.handleSandboxPositionsBalancePostRequest([0]string{}, w, r)

						return
					}
				case 'r': // Prefix: "re"
					if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleSandboxRemovePostRequest([0]string{}, w, r)

						return
					}
					switch elem[0] {
					case 'g': // Prefix: "gister"
						if l := len("gister"); len(elem) >= l && elem[0:l] == "gister" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SandboxRegisterPost
							s.handleSandboxRegisterPostRequest([0]string{}, w, r)

							return
						}
					case 'm': // Prefix: "move"
						if l := len("move"); len(elem) >= l && elem[0:l] == "move" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SandboxRemovePost
							s.handleSandboxRemovePostRequest([0]string{}, w, r)

							return
						}
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name  string
	count int
	args  [0]string
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.name
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [0]string{}
		elem = path
	)
	r.args = args

	// Static code generated router with unwrapped path search.
	switch method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "OperationsGet"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'm': // Prefix: "market/"
				if l := len("market/"); len(elem) >= l && elem[0:l] == "market/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "MarketCandlesGet"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'b': // Prefix: "bonds"
					if l := len("bonds"); len(elem) >= l && elem[0:l] == "bonds" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: MarketBondsGet
						r.name = "MarketBondsGet"
						r.args = args
						r.count = 0
						return r, true
					}
				case 'c': // Prefix: "c"
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						r.name = "MarketCurrenciesGet"
						r.args = args
						r.count = 0
						return r, true
					}
					switch elem[0] {
					case 'a': // Prefix: "andles"
						if l := len("andles"); len(elem) >= l && elem[0:l] == "andles" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: MarketCandlesGet
							r.name = "MarketCandlesGet"
							r.args = args
							r.count = 0
							return r, true
						}
					case 'u': // Prefix: "urrencies"
						if l := len("urrencies"); len(elem) >= l && elem[0:l] == "urrencies" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: MarketCurrenciesGet
							r.name = "MarketCurrenciesGet"
							r.args = args
							r.count = 0
							return r, true
						}
					}
				case 'e': // Prefix: "etfs"
					if l := len("etfs"); len(elem) >= l && elem[0:l] == "etfs" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: MarketEtfsGet
						r.name = "MarketEtfsGet"
						r.args = args
						r.count = 0
						return r, true
					}
				case 'o': // Prefix: "orderbook"
					if l := len("orderbook"); len(elem) >= l && elem[0:l] == "orderbook" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: MarketOrderbookGet
						r.name = "MarketOrderbookGet"
						r.args = args
						r.count = 0
						return r, true
					}
				case 's': // Prefix: "s"
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						r.name = "MarketStocksGet"
						r.args = args
						r.count = 0
						return r, true
					}
					switch elem[0] {
					case 'e': // Prefix: "earch/by-"
						if l := len("earch/by-"); len(elem) >= l && elem[0:l] == "earch/by-" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							r.name = "MarketSearchByTickerGet"
							r.args = args
							r.count = 0
							return r, true
						}
						switch elem[0] {
						case 'f': // Prefix: "figi"
							if l := len("figi"); len(elem) >= l && elem[0:l] == "figi" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: MarketSearchByFigiGet
								r.name = "MarketSearchByFigiGet"
								r.args = args
								r.count = 0
								return r, true
							}
						case 't': // Prefix: "ticker"
							if l := len("ticker"); len(elem) >= l && elem[0:l] == "ticker" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: MarketSearchByTickerGet
								r.name = "MarketSearchByTickerGet"
								r.args = args
								r.count = 0
								return r, true
							}
						}
					case 't': // Prefix: "tocks"
						if l := len("tocks"); len(elem) >= l && elem[0:l] == "tocks" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: MarketStocksGet
							r.name = "MarketStocksGet"
							r.args = args
							r.count = 0
							return r, true
						}
					}
				}
			case 'o': // Prefix: "o"
				if l := len("o"); len(elem) >= l && elem[0:l] == "o" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "OrdersGet"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'p': // Prefix: "perations"
					if l := len("perations"); len(elem) >= l && elem[0:l] == "perations" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OperationsGet
						r.name = "OperationsGet"
						r.args = args
						r.count = 0
						return r, true
					}
				case 'r': // Prefix: "rders"
					if l := len("rders"); len(elem) >= l && elem[0:l] == "rders" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OrdersGet
						r.name = "OrdersGet"
						r.args = args
						r.count = 0
						return r, true
					}
				}
			case 'p': // Prefix: "portfolio"
				if l := len("portfolio"); len(elem) >= l && elem[0:l] == "portfolio" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "PortfolioGet"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/currencies"
					if l := len("/currencies"); len(elem) >= l && elem[0:l] == "/currencies" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: PortfolioCurrenciesGet
						r.name = "PortfolioCurrenciesGet"
						r.args = args
						r.count = 0
						return r, true
					}
				}
			case 'u': // Prefix: "user/accounts"
				if l := len("user/accounts"); len(elem) >= l && elem[0:l] == "user/accounts" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: UserAccountsGet
					r.name = "UserAccountsGet"
					r.args = args
					r.count = 0
					return r, true
				}
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "SandboxClearPost"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'o': // Prefix: "orders/"
				if l := len("orders/"); len(elem) >= l && elem[0:l] == "orders/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "OrdersLimitOrderPost"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'c': // Prefix: "cancel"
					if l := len("cancel"); len(elem) >= l && elem[0:l] == "cancel" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OrdersCancelPost
						r.name = "OrdersCancelPost"
						r.args = args
						r.count = 0
						return r, true
					}
				case 'l': // Prefix: "limit-order"
					if l := len("limit-order"); len(elem) >= l && elem[0:l] == "limit-order" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OrdersLimitOrderPost
						r.name = "OrdersLimitOrderPost"
						r.args = args
						r.count = 0
						return r, true
					}
				case 'm': // Prefix: "market-order"
					if l := len("market-order"); len(elem) >= l && elem[0:l] == "market-order" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: OrdersMarketOrderPost
						r.name = "OrdersMarketOrderPost"
						r.args = args
						r.count = 0
						return r, true
					}
				}
			case 's': // Prefix: "sandbox/"
				if l := len("sandbox/"); len(elem) >= l && elem[0:l] == "sandbox/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "SandboxPositionsBalancePost"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'c': // Prefix: "c"
					if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						r.name = "SandboxCurrenciesBalancePost"
						r.args = args
						r.count = 0
						return r, true
					}
					switch elem[0] {
					case 'l': // Prefix: "lear"
						if l := len("lear"); len(elem) >= l && elem[0:l] == "lear" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SandboxClearPost
							r.name = "SandboxClearPost"
							r.args = args
							r.count = 0
							return r, true
						}
					case 'u': // Prefix: "urrencies/balance"
						if l := len("urrencies/balance"); len(elem) >= l && elem[0:l] == "urrencies/balance" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SandboxCurrenciesBalancePost
							r.name = "SandboxCurrenciesBalancePost"
							r.args = args
							r.count = 0
							return r, true
						}
					}
				case 'p': // Prefix: "positions/balance"
					if l := len("positions/balance"); len(elem) >= l && elem[0:l] == "positions/balance" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: SandboxPositionsBalancePost
						r.name = "SandboxPositionsBalancePost"
						r.args = args
						r.count = 0
						return r, true
					}
				case 'r': // Prefix: "re"
					if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						r.name = "SandboxRemovePost"
						r.args = args
						r.count = 0
						return r, true
					}
					switch elem[0] {
					case 'g': // Prefix: "gister"
						if l := len("gister"); len(elem) >= l && elem[0:l] == "gister" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SandboxRegisterPost
							r.name = "SandboxRegisterPost"
							r.args = args
							r.count = 0
							return r, true
						}
					case 'm': // Prefix: "move"
						if l := len("move"); len(elem) >= l && elem[0:l] == "move" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SandboxRemovePost
							r.name = "SandboxRemovePost"
							r.args = args
							r.count = 0
							return r, true
						}
					}
				}
			}
		}
	}
	return r, false
}
