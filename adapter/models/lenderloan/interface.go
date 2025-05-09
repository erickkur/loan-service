package lenderloan

import (
	"context"

	pg "github.com/loan-service/adapter/database/postgres"
)

type LenderLoanModelInterface interface {
	CreateLenderloan(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		l LenderLoan,
	) (*LenderLoan, error)
	GetLenderLoansByLoanID(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		loanID int64,
	) ([]LenderLoan, error)
}
