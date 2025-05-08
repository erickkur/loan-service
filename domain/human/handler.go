package human

import (
	"net/http"

	"github.com/loan-service/adapter/httpserver"
	"github.com/loan-service/application/services/human"
	"github.com/loan-service/internal/handler"
)

type HandlerDependency struct {
	Human  human.HumanServiceInterface
	Helper httpserver.HelperInterface
}

type Handler struct {
	human  human.HumanServiceInterface
	helper httpserver.HelperInterface
	resp   handler.ResponseInterface
}

func NewHandler(d HandlerDependency) Handler {
	return Handler{
		human:  d.Human,
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
			return h.resp.SetErrorWithStatus(http.StatusBadRequest, err, 1, "failed to get human data")
		}

		return h.resp.SetOk(response)
	}
}
