package human

import (
	"net/http"

	"github.com/loan-service/adapter/httpserver"
	"github.com/loan-service/application/services/human"
	"github.com/loan-service/internal/handler"
	"github.com/loan-service/internal/logger"
)

type HandlerDependency struct {
	Human  human.HumanServiceInterface
	Logger logger.Interface
	Helper httpserver.HelperInterface
}

type Handler struct {
	human  human.HumanServiceInterface
	helper httpserver.HelperInterface
	logger logger.Interface
	resp   handler.ResponseInterface
}

func NewHandler(d HandlerDependency) Handler {
	return Handler{
		human:  d.Human,
		logger: d.Logger,
		helper: d.Helper,
		resp:   handler.NewResponse(handler.Dep{}),
	}
}

func (h Handler) GetHumans() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) (res handler.ResponseInterface) {
		entity := NewHumanEntity(EntityDependency{
			HumanService: h.human,
		})

		response, err := entity.GetHumans()
		if err != nil {
			h.logger.Error(err.Error())
			return h.resp.SetErrorWithStatus(http.StatusBadRequest, err, 1, "failed to get human data")
		}

		return h.resp.SetOk(response)
	}
}
