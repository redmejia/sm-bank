package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smbank/bank"
	"github.com/smbank/logers"
)

func HandleTransaction(w http.ResponseWriter, r *http.Request) {
	log := logers.NewLogers()

	var transaction bank.Transaction
	data := json.NewDecoder(r.Body)
	err := data.Decode(&transaction)
	log.CheckErr(err)
	transaction.Deposit(w)

}
