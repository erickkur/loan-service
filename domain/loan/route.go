package loan

import (
	"net/http"

	lService "github.com/loan-service/application/services/loan"
	rs "github.com/loan-service/application/services/router"
	cs "github.com/loan-service/internal/constant"
	"github.com/loan-service/internal/logger"
)

type RouteDependency struct {
	Context     rs.Context
	Logger      logger.Interface
	LoanService lService.LoanServiceInterface
}

type Route struct {
	Context     rs.Context
	Logger      logger.Interface
	LoanService lService.LoanServiceInterface
}

func NewDomain(d RouteDependency) {
	route := Route(d)

	route.InitEndpoints()
}

func (r Route) InitEndpoints() {
	h := NewHandler(HandlerDependency{
		Logger:      r.Logger,
		Context:     r.Context,
		LoanService: r.LoanService,
	})

	r.Context.RegisterEndpoint(r.CreateLoanEndpoint(h))
}

func (r Route) CreateLoanEndpoint(h Handler) rs.EndpointInfo {
	return rs.EndpointInfo{
		HTTPMethod: http.MethodPost,
		URLPattern: "/loans",
		Handler:    h.CreateLoanHandler(),
		Verifications: []cs.VerificationType{
			cs.VerificationTypeConstants.AppToken,
		},
	}
}
