package httpserver

import (
	"net/http"

	"github.com/loan-service/infra/way"
)

type RouterInterface interface {
	Handle(method string, pattern string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type Router struct {
	Server *way.Router
}

// Adapter ...
type Adapter struct {
	Router *way.Router
}

func NewAdapter(a *Adapter) *Router {
	return &Router{
		Server: a.Router,
	}
}
