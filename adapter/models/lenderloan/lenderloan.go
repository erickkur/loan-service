package lenderloan

import (
	"github.com/google/uuid"
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

	bun.BaseModel `bun:"lenderloans,alias:ll"`
}

type LenderLoanModel struct {
}

func NewModel() *LenderLoanModel {
	return &LenderLoanModel{}
}
