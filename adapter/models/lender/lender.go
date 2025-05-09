package lender

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
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

func (b *LenderModel) GetLenderByGUID(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	guid uuid.UUID,
) (*Lender, error) {
	var l Lender
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewSelect().
		Model(&l).
		Where("guid = ?", guid)
	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
