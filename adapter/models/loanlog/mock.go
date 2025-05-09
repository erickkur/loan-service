package loanlog

import (
	"context"

	pg "github.com/loan-service/adapter/database/postgres"
	co "github.com/loan-service/internal/constant"
)

type MockLoanLogModel struct {
}

func (lm *MockLoanLogModel) CreateLoanLog(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	lg LoanLog,
) (*LoanLog, error) {
	if lg.LoanID == 1 {
		return &LoanLog{ID: 1}, nil
	}

	return nil, nil
}

func (lm *MockLoanLogModel) GetLatestLoanLog(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	loanID int64,
) (*LoanLog, error) {
	switch loanID {
	case 1:
		return &LoanLog{ID: 1, Status: co.LoanStatusProposed}, nil
	case 2:
		return &LoanLog{ID: 2, Status: co.LoanStatusDisbursed}, nil
	case 3:
		return &LoanLog{ID: 3, Status: co.LoanStatusInvested}, nil
	case 4:
		return &LoanLog{ID: 4, Status: co.LoanStatusApproved}, nil
	}

	return nil, nil
}
