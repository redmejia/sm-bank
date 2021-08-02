package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smbank/bank"
	"github.com/smbank/logers"
)

func HandleDeposit(w http.ResponseWriter, r *http.Request) {
	log := logers.NewLogers()

	var transaction bank.Transaction
	data := json.NewDecoder(r.Body)
	err := data.Decode(&transaction)
	log.CheckErr(err)
	transaction.Deposit(w)

}

func HandleWithdraw(w http.ResponseWriter, r *http.Request) {
	logr := logers.NewLogers()

	var transaction bank.IBank = &bank.Transaction{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&transaction)
	logr.CheckErr(err)
	transaction.Withdraw(w)
}
