package httpserver

import (
	"net/http"

	"github.com/loan-service/infra/way"
	"github.com/loan-service/internal/logger"
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
	Log    logger.Interface
}

func NewAdapter(a *Adapter) *Router {
	return &Router{
		Server: a.Router,
	}
}
