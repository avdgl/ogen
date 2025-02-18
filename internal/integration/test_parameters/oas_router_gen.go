// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
	}
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
	args := [1]string{}

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
			case 'c': // Prefix: "co"
				if l := len("co"); len(elem) >= l && elem[0:l] == "co" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "mplicatedParameterName"
					if l := len("mplicatedParameterName"); len(elem) >= l && elem[0:l] == "mplicatedParameterName" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleComplicatedParameterNameGetRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'n': // Prefix: "ntentParameters/"
					if l := len("ntentParameters/"); len(elem) >= l && elem[0:l] == "ntentParameters/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "path"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleContentParametersRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'o': // Prefix: "okieParameter"
					if l := len("okieParameter"); len(elem) >= l && elem[0:l] == "okieParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleCookieParameterRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'h': // Prefix: "headerParameter"
				if l := len("headerParameter"); len(elem) >= l && elem[0:l] == "headerParameter" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleHeaderParameterRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'o': // Prefix: "object"
				if l := len("object"); len(elem) >= l && elem[0:l] == "object" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'C': // Prefix: "CookieParameter"
					if l := len("CookieParameter"); len(elem) >= l && elem[0:l] == "CookieParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleObjectCookieParameterRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'Q': // Prefix: "QueryParameter"
					if l := len("QueryParameter"); len(elem) >= l && elem[0:l] == "QueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleObjectQueryParameterRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 's': // Prefix: "same_name/"
				if l := len("same_name/"); len(elem) >= l && elem[0:l] == "same_name/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "path"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleSameNameRequest([1]string{
							args[0],
						}, w, r)
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
	pathPattern string
	count       int
	args        [1]string
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

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
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
			case 'c': // Prefix: "co"
				if l := len("co"); len(elem) >= l && elem[0:l] == "co" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'm': // Prefix: "mplicatedParameterName"
					if l := len("mplicatedParameterName"); len(elem) >= l && elem[0:l] == "mplicatedParameterName" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ComplicatedParameterNameGet
							r.name = "ComplicatedParameterNameGet"
							r.operationID = ""
							r.pathPattern = "/complicatedParameterName"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'n': // Prefix: "ntentParameters/"
					if l := len("ntentParameters/"); len(elem) >= l && elem[0:l] == "ntentParameters/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "path"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ContentParameters
							r.name = "ContentParameters"
							r.operationID = "contentParameters"
							r.pathPattern = "/contentParameters/{path}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
				case 'o': // Prefix: "okieParameter"
					if l := len("okieParameter"); len(elem) >= l && elem[0:l] == "okieParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: CookieParameter
							r.name = "CookieParameter"
							r.operationID = "cookieParameter"
							r.pathPattern = "/cookieParameter"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'h': // Prefix: "headerParameter"
				if l := len("headerParameter"); len(elem) >= l && elem[0:l] == "headerParameter" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: HeaderParameter
						r.name = "HeaderParameter"
						r.operationID = "headerParameter"
						r.pathPattern = "/headerParameter"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'o': // Prefix: "object"
				if l := len("object"); len(elem) >= l && elem[0:l] == "object" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'C': // Prefix: "CookieParameter"
					if l := len("CookieParameter"); len(elem) >= l && elem[0:l] == "CookieParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ObjectCookieParameter
							r.name = "ObjectCookieParameter"
							r.operationID = "objectCookieParameter"
							r.pathPattern = "/objectCookieParameter"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'Q': // Prefix: "QueryParameter"
					if l := len("QueryParameter"); len(elem) >= l && elem[0:l] == "QueryParameter" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: ObjectQueryParameter
							r.name = "ObjectQueryParameter"
							r.operationID = "objectQueryParameter"
							r.pathPattern = "/objectQueryParameter"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 's': // Prefix: "same_name/"
				if l := len("same_name/"); len(elem) >= l && elem[0:l] == "same_name/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "path"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: SameName
						r.name = "SameName"
						r.operationID = "sameName"
						r.pathPattern = "/same_name/{path}"
						r.args = args
						r.count = 1
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
