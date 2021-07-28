package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/smbank/bank"
)

func HandleAccount(w http.ResponseWriter, r *http.Request) {
	var account bank.IBank = &bank.Account{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&account)
	if err != nil {
		log.Println("Data ", err)
	}

	account.Save()
}
