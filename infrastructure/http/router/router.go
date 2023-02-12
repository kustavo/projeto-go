package router

import (
	nethttp "net/http"

	"github.com/kustavo/benchmark/go/application/interfaces"
	"github.com/kustavo/benchmark/go/infrastructure/http/middleware"
)

type Handler func(w nethttp.ResponseWriter, r *nethttp.Request)

type route struct {
	url         string
	method      string
	handler     nethttp.HandlerFunc
	requireAuth bool
}

type Router struct {
	routes         map[string]route
	authentication interfaces.Authentication
}

func NewRouter(authentication interfaces.Authentication) *Router {
	return &Router{
		routes:         make(map[string]route),
		authentication: authentication,
	}
}

func (r *Router) AddRoutes(routes []route) {
	for _, route := range routes {
		r.routes[route.url+"#"+route.method] = route
	}
}

func (r *Router) AddRoute(method string, url string, handler nethttp.HandlerFunc, requireAuth bool) {
	r.routes[url+"#"+method] = route{method: method, url: url, handler: handler, requireAuth: requireAuth}
}

func (r *Router) ServeHTTP(w nethttp.ResponseWriter, req *nethttp.Request) {
	url := req.URL.Path
	method := req.Method

	route := r.routes[url+"#"+method]

	var handler nethttp.HandlerFunc

	if route.requireAuth {
		handler = middleware.AuthMiddleware(r.authentication, route.handler)
	} else {
		handler = route.handler
	}

	middleware.CorsMiddleware(handler).ServeHTTP(w, req)
}
