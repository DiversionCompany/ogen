// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"strings"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
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
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "anyContentTypeBinaryStringSchema"
				if l := len("anyContentTypeBinaryStringSchema"); len(elem) >= l && elem[0:l] == "anyContentTypeBinaryStringSchema" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleAnyContentTypeBinaryStringSchemaRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
				switch elem[0] {
				case 'D': // Prefix: "Default"
					if l := len("Default"); len(elem) >= l && elem[0:l] == "Default" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleAnyContentTypeBinaryStringSchemaDefaultRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'c': // Prefix: "combined"
				if l := len("combined"); len(elem) >= l && elem[0:l] == "combined" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleCombinedRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'h': // Prefix: "headers"
				if l := len("headers"); len(elem) >= l && elem[0:l] == "headers" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '2': // Prefix: "200"
					if l := len("200"); len(elem) >= l && elem[0:l] == "200" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleHeaders200Request([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'C': // Prefix: "Combined"
					if l := len("Combined"); len(elem) >= l && elem[0:l] == "Combined" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleHeadersCombinedRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'D': // Prefix: "Default"
					if l := len("Default"); len(elem) >= l && elem[0:l] == "Default" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleHeadersDefaultRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'P': // Prefix: "Pattern"
					if l := len("Pattern"); len(elem) >= l && elem[0:l] == "Pattern" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleHeadersPatternRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'i': // Prefix: "intersectPatternCode"
				if l := len("intersectPatternCode"); len(elem) >= l && elem[0:l] == "intersectPatternCode" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleIntersectPatternCodeRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'm': // Prefix: "multipleGenericResponses"
				if l := len("multipleGenericResponses"); len(elem) >= l && elem[0:l] == "multipleGenericResponses" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleMultipleGenericResponsesRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'o': // Prefix: "octetStream"
				if l := len("octetStream"); len(elem) >= l && elem[0:l] == "octetStream" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'B': // Prefix: "BinaryStringSchema"
					if l := len("BinaryStringSchema"); len(elem) >= l && elem[0:l] == "BinaryStringSchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleOctetStreamBinaryStringSchemaRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'E': // Prefix: "EmptySchema"
					if l := len("EmptySchema"); len(elem) >= l && elem[0:l] == "EmptySchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleOctetStreamEmptySchemaRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 't': // Prefix: "textPlainBinaryStringSchema"
				if l := len("textPlainBinaryStringSchema"); len(elem) >= l && elem[0:l] == "textPlainBinaryStringSchema" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleTextPlainBinaryStringSchemaRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
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
	if elem == "" {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
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
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "anyContentTypeBinaryStringSchema"
				if l := len("anyContentTypeBinaryStringSchema"); len(elem) >= l && elem[0:l] == "anyContentTypeBinaryStringSchema" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "AnyContentTypeBinaryStringSchema"
						r.operationID = "anyContentTypeBinaryStringSchema"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case 'D': // Prefix: "Default"
					if l := len("Default"); len(elem) >= l && elem[0:l] == "Default" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: AnyContentTypeBinaryStringSchemaDefault
							r.name = "AnyContentTypeBinaryStringSchemaDefault"
							r.operationID = "anyContentTypeBinaryStringSchemaDefault"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'c': // Prefix: "combined"
				if l := len("combined"); len(elem) >= l && elem[0:l] == "combined" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: Combined
						r.name = "Combined"
						r.operationID = "combined"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'h': // Prefix: "headers"
				if l := len("headers"); len(elem) >= l && elem[0:l] == "headers" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '2': // Prefix: "200"
					if l := len("200"); len(elem) >= l && elem[0:l] == "200" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: Headers200
							r.name = "Headers200"
							r.operationID = "headers200"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'C': // Prefix: "Combined"
					if l := len("Combined"); len(elem) >= l && elem[0:l] == "Combined" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: HeadersCombined
							r.name = "HeadersCombined"
							r.operationID = "headersCombined"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'D': // Prefix: "Default"
					if l := len("Default"); len(elem) >= l && elem[0:l] == "Default" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: HeadersDefault
							r.name = "HeadersDefault"
							r.operationID = "headersDefault"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'P': // Prefix: "Pattern"
					if l := len("Pattern"); len(elem) >= l && elem[0:l] == "Pattern" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: HeadersPattern
							r.name = "HeadersPattern"
							r.operationID = "headersPattern"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'i': // Prefix: "intersectPatternCode"
				if l := len("intersectPatternCode"); len(elem) >= l && elem[0:l] == "intersectPatternCode" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: IntersectPatternCode
						r.name = "IntersectPatternCode"
						r.operationID = "intersectPatternCode"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'm': // Prefix: "multipleGenericResponses"
				if l := len("multipleGenericResponses"); len(elem) >= l && elem[0:l] == "multipleGenericResponses" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: MultipleGenericResponses
						r.name = "MultipleGenericResponses"
						r.operationID = "multipleGenericResponses"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'o': // Prefix: "octetStream"
				if l := len("octetStream"); len(elem) >= l && elem[0:l] == "octetStream" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'B': // Prefix: "BinaryStringSchema"
					if l := len("BinaryStringSchema"); len(elem) >= l && elem[0:l] == "BinaryStringSchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: OctetStreamBinaryStringSchema
							r.name = "OctetStreamBinaryStringSchema"
							r.operationID = "octetStreamBinaryStringSchema"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'E': // Prefix: "EmptySchema"
					if l := len("EmptySchema"); len(elem) >= l && elem[0:l] == "EmptySchema" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: OctetStreamEmptySchema
							r.name = "OctetStreamEmptySchema"
							r.operationID = "octetStreamEmptySchema"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 't': // Prefix: "textPlainBinaryStringSchema"
				if l := len("textPlainBinaryStringSchema"); len(elem) >= l && elem[0:l] == "textPlainBinaryStringSchema" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: TextPlainBinaryStringSchema
						r.name = "TextPlainBinaryStringSchema"
						r.operationID = "textPlainBinaryStringSchema"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			}
		}
	}
	return r, false
}
