package loan

import (
	"context"
	"database/sql"

	pg "github.com/loan-service/adapter/database/postgres"
	borrowerAdapter "github.com/loan-service/adapter/models/borrower"
	employeeAdapter "github.com/loan-service/adapter/models/employee"
	loanAdapter "github.com/loan-service/adapter/models/loan"
	loanlogAdapter "github.com/loan-service/adapter/models/loanlog"
	"github.com/loan-service/application/dto"
	"github.com/loan-service/internal/constant"
	errs "github.com/loan-service/internal/error"
)

type Dependency struct {
	LoanModel     loanAdapter.LoanModelInterface
	BorrowerModel borrowerAdapter.BorrowerModelInterface
	LoanLogModel  loanlogAdapter.LoanLogModelInterface
	EmployeeModel employeeAdapter.EmployeeModelInterface
	DBClient      pg.DatabaseAdapterInterface
}

type LoanService struct {
	loanModel     loanAdapter.LoanModelInterface
	borrowerModel borrowerAdapter.BorrowerModelInterface
	loanLogModel  loanlogAdapter.LoanLogModelInterface
	employeeModel employeeAdapter.EmployeeModelInterface
	dbClient      pg.DatabaseAdapterInterface
}

func NewLoanService(d Dependency) *LoanService {
	return &LoanService{
		loanModel:     d.LoanModel,
		borrowerModel: d.BorrowerModel,
		loanLogModel:  d.LoanLogModel,
		employeeModel: d.EmployeeModel,
		dbClient:      d.DBClient,
	}
}

func (ls *LoanService) CreateLoan(ctx context.Context, request dto.CreateLoanRequest) (*loanAdapter.Loan, error) {
	trxDB, err := ls.dbClient.BeginTransaction()
	if err != nil {
		return nil, err
	}

	defer trxDB.Rollback()

	bw, err := ls.borrowerModel.GetBorrowerByGUID(
		ls.dbClient,
		ctx,
		request.BorrowerGUID,
	)
	if err != nil {
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
		return nil, err
	}

	err = trxDB.Commit()
	if err != nil {
		return nil, err
	}

	return createdLoan, nil
}

func (ls *LoanService) UpdateLoanToApproved(ctx context.Context, request dto.UpdateLoanRequest) (*loanAdapter.Loan, error) {
	trxDB, err := ls.dbClient.BeginTransaction()
	if err != nil {
		return nil, err
	}

	defer trxDB.Rollback()

	loan, err := ls.loanModel.GetLoanByGUID(
		ls.dbClient,
		ctx,
		*request.LoanGUID,
	)
	if err != nil {
		return nil, err
	}

	latestLoanLog, err := ls.loanLogModel.GetLatestLoanLog(
		ls.dbClient,
		ctx,
		int64(loan.ID),
	)
	if err != nil {
		return nil, err
	}

	if latestLoanLog.Status != constant.LoanStatusProposed {
		return nil, errs.CustomErrorInformation{
			ErrorInformation: "loan status must be proposed for updating loan to approved",
		}
	}

	employee, err := ls.employeeModel.GetEmployeeByGUID(
		ls.dbClient,
		ctx,
		request.EmployeeGUID,
	)
	if err != nil {
		return nil, err
	}

	llog := loanlogAdapter.LoanLog{
		LoanID:     int64(loan.ID),
		EmployeeID: sql.NullInt64{Valid: true, Int64: int64(employee.ID)},
		Status:     constant.LoanStatusApproved,
		CreatedAt:  request.DateOfApproval,
	}
	_, err = ls.loanLogModel.CreateLoanLog(
		ls.dbClient,
		ctx,
		llog,
	)
	if err != nil {
		return nil, err
	}

	err = trxDB.Commit()
	if err != nil {
		return nil, err
	}

	return loan, nil
}
