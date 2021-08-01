package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smbank/bank"
	"github.com/smbank/logers"
)

func HandleAccount(w http.ResponseWriter, r *http.Request) {
	var account bank.IBank = &bank.Account{}
	log := logers.NewLogers()
	data := json.NewDecoder(r.Body)
	err := data.Decode(&account)
	log.CheckErr(err)

	account.Save()
	log.LogSuccess("New account was created")
}
