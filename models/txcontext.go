package models

import "context"

type TxContext interface {
	WithTxContext(ctx context.Context, f func(context.Context) error) error
}
