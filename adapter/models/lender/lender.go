package lender

import (
	"github.com/google/uuid"
	"github.com/loan-service/adapter/models/recordtimestamp"
	"github.com/uptrace/bun"
)

type Lender struct {
	ID        int       `bun:",pk,autoincrement"`
	GUID      uuid.UUID `bun:",nullzero"`
	FirstName string    `bun:""`
	LastName  string    `bun:""`
	recordtimestamp.RecordTimestamp

	bun.BaseModel `bun:"lenders,alias:l"`
}

type LenderModel struct {
}

func NewModel() *LenderModel {
	return &LenderModel{}
}
