package loan

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
)

type LoanModelInterface interface {
	CreateLoan(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		l Loan,
	) (*Loan, error)
	GetLoanByGUID(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		guid uuid.UUID,
	) (*Loan, error)
	UpdateLoanAgrrementLetter(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		loanID int,
		aggrementLetter string,
	) error
}
