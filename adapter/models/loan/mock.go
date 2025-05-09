package loan

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
)

type MockLoanModel struct {
}

func (lm *MockLoanModel) CreateLoan(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	l Loan,
) (*Loan, error) {
	if l.BorrowerID == 1 {
		loan := Loan{ID: 1}
		return &loan, nil
	}

	return nil, nil
}

func (lm *MockLoanModel) GetLoanByGUID(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	guid uuid.UUID,
) (*Loan, error) {
	switch guid {
	case uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1123"):
		loan := Loan{ID: 1}
		return &loan, nil
	case uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1444"):
		loan := Loan{ID: 2}
		return &loan, nil
	case uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e5555"):
		loan := Loan{ID: 3}
		return &loan, nil
	case uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e6666"):
		loan := Loan{ID: 4}
		return &loan, nil
	}

	return nil, nil
}

func (lm *MockLoanModel) UpdateLoanAgrrementLetter(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	loanID int,
	aggrementLetter string,
) error {
	return nil
}
