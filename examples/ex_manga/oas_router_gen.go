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
	args := [3]string{}
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
				s.handleGetPageCoverImageRequest([2]string{
					args[0],
					args[1],
				}, w, r)

				return
			}
			switch elem[0] {
			case 'a': // Prefix: "api/galler"
				if l := len("api/galler"); len(elem) >= l && elem[0:l] == "api/galler" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleSearchRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case 'i': // Prefix: "ies/"
					if l := len("ies/"); len(elem) >= l && elem[0:l] == "ies/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleSearchByTagIDRequest([0]string{}, w, r)

						return
					}
					switch elem[0] {
					case 's': // Prefix: "search"
						if l := len("search"); len(elem) >= l && elem[0:l] == "search" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: Search
							s.handleSearchRequest([0]string{}, w, r)

							return
						}
					case 't': // Prefix: "tagged"
						if l := len("tagged"); len(elem) >= l && elem[0:l] == "tagged" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SearchByTagID
							s.handleSearchByTagIDRequest([0]string{}, w, r)

							return
						}
					}
				case 'y': // Prefix: "y/"
					if l := len("y/"); len(elem) >= l && elem[0:l] == "y/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "book_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: GetBook
						s.handleGetBookRequest([1]string{
							args[0],
						}, w, r)

						return
					}
				}
			case 'g': // Prefix: "galleries/"
				if l := len("galleries/"); len(elem) >= l && elem[0:l] == "galleries/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "media_id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx > 0 {
					args[0] = elem[:idx]
					elem = elem[idx:]

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
							s.handleGetPageImageRequest([3]string{
								args[0],
								args[1],
								args[2],
							}, w, r)

							return
						}
						switch elem[0] {
						case 'c': // Prefix: "cover."
							if l := len("cover."); len(elem) >= l && elem[0:l] == "cover." {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "format"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf: GetPageCoverImage
								s.handleGetPageCoverImageRequest([2]string{
									args[0],
									args[1],
								}, w, r)

								return
							}
						}
						// Param: "page"
						// Match until one of ".t"
						idx := strings.IndexAny(elem, ".t")
						if idx > 0 {
							args[1] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '.': // Prefix: "."
								if l := len("."); len(elem) >= l && elem[0:l] == "." {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "format"
								// Leaf parameter
								args[2] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf: GetPageImage
									s.handleGetPageImageRequest([3]string{
										args[0],
										args[1],
										args[2],
									}, w, r)

									return
								}
							case 't': // Prefix: "t."
								if l := len("t."); len(elem) >= l && elem[0:l] == "t." {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "format"
								// Leaf parameter
								args[2] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf: GetPageThumbnailImage
									s.handleGetPageThumbnailImageRequest([3]string{
										args[0],
										args[1],
										args[2],
									}, w, r)

									return
								}
							}
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
	args  [3]string
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
		args = [3]string{}
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
				r.name = "GetPageCoverImage"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'a': // Prefix: "api/galler"
				if l := len("api/galler"); len(elem) >= l && elem[0:l] == "api/galler" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "Search"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case 'i': // Prefix: "ies/"
					if l := len("ies/"); len(elem) >= l && elem[0:l] == "ies/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						r.name = "SearchByTagID"
						r.args = args
						r.count = 0
						return r, true
					}
					switch elem[0] {
					case 's': // Prefix: "search"
						if l := len("search"); len(elem) >= l && elem[0:l] == "search" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: Search
							r.name = "Search"
							r.args = args
							r.count = 0
							return r, true
						}
					case 't': // Prefix: "tagged"
						if l := len("tagged"); len(elem) >= l && elem[0:l] == "tagged" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: SearchByTagID
							r.name = "SearchByTagID"
							r.args = args
							r.count = 0
							return r, true
						}
					}
				case 'y': // Prefix: "y/"
					if l := len("y/"); len(elem) >= l && elem[0:l] == "y/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "book_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf: GetBook
						r.name = "GetBook"
						r.args = args
						r.count = 1
						return r, true
					}
				}
			case 'g': // Prefix: "galleries/"
				if l := len("galleries/"); len(elem) >= l && elem[0:l] == "galleries/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "media_id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx > 0 {
					args[0] = elem[:idx]
					elem = elem[idx:]

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
							r.name = "GetPageImage"
							r.args = args
							r.count = 1
							return r, true
						}
						switch elem[0] {
						case 'c': // Prefix: "cover."
							if l := len("cover."); len(elem) >= l && elem[0:l] == "cover." {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "format"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf: GetPageCoverImage
								r.name = "GetPageCoverImage"
								r.args = args
								r.count = 2
								return r, true
							}
						}
						// Param: "page"
						// Match until one of ".t"
						idx := strings.IndexAny(elem, ".t")
						if idx > 0 {
							args[1] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '.': // Prefix: "."
								if l := len("."); len(elem) >= l && elem[0:l] == "." {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "format"
								// Leaf parameter
								args[2] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf: GetPageImage
									r.name = "GetPageImage"
									r.args = args
									r.count = 3
									return r, true
								}
							case 't': // Prefix: "t."
								if l := len("t."); len(elem) >= l && elem[0:l] == "t." {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "format"
								// Leaf parameter
								args[2] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf: GetPageThumbnailImage
									r.name = "GetPageThumbnailImage"
									r.args = args
									r.count = 3
									return r, true
								}
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
