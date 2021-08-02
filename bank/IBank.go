package bank

import "net/http"

type IBank interface {
	GetInfo()
	Withdraw(w http.ResponseWriter)
	// Deposit(w http.ResponseWriter)
}
