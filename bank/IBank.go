package bank

import "net/http"

type IBank interface {
	Save()
	GetInfo()

	Withdraw(w http.ResponseWriter)
	Deposit(w http.ResponseWriter)
}
