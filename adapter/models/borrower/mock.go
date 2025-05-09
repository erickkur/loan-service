package borrower

import (
	"context"

	"errors"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
)

type MockBorrowerModel struct {
}

func (b *MockBorrowerModel) GetBorrowerByGUID(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	guid uuid.UUID,
) (*Borrower, error) {
	if guid == uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e") {
		return &Borrower{ID: 1}, nil
	}

	if guid == uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e12") {
		return nil, errors.New("borrower data not found")
	}

	return nil, nil
}
