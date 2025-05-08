package human

import (
	hService "github.com/loan-service/application/services/human"

	"github.com/loan-service/application/dto"
)

type EntityDependency struct {
	HumanService hService.HumanServiceInterface
}

type HumanEntity struct {
	humanService hService.HumanServiceInterface
}

func NewHumanEntity(d EntityDependency) HumanEntity {
	return HumanEntity{
		humanService: d.HumanService,
	}
}

func (e HumanEntity) GetHumans() ([]dto.GetHumansResponse, error) {
	humans, err := e.humanService.GetHumans(10)
	if err != nil {
		return nil, err
	}

	response := e.humanService.BuildGetHumansResponse(humans)

	return response, nil
}
