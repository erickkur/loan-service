package lenderloan

import (
	"context"

	pg "github.com/loan-service/adapter/database/postgres"
	lenderAdapter "github.com/loan-service/adapter/models/lender"
	lenderLoanAdapter "github.com/loan-service/adapter/models/lenderloan"
	"github.com/loan-service/application/dto"
)

type Dependency struct {
	LenderLoanModel lenderLoanAdapter.LenderLoanModelInterface
	LenderModel     lenderAdapter.LenderModelInterface
	DBClient        pg.DatabaseAdapterInterface
}

type LenderLoanService struct {
	lenderLoanModel lenderLoanAdapter.LenderLoanModelInterface
	lenderModel     lenderAdapter.LenderModelInterface
	dbClient        pg.DatabaseAdapterInterface
}

func NewLenderLoanService(d Dependency) *LenderLoanService {
	return &LenderLoanService{
		lenderLoanModel: d.LenderLoanModel,
		lenderModel:     d.LenderModel,
		dbClient:        d.DBClient,
	}
}

func (ls *LenderLoanService) CreateLenderLoan(ctx context.Context, request dto.CreateLenderLoanRequest, loanID int) (*lenderLoanAdapter.LenderLoan, error) {
	trxDB, err := ls.dbClient.BeginTransaction()
	if err != nil {
		return nil, err
	}

	defer trxDB.Rollback()

	lender, err := ls.lenderModel.GetLenderByGUID(
		ls.dbClient,
		ctx,
		request.LenderGUID,
	)
	if err != nil {
		return nil, err
	}

	l := lenderLoanAdapter.LenderLoan{
		LenderID:       int64(lender.ID),
		LoanID:         int64(loanID),
		InvestedAmount: request.InvestedAmount,
	}
	ll, err := ls.lenderLoanModel.CreateLenderloan(
		ls.dbClient,
		ctx,
		l,
	)
	if err != nil {
		return nil, err
	}

	err = trxDB.Commit()
	if err != nil {
		return nil, err
	}

	return ll, nil
}

func (ls *LenderLoanService) GetLenderLoansByLoanID(ctx context.Context, loanID int64) ([]lenderLoanAdapter.LenderLoan, error) {
	return ls.lenderLoanModel.GetLenderLoansByLoanID(
		ls.dbClient,
		ctx,
		loanID,
	)
}
