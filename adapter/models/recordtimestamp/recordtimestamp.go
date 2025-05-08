package recordtimestamp

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type RecordTimestamp struct {
	CreatedAt time.Time
	UpdatedAt bun.NullTime
	DeletedAt bun.NullTime `bun:",soft_delete,nullzero"`
}

var _ bun.BeforeAppendModelHook = (*RecordTimestamp)(nil)

func (recordTimestamp *RecordTimestamp) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		recordTimestamp.CreatedAt = time.Now()
		recordTimestamp.UpdatedAt = bun.NullTime{Time: time.Now()}
	case *bun.UpdateQuery:
		recordTimestamp.UpdatedAt = bun.NullTime{Time: time.Now()}
	}
	return nil
}
