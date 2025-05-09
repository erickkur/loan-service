package lender

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
)

type LenderModelInterface interface {
	GetLenderByGUID(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		guid uuid.UUID,
	) (*Lender, error)
}
