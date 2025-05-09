package recordtimestamp

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type RecordTimestamp struct {
	CreatedAt time.Time    `bun:",,default:now()"`
	UpdatedAt time.Time    `bun:",,default:now()"`
	DeletedAt bun.NullTime `bun:",soft_delete,nullzero"`
}

var _ bun.BeforeAppendModelHook = (*RecordTimestamp)(nil)

func (recordTimestamp *RecordTimestamp) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	if _, ok := query.(*bun.UpdateQuery); ok {
		recordTimestamp.UpdatedAt = time.Now()
	}

	return nil
}
