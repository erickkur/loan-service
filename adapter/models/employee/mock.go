package employee

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
)

type MockEmployeeModel struct {
}

func (m *MockEmployeeModel) GetEmployeeByGUID(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	guid uuid.UUID,
) (*Employee, error) {
	if guid == uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1124") {
		e := Employee{ID: 1}
		return &e, nil
	}

	return nil, nil
}
