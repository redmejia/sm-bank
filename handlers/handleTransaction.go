package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smbank/bank"
	"github.com/smbank/logers"
)

var logr = logers.NewLogers()
var transaction bank.IBank

func HandleDeposit(w http.ResponseWriter, r *http.Request) {
	transaction = &bank.Transaction{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&transaction)
	logr.CheckErr(err)

	logr.InfoLog(r)
	transaction.Deposit(w)
}

func HandleWithdraw(w http.ResponseWriter, r *http.Request) {
	transaction = &bank.Transaction{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&transaction)
	logr.CheckErr(err)

	logr.InfoLog(r)
	transaction.Withdraw(w)
}

func HandlePurchase(w http.ResponseWriter, r *http.Request) {
	transaction = &bank.Purchase{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&transaction)
	logr.CheckErr(err)

	logr.InfoLog(r)
	transaction.Withdraw(w)
}
