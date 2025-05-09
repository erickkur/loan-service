package employee

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
	"github.com/loan-service/adapter/models/recordtimestamp"
	"github.com/uptrace/bun"
)

type Employee struct {
	ID        int       `bun:",pk,autoincrement"`
	GUID      uuid.UUID `bun:",nullzero"`
	FirstName string    `bun:""`
	LastName  string    `bun:""`
	recordtimestamp.RecordTimestamp

	bun.BaseModel `bun:"employees,alias:e"`
}

type EmployeeModel struct {
}

func NewModel() *EmployeeModel {
	return &EmployeeModel{}
}

func (b *EmployeeModel) GetEmployeeByGUID(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	guid uuid.UUID,
) (*Employee, error) {
	var e Employee
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewSelect().
		Model(&e).
		Where("guid = ?", guid)
	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
