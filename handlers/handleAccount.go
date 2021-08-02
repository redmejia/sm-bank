package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smbank/bank"
	"github.com/smbank/logers"
)

func HandleAccount(w http.ResponseWriter, r *http.Request) {
	log := logers.NewLogers()

	var account bank.IBank = &bank.Account{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&account)
	log.CheckErr(err)

	account.Save()
}
