package lenderloan

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
	"github.com/loan-service/adapter/models/recordtimestamp"
	"github.com/uptrace/bun"
)

type LenderLoan struct {
	ID             int       `bun:",pk,autoincrement"`
	GUID           uuid.UUID `bun:",nullzero"`
	LoanID         int64     `bun:",notnull"`
	LenderID       int64     `bun:",notnull"`
	InvestedAmount int64     `bun:",notnull"`
	recordtimestamp.RecordTimestamp

	bun.BaseModel `bun:"lender_loans,alias:ll"`
}

type LenderLoanModel struct {
}

func NewModel() *LenderLoanModel {
	return &LenderLoanModel{}
}

func (m *LenderLoanModel) CreateLenderloan(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	l LenderLoan,
) (*LenderLoan, error) {
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewInsert().
		Model(&l)
	_, err = query.Exec(ctx)

	return &l, err
}

func (m *LenderLoanModel) GetLenderLoansByLoanID(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	loanID int64,
) ([]LenderLoan, error) {
	var ll []LenderLoan
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewSelect().
		Model(&ll).
		Where("loan_id = ?", loanID)
	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return ll, nil
}
