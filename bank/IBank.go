package bank

import "net/http"

type IBank interface {
	Save()
	GetInfo()

	Withdraw(w http.ResponseWriter) // for making transaction and purchase
	Deposit(w http.ResponseWriter)
}
