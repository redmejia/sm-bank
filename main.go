package main

import (
	"log"
	"net/http"

	"github.com/smbank/database"
	"github.com/smbank/handlers"
	"github.com/smbank/logers"
)

func main() {
	loger := logers.NewLogers()
	db, err := database.DbConnect()
	loger.CheckDBErr(err)
	defer db.Close()

	http.HandleFunc("/v1/account", handlers.HandleAccount)
	http.HandleFunc("/v1/deposit/transaction", handlers.HandleDeposit)
	http.HandleFunc("/v1/withdraw/transaction", handlers.HandleWithdraw)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
