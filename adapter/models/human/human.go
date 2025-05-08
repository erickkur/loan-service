package human

import (
	"context"

	pg "github.com/loan-service/adapter/database/postgres"
	"github.com/uptrace/bun"
)

type Human struct {
	ID   int    `pg:",pk" json:"id"`
	Name string `json:"name"`

	bun.BaseModel `bun:"humans,alias:h"`
}

type HumanModel struct {
}

func NewModel() *HumanModel {
	return &HumanModel{}
}

func (m *HumanModel) GetHumansData(dbClient pg.DatabaseAdapterInterface, limit int) ([]Human, error) {
	var humans []Human
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	err = db.GetConnectionDB().
		NewSelect().
		Model(&humans).
		Limit(limit).
		Scan(context.TODO())
	if err != nil {
		return nil, err
	}

	return humans, nil
}
