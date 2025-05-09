package borrower

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
	"github.com/loan-service/adapter/models/recordtimestamp"
	"github.com/uptrace/bun"
)

type Borrower struct {
	ID        int       `bun:",pk,autoincrement"`
	GUID      uuid.UUID `bun:",nullzero"`
	FirstName string    `bun:""`
	LastName  string    `bun:""`
	recordtimestamp.RecordTimestamp

	bun.BaseModel `bun:"borrowers,alias:b"`
}

type BorrowerModel struct {
}

func NewModel() *BorrowerModel {
	return &BorrowerModel{}
}

func (b *BorrowerModel) GetBorrowerByGUID(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	guid uuid.UUID,
) (*Borrower, error) {
	var bo Borrower
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewSelect().
		Model(&bo).
		Where("guid = ?", guid)
	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &bo, nil
}
