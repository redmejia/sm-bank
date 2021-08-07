package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smbank/bank"
)

func HandleAccount(w http.ResponseWriter, r *http.Request) {
	client = &bank.Account{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&client)
	logr.CheckErr(err)

	logr.InfoLog(r)
	client.GetInfo()
	client.Save()
}
