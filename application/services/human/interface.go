package human

import (
	"github.com/loan-service/adapter/models/human"
	"github.com/loan-service/application/dto"
)

type HumanServiceInterface interface {
	BuildGetHumansResponse([]human.Human) []dto.GetHumansResponse
	GetHumans(limit int) ([]human.Human, error)
}
