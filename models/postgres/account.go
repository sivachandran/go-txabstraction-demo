package postgres

import (
	"context"
	"database/sql"

	"github.com/sivachandran/go-txabstraction-demo/models"
)

type account struct {
	TxContext // postgres.TxContext
	id        string
	name      string
}

func NewAccount(db *sql.DB, id, name string) models.Account {
	return &account{TxContext: TxContext{DB: db}, id: id, name: name}
}

func (a *account) ID() string   { return a.id }
func (a *account) Name() string { return a.name }

func (a *account) RecordEntry(ctx context.Context, amount int64, remarks string) error {
	tx := a.GetContextTx(ctx)
	_, err := tx.Exec(`INSERT INTO account_entries(accound_id, amount, remarks)
                    VALUES ($1, $2, $3)`, a.ID, amount, remarks)
	return err
}
