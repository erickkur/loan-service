package human

import (
	"net/http"

	"github.com/loan-service/application/services/human"
	rs "github.com/loan-service/application/services/router"
	cs "github.com/loan-service/internal/constant"
	"github.com/loan-service/internal/logger"
)

type RouteDependency struct {
	HumanService human.HumanServiceInterface
	Logger       logger.Interface

	Context rs.Context
}

type HumanRoute struct {
	HumanService human.HumanServiceInterface
	Logger       logger.Interface

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
		Logger: r.Logger,
	})

	r.Context.RegisterEndpoint(r.GetHumans(h))
}

func (r HumanRoute) GetHumans(h Handler) rs.EndpointInfo {
	return rs.EndpointInfo{
		HTTPMethod: http.MethodGet,
		URLPattern: "/humans",
		Handler:    h.GetHumans(),
		Verifications: []cs.VerificationType{
			cs.VerificationTypeConstants.AppToken,
		},
	}
}
