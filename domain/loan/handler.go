package loan

import (
	"net/http"

	"github.com/loan-service/application/dto"
	lService "github.com/loan-service/application/services/loan"
	rs "github.com/loan-service/application/services/router"
	errs "github.com/loan-service/internal/error"
	"github.com/loan-service/internal/handler"
	"github.com/loan-service/internal/json"
	"github.com/loan-service/internal/logger"
)

type HandlerDependency struct {
	Logger      logger.Interface
	Context     rs.Context
	LoanService lService.LoanServiceInterface
}

type Handler struct {
	logger      logger.Interface
	resp        handler.ResponseInterface
	context     rs.Context
	loanService lService.LoanServiceInterface
	entity      Entity
}

func NewHandler(d HandlerDependency) Handler {
	entity := NewEntity(
		EntityDependency{
			LoanService: d.LoanService,
		},
	)

	return Handler{
		logger:      d.Logger,
		resp:        handler.NewResponse(handler.Dep{}),
		context:     d.Context,
		loanService: d.LoanService,
		entity:      entity,
	}
}

func (h Handler) CreateLoanHandler() handler.EndpointHandler {
	return func(rw http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		var request dto.CreateLoanRequest

		err := json.DecodeBody(&request, r.Body)
		if err != nil {
			decodingErr := err.WrapError(errs.LoanPrefix)
			return h.resp.ImportJSONWrapError(&decodingErr)
		}

		response, jsonWebErr := h.entity.CreateLoan(request)
		if jsonWebErr != nil {
			return h.resp.ImportJSONWrapError(jsonWebErr)
		}

		return h.resp.SetOkWithStatus(http.StatusCreated, &response)
	}
}
