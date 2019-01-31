package main

import (
	"database/sql"
	"log"

	"github.com/sivachandran/go-txabstraction-demo/controllers"
	"github.com/sivachandran/go-txabstraction-demo/models/postgres"
)

const PaisePerRupee = 100

func main() {
	db, err := sql.Open("postgres", "localhost:5234")
	if err != nil {
		log.Fatal(err)
	}

	sarveshAccount := postgres.NewAccount(db, "123", "Sarvesh")
	nikileshAccount := postgres.NewAccount(db, "456", "Nikilesh")

	bank := controllers.Bank{}
	err = bank.Transfer(1000*PaisePerRupee, sarveshAccount, nikileshAccount)
	if err != nil {
		log.Fatal(err)
	}
}
