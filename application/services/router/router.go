package router

import (
	"path"

	"github.com/loan-service/adapter/httpserver"
	ma "github.com/loan-service/adapter/middleware"
	cs "github.com/loan-service/internal/constant"
	"github.com/loan-service/internal/handler"
	"github.com/thoas/go-funk"
)

// Context ...
type Context struct {
	router     httpserver.RouterInterface
	middleware *ma.Middleware
	prefix     string
	Helper     httpserver.HelperInterface
}

type EndpointInfo struct {
	HTTPMethod    string
	URLPattern    string
	Handler       handler.EndpointHandler
	Verifications []cs.VerificationType
}

func NewService(router httpserver.RouterInterface, helper httpserver.HelperInterface, middleware ma.Middleware, prefix string) Context {
	return Context{
		router:     router,
		middleware: &middleware,
		prefix:     prefix,
		Helper:     helper,
	}
}

// RegisterEndpoint ...
func (r *Context) RegisterEndpoint(info EndpointInfo) {
	r.RegisterEndpointWithPrefix(info, r.prefix)
}

// RegisterEndpointWithPrefix ...
func (r *Context) RegisterEndpointWithPrefix(info EndpointInfo, prefix string) {
	m := r.middleware
	urlPattern := getFullURLPattern(info, prefix)

	verificationFns := getVerificationMethod(m, info.Verifications)

	r.router.Handle(info.HTTPMethod, urlPattern, m.Verify(info.Handler, verificationFns...))
}

func getVerificationMethod(m *ma.Middleware, verifications []cs.VerificationType) []ma.MiddlewareFunc {
	return funk.Map(verifications, func(t cs.VerificationType) ma.MiddlewareFunc {
		switch t {
		case cs.VerificationTypeConstants.InternalToolToken:
			return m.InternalToolToken
		default:
			return m.AppToken
		}
	}).([]ma.MiddlewareFunc)
}

func getFullURLPattern(info EndpointInfo, prefix string) string {
	if prefix == "" {
		return info.URLPattern
	}
	return path.Join(prefix, info.URLPattern)
}
