package routes

import (
	"net/http"

	"github.com/smbank/handlers"
)

func Routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/account", handlers.HandleAccount)

	mux.HandleFunc("/v1/deposit/transaction", handlers.HandleDeposit)
	mux.HandleFunc("/v1/withdraw/transaction", handlers.HandleWithdraw)
	mux.HandleFunc("/v1/purchase/transaction", handlers.HandlePurchase)

	return mux
}
