package controllers

import (
	"context"
	"fmt"

	"github.com/sivachandran/go-txabstraction-demo/models"
)

type Bank struct{}

func (b *Bank) Transfer(amount int64, from, to models.Account) error {
	return from.WithTxContext(context.Background(), func(ctx context.Context) error {
		err := from.RecordEntry(ctx, -amount, fmt.Sprintf("fund transfer to %s(%s)", to.ID, to.Name))
		if err != nil {
			return err
		}

		return to.RecordEntry(ctx, amount, fmt.Sprintf("fund transfer from %s(%s)", from.ID, from.Name))
	})
}
