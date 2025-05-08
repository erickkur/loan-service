package human

import (
	pg "github.com/loan-service/adapter/database/postgres"
	model "github.com/loan-service/adapter/models/human"
	"github.com/loan-service/application/dto"
)

type Dependency struct {
	HumanModel model.HumanModelInterface
	DBClient   pg.DatabaseAdapterInterface
}

type HumanService struct {
	humanModel model.HumanModelInterface
	dbClient   pg.DatabaseAdapterInterface
}

func NewHumanService(d Dependency) *HumanService {
	return &HumanService{
		humanModel: d.HumanModel,
		dbClient:   d.DBClient,
	}
}

func (h *HumanService) BuildGetHumansResponse(humans []model.Human) []dto.GetHumansResponse {
	response := make([]dto.GetHumansResponse, 0)
	for _, h := range humans {
		r := dto.GetHumansResponse{
			ID:   h.ID,
			Name: h.Name,
		}
		response = append(response, r)
	}
	return response
}

func (h *HumanService) GetHumans(limit int) ([]model.Human, error) {
	humans, err := h.humanModel.GetHumansData(h.dbClient, limit)
	if err != nil {
		return nil, err
	}

	return humans, nil
}
