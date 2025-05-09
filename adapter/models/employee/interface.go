package employee

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
)

type EmployeeModelInterface interface {
	GetEmployeeByGUID(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		guid uuid.UUID,
	) (*Employee, error)
}
