package loanlog

import (
	"context"

	pg "github.com/loan-service/adapter/database/postgres"
)

type LoanLogModelInterface interface {
	CreateLoanLog(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		lg LoanLog,
	) (*LoanLog, error)
	GetLatestLoanLog(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		loanID int64,
	) (*LoanLog, error)
}
