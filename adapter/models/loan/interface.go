package loan

import (
	"context"

	pg "github.com/loan-service/adapter/database/postgres"
)

type LoanModelInterface interface {
	CreateLoan(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		l Loan,
	) (*Loan, error)
}
