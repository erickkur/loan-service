package borrower

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
)

type BorrowerModelInterface interface {
	GetBorrowerByGUID(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		guid uuid.UUID,
	) (*Borrower, error)
}
