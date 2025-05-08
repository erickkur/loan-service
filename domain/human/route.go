package human

import (
	"net/http"

	ma "github.com/loan-service/adapter/middleware"
	"github.com/loan-service/application/services/human"
	rs "github.com/loan-service/application/services/router"
)

type RouteDependency struct {
	HumanService human.HumanServiceInterface

	Context rs.Context
}

type HumanRoute struct {
	HumanService human.HumanServiceInterface

	Context rs.Context
}

func NewDomain(d RouteDependency) {
	route := HumanRoute(d)

	route.InitEndpoints()
}

func (r HumanRoute) InitEndpoints() {
	h := NewHandler(HandlerDependency{
		Human:  r.HumanService,
		Helper: r.Context.Helper,
	})

	r.Context.RegisterEndpoint(r.GetHumans(h))
}

func (r HumanRoute) GetHumans(h Handler) rs.EndpointInfo {
	return rs.EndpointInfo{
		HTTPMethod: http.MethodGet,
		URLPattern: "/humans",
		Handler:    h.GetHumans(),
		Verifications: []ma.VerificationType{
			ma.VerificationTypeConstants.AppToken,
		},
	}
}
