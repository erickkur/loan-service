package way

import (
	"context"
	"net/http"
	"sort"
	"strings"
)

// wayContextKey is the context key type for storing
// parameters in context.Context.
type wayContextKey string

type route struct {
	method  string
	segs    []string
	handler http.Handler
	prefix  bool
}

// Router routes HTTP requests.
type Router struct {
	routes []*route
	// NotFound is the http.Handler to call when no routes
	// match. By default uses http.NotFoundHandler().
	NotFound http.Handler
}

// NewRouter makes a new Router.
func NewRouter() *Router {
	return &Router{
		NotFound: http.NotFoundHandler(),
	}
}

func (r *Router) pathSegments(p string) []string {
	return strings.Split(strings.Trim(p, "/"), "/")
}

// Handle adds a handler with the specified method and pattern.
// Method can be any HTTP method string or "*" to match all methods.
// Pattern can contain path segments such as: /item/:id which is
// accessible via the Param function.
// If pattern ends with trailing /, it acts as a prefix.
func (r *Router) Handle(method, pattern string, handler http.Handler) {
	route := &route{
		method:  strings.ToLower(method),
		segs:    r.pathSegments(pattern),
		handler: handler,
		prefix:  strings.HasSuffix(pattern, "/") || strings.HasSuffix(pattern, "..."),
	}
	r.routes = append(r.routes, route)
}

// ServeHTTP routes the incoming http.Request based on method and path
// extracting path parameters as it goes.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := strings.ToLower(req.Method)
	segs := r.pathSegments(req.URL.Path)
	for _, route := range r.routes {
		if route.method != method && route.method != "*" {
			continue
		}
		if ctx, ok := route.match(req.Context(), r, segs); ok {
			route.handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}
	}
	r.NotFound.ServeHTTP(w, req)
}

// Param gets the path parameter from the specified Context.
// Returns an empty string if the parameter was not found.
func Param(ctx context.Context, param string) string {
	vStr, ok := ctx.Value(wayContextKey(param)).(string)
	if !ok {
		return ""
	}
	return vStr
}

// ReArrange all registered routes within router in descending order.
// To avoid router conflict if has wildcard params in segments.
func (r *Router) ReArrange() {
	var nonWildcard []*route
	var withWildcard []*route

	for _, r := range r.routes {
		path := strings.Join(r.segs, "/")
		hasWildcard := strings.Contains(path, ":")
		if hasWildcard {
			withWildcard = append(withWildcard, r)
		} else {
			nonWildcard = append(nonWildcard, r)
		}
	}

	sort.SliceStable(nonWildcard, func(next, prev int) bool {
		nstr := strings.Join(nonWildcard[next].segs, "/")
		pstr := strings.Join(nonWildcard[prev].segs, "/")
		return nstr > pstr
	})

	sort.SliceStable(withWildcard, func(next, prev int) bool {
		if len(withWildcard[next].segs) > len(withWildcard[prev].segs) {
			return len(withWildcard[next].segs) > len(withWildcard[prev].segs)
		}

		nstr := strings.Join(withWildcard[next].segs, "/")
		pstr := strings.Join(withWildcard[prev].segs, "/")

		return strings.Count(nstr, ":") < strings.Count(pstr, ":")
	})

	newRoutes := []*route{}
	newRoutes = append(newRoutes, nonWildcard...)
	newRoutes = append(newRoutes, withWildcard...)

	r.routes = newRoutes
}

func (r *route) match(ctx context.Context, router *Router, segs []string) (context.Context, bool) {
	if len(segs) > len(r.segs) && !r.prefix {
		return nil, false
	}

	for i, seg := range r.segs {
		if i > len(segs)-1 {
			return nil, false
		}

		isParam := false
		if strings.HasPrefix(seg, ":") {
			isParam = true
			seg = strings.TrimPrefix(seg, ":")
		}

		if !isParam { // verbatim check
			if strings.HasSuffix(seg, "...") {
				if strings.HasPrefix(segs[i], seg[:len(seg)-3]) {
					return ctx, true
				}
			}
			if seg != segs[i] {
				return nil, false
			}
		}

		if isParam {
			ctx = context.WithValue(ctx, wayContextKey(seg), segs[i])
		}
	}

	return ctx, true
}
