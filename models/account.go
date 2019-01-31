package models

import "context"

type Account interface {
	TxContext // models.TxContext

	ID() string
	Name() string
	RecordEntry(ctx context.Context, amount int64, remarks string) error
}
