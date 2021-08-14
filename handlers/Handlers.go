package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smbank/bank"
	"github.com/smbank/logers"
)

var logr = logers.NewLogers()

// HandleAccount creating new client account and saving to db
func HandleAccount(w http.ResponseWriter, r *http.Request) {
	var client bank.Account
	data := json.NewDecoder(r.Body)
	err := data.Decode(&client)
	logr.CheckErr(err)

	logr.InfoLog(r)
	client.GetInfo()
	client.Save()
}

// Handle client transaction
// deposit client interface for deposit to saving/checking or deposit refound to checking
func deposit(client bank.IBank, w http.ResponseWriter) {
	client.Deposit(w)
}

// Handle client transaction
// withdraw client interface for withdraw from saving/checking
func withdraw(client bank.IBank, w http.ResponseWriter) {
	client.Withdraw(w)
}

// HandleDeposit deposit transaction
func HandleDeposit(w http.ResponseWriter, r *http.Request) {
	var client bank.Transaction
	data := json.NewDecoder(r.Body)
	err := data.Decode(&client)
	logr.CheckErr(err)

	logr.InfoLog(r)
	deposit(client, w)
}

// HandleWithdraw withdraw transaction
func HandleWithdraw(w http.ResponseWriter, r *http.Request) {
	var client bank.Transaction
	data := json.NewDecoder(r.Body)
	err := data.Decode(&client)
	logr.CheckErr(err)

	logr.InfoLog(r)
	withdraw(client, w)
}

// HandlePurchase purchase transaction
func HandlePurchase(w http.ResponseWriter, r *http.Request) {
	var client bank.Purchase
	data := json.NewDecoder(r.Body)
	err := data.Decode(&client)
	logr.CheckErr(err)

	logr.InfoLog(r)
	withdraw(client, w)
}

// HandleRefound deposit the amount refound to checking account
func HandleRefound(w http.ResponseWriter, r *http.Request) {
	var client bank.Purchase
	data := json.NewDecoder(r.Body)
	err := data.Decode(&client)
	logr.CheckErr(err)

	logr.InfoLog(r)
	deposit(client, w)
}
