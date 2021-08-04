package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smbank/bank"
)

func HandleDeposit(w http.ResponseWriter, r *http.Request) {
	client = &bank.Transaction{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&client)
	logr.CheckErr(err)

	logr.InfoLog(r)
	client.Deposit(w)
}

func HandleWithdraw(w http.ResponseWriter, r *http.Request) {
	client = &bank.Transaction{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&client)
	logr.CheckErr(err)

	logr.InfoLog(r)
	client.Withdraw(w)
}

func HandlePurchase(w http.ResponseWriter, r *http.Request) {
	client = &bank.Purchase{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&client)
	logr.CheckErr(err)

	logr.InfoLog(r)
	client.Withdraw(w)
}
