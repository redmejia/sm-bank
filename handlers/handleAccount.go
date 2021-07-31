package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smbank/bank"
	"github.com/smbank/logers"
)

func HandleAccount(w http.ResponseWriter, r *http.Request) {
	var account bank.IBank = &bank.Account{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&account)
	logers.CheckErrLog(err)

	account.Save()
	logers.LogSuccess("New account was created")
}
