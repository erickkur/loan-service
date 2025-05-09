package lenderloan

import (
	"net/http"

	"github.com/loan-service/application/dto"
	llService "github.com/loan-service/application/services/lenderloan"
	lService "github.com/loan-service/application/services/loan"
	llogService "github.com/loan-service/application/services/loanlog"
	rs "github.com/loan-service/application/services/router"
	errs "github.com/loan-service/internal/error"
	"github.com/loan-service/internal/handler"
	"github.com/loan-service/internal/json"
	"github.com/loan-service/internal/logger"
)

type HandlerDependency struct {
	Logger            logger.Interface
	Context           rs.Context
	LenderLoanService llService.LenderLoanServiceInterface
	LoanService       lService.LoanServiceInterface
	LoanLogService    llogService.LoanLogServiceInterface
}

type Handler struct {
	logger            logger.Interface
	resp              handler.ResponseInterface
	context           rs.Context
	lenderLoanService llService.LenderLoanServiceInterface
	loanService       lService.LoanServiceInterface
	loanLogService    llogService.LoanLogServiceInterface
	entity            Entity
}

func NewHandler(d HandlerDependency) Handler {
	entity := NewEntity(
		EntityDependency{
			LenderLoanService: d.LenderLoanService,
			LoanService:       d.LoanService,
			LoanLogService:    d.LoanLogService,
		},
	)

	return Handler{
		logger:            d.Logger,
		resp:              handler.NewResponse(handler.Dep{}),
		context:           d.Context,
		lenderLoanService: d.LenderLoanService,
		loanService:       d.LoanService,
		loanLogService:    d.LoanLogService,
		entity:            entity,
	}
}

func (h Handler) CreateLenderLoanHandler() handler.EndpointHandler {
	return func(rw http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		ctx := r.Context()
		var request dto.CreateLenderLoanRequest
		err := json.DecodeBody(&request, r.Body)
		if err != nil {
			decodingErr := err.WrapError(errs.LoanPrefix)
			return h.resp.ImportJSONWrapError(&decodingErr)
		}

		response, jsonWebErr := h.entity.CreateLenderLoan(ctx, request)
		if jsonWebErr != nil {
			return h.resp.ImportJSONWrapError(jsonWebErr)
		}

		return h.resp.SetOkWithStatus(http.StatusCreated, &response)
	}
}
