package lenderloan

import (
	"net/http"

	rs "github.com/loan-service/application/services/router"
	cs "github.com/loan-service/internal/constant"
	"github.com/loan-service/internal/logger"

	llService "github.com/loan-service/application/services/lenderloan"
	lService "github.com/loan-service/application/services/loan"
	llogService "github.com/loan-service/application/services/loanlog"
)

type RouteDependency struct {
	Context           rs.Context
	Logger            logger.Interface
	LenderLoanService llService.LenderLoanServiceInterface
	LoanService       lService.LoanServiceInterface
	LoanLogService    llogService.LoanLogServiceInterface
}

type Route struct {
	Context           rs.Context
	Logger            logger.Interface
	LenderLoanService llService.LenderLoanServiceInterface
	LoanService       lService.LoanServiceInterface
	LoanLogService    llogService.LoanLogServiceInterface
}

func NewDomain(d RouteDependency) {
	route := Route(d)

	route.InitEndpoints()
}

func (r Route) InitEndpoints() {
	h := NewHandler(HandlerDependency{
		Logger:            r.Logger,
		Context:           r.Context,
		LenderLoanService: r.LenderLoanService,
		LoanService:       r.LoanService,
		LoanLogService:    r.LoanLogService,
	})

	r.Context.RegisterEndpoint(r.CreateLenderLoanEndpoint(h))
}

func (r Route) CreateLenderLoanEndpoint(h Handler) rs.EndpointInfo {
	return rs.EndpointInfo{
		HTTPMethod: http.MethodPost,
		URLPattern: "/lender-loans",
		Handler:    h.CreateLenderLoanHandler(),
		Verifications: []cs.VerificationType{
			cs.VerificationTypeConstants.AppToken,
		},
	}
}
