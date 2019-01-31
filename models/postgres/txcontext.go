package postgres

import (
	"context"
	"database/sql"
)

type contextKey int

var contextKeyTx = contextKey(0)

type TxContext struct {
	*sql.DB
}

func (t *TxContext) WithTxContext(ctx context.Context, f func(context.Context) error) error {
	tx, err := t.Begin()
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, contextKeyTx, tx)

	err = f(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

type Tx interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func (t *TxContext) GetContextTx(ctx context.Context) Tx {
	txValue := ctx.Value(contextKeyTx)
	switch contextTx := txValue.(type) {
	case *sql.Tx:
		return contextTx
	default:
		return t.DB
	}
}
