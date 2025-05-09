package loan

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	pa "github.com/loan-service/adapter/database/postgres"
	b "github.com/loan-service/adapter/models/borrower"
	e "github.com/loan-service/adapter/models/employee"
	l "github.com/loan-service/adapter/models/loan"
	ll "github.com/loan-service/adapter/models/loanlog"
	"github.com/loan-service/application/dto"
)

func TestCreateLoan(t *testing.T) {
	mockTrxDB := &pa.MockDatabaseAdapter{}
	dbClient := &pa.MockDatabaseAdapter{MockBeginTxAdapter: mockTrxDB}
	borrowerModel := &b.MockBorrowerModel{}
	loanModel := &l.MockLoanModel{}
	loanLogModel := &ll.MockLoanLogModel{}

	ls := &LoanService{
		dbClient:      dbClient,
		borrowerModel: borrowerModel,
		loanModel:     loanModel,
		loanLogModel:  loanLogModel,
	}

	t.Run("When all inputs are valid, then return success", func(t *testing.T) {
		request := dto.CreateLoanRequest{
			BorrowerGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
			PrincipalAmount:    10000,
			Rate:               0.05,
			ReturnOfInvestment: 12000,
		}

		createdLoan, err := ls.CreateLoan(context.Background(), request)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if createdLoan == nil {
			t.Errorf("expected created loan to be not nil")
		}
	})

	t.Run("When borrower data is invalid, then return error ", func(t *testing.T) {
		request := dto.CreateLoanRequest{
			BorrowerGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e12"),
			PrincipalAmount:    10000,
			Rate:               0.05,
			ReturnOfInvestment: 12000,
		}

		_, err := ls.CreateLoan(context.Background(), request)

		if err == nil {
			t.Fatal("expected an error but got nil")
		}

		expected := "borrower data not found"
		if err.Error() != expected {
			t.Fatalf("expected error %q, got %q", expected, err.Error())
		}
	})
}

func TestUpdateLoan(t *testing.T) {
	mockTrxDB := &pa.MockDatabaseAdapter{}
	dbClient := &pa.MockDatabaseAdapter{MockBeginTxAdapter: mockTrxDB}
	borrowerModel := &b.MockBorrowerModel{}
	loanModel := &l.MockLoanModel{}
	loanLogModel := &ll.MockLoanLogModel{}
	employeeModel := &e.MockEmployeeModel{}

	ls := &LoanService{
		dbClient:      dbClient,
		borrowerModel: borrowerModel,
		loanModel:     loanModel,
		loanLogModel:  loanLogModel,
		employeeModel: employeeModel,
	}

	t.Run("When all inputs are valid and want to change loan status to approved, then return success", func(t *testing.T) {
		loanGUID := uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1123")
		request := dto.UpdateLoanRequest{
			LoanGUID:       &loanGUID,
			Status:         "approved",
			PictureProof:   "picture.jpg",
			EmployeeGUID:   uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1124"),
			DateOfApproval: time.Now(),
		}

		l, err := ls.UpdateLoanToApprovedOrDisbursed(context.Background(), request)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if l == nil {
			t.Errorf("expected created loan to be not nil")
		}
	})

	t.Run("When current loan status is dibursed and want to change loan status to approved, then return error", func(t *testing.T) {
		loanGUID := uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1444")
		request := dto.UpdateLoanRequest{
			LoanGUID:       &loanGUID,
			Status:         "approved",
			PictureProof:   "picture.jpg",
			EmployeeGUID:   uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1124"),
			DateOfApproval: time.Now(),
		}

		_, err := ls.UpdateLoanToApprovedOrDisbursed(context.Background(), request)

		if err == nil {
			t.Fatal("expected an error but got nil")
		}

		expected := "current loan status must be proposed for updating loan to approved"
		if err.Error() != expected {
			t.Fatalf("expected error %q, got %q", expected, err.Error())
		}
	})

	t.Run("When all inputs are valid and want to change loan status to disbursed, then return success", func(t *testing.T) {
		loanGUID := uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e5555")
		request := dto.UpdateLoanRequest{
			LoanGUID:       &loanGUID,
			Status:         "disbursed",
			PictureProof:   "picture.jpg",
			EmployeeGUID:   uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1124"),
			DateOfApproval: time.Now(),
		}

		l, err := ls.UpdateLoanToApprovedOrDisbursed(context.Background(), request)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if l == nil {
			t.Errorf("expected created loan to be not nil")
		}
	})

	t.Run("When current loan status is approved and want to change loan status to approved, then return error", func(t *testing.T) {
		loanGUID := uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e6666")
		request := dto.UpdateLoanRequest{
			LoanGUID:       &loanGUID,
			Status:         "disbursed",
			PictureProof:   "picture.jpg",
			EmployeeGUID:   uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1124"),
			DateOfApproval: time.Now(),
		}

		_, err := ls.UpdateLoanToApprovedOrDisbursed(context.Background(), request)

		if err == nil {
			t.Fatal("expected an error but got nil")
		}

		expected := "current loan status must be invested for updating loan to disbursed"
		if err.Error() != expected {
			t.Fatalf("expected error %q, got %q", expected, err.Error())
		}
	})
}
