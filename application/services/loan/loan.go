package loan

import (
	"context"
	"fmt"

	pg "github.com/loan-service/adapter/database/postgres"
	borrowerAdapter "github.com/loan-service/adapter/models/borrower"
	loanAdapter "github.com/loan-service/adapter/models/loan"
	loanlogAdapter "github.com/loan-service/adapter/models/loanlog"
	"github.com/loan-service/application/dto"
)

type Dependency struct {
	LoanModel     loanAdapter.LoanModelInterface
	BorrowerModel borrowerAdapter.BorrowerModelInterface
	LoanLogModel  loanlogAdapter.LoanLogModelInterface
	DBClient      pg.DatabaseAdapterInterface
}

type LoanService struct {
	loanModel     loanAdapter.LoanModelInterface
	borrowerModel borrowerAdapter.BorrowerModelInterface
	loanLogModel  loanlogAdapter.LoanLogModelInterface
	dbClient      pg.DatabaseAdapterInterface
}

func NewLoanService(d Dependency) *LoanService {
	return &LoanService{
		loanModel:     d.LoanModel,
		borrowerModel: d.BorrowerModel,
		loanLogModel:  d.LoanLogModel,
		dbClient:      d.DBClient,
	}
}

func (ls *LoanService) CreateLoan(ctx context.Context, request dto.CreateLoanRequest) (*loanAdapter.Loan, error) {
	trxDB, err := ls.dbClient.BeginTransaction()
	if err != nil {
		fmt.Println("Failed to initiate transaction")
		return nil, err
	}

	defer trxDB.Rollback()

	bw, err := ls.borrowerModel.GetBorrowerByGUID(
		ls.dbClient,
		ctx,
		request.BorrowerGUID,
	)
	if err != nil {
		fmt.Println("Failed to GetBorrowerByGUID")
		return nil, err
	}

	l := loanAdapter.Loan{
		BorrowerID:         int64(bw.ID),
		PrincipalAmount:    request.PrincipalAmount,
		Rate:               request.Rate,
		ReturnOfInvestment: request.ReturnOfInvestment,
	}
	createdLoan, err := ls.loanModel.CreateLoan(
		ls.dbClient,
		ctx,
		l,
	)
	if err != nil {
		fmt.Println("Failed to CreateLoan")
		return nil, err
	}

	llog := loanlogAdapter.LoanLog{
		LoanID: int64(createdLoan.ID),
	}
	_, err = ls.loanLogModel.CreateLoanLog(
		ls.dbClient,
		ctx,
		llog,
	)
	if err != nil {
		fmt.Println("Failed to CreateLoanLog")
		return nil, err
	}

	err = trxDB.Commit()
	if err != nil {
		fmt.Println("Failed to commit")
		return nil, err
	}

	return createdLoan, nil
}
