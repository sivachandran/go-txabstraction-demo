package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/sivachandran/go-txabstraction-demo/controllers"
)

type MockAccount struct {
	balance int64
	err     error
}

func (a *MockAccount) ID() string   { return "0" }
func (a *MockAccount) Name() string { return "MockAccount" }

func (a *MockAccount) WithTxContext(ctx context.Context, f func(ctx context.Context) error) error {
	bal := a.balance
	err := f(ctx)
	if err != nil {
		a.balance = bal
		return err
	}

	return nil
}

func (a *MockAccount) RecordEntry(ctx context.Context, amount int64, remarks string) error {
	if a.err != nil {
		return a.err
	}

	a.balance += amount
	return nil
}

func Test_Transfer(t *testing.T) {
	accountA := &MockAccount{balance: 1000}
	accountB := &MockAccount{balance: 1000}

	bank := controllers.Bank{}
	err := bank.Transfer(100, accountA, accountB)
	if err != nil {
		t.Error(err.Error())
	}

	if accountA.balance != 900 || accountB.balance != 1100 {
		t.Logf("accountA.balance: %d, accountB.balance: %d", accountA.balance, accountB.balance)
		t.Fail()
	}
}

func Test_TxRollback(t *testing.T) {
	mockErr := errors.New("error while crediting")

	accountA := &MockAccount{balance: 1000}
	accountB := &MockAccount{balance: 1000, err: mockErr}

	bank := controllers.Bank{}
	err := bank.Transfer(100, accountA, accountB)
	if err != mockErr {
		t.Logf("expected error '%v', got '%v'", mockErr, err)
		t.Fail()
	}

	if accountA.balance != 1000 || accountB.balance != 1000 {
		t.Logf("accountA.balance: %d, accountB.balance: %d", accountA.balance, accountB.balance)
		t.Fail()
	}
}
