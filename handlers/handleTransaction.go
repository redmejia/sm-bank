package handlers

import (
	"fmt"
	"net/http"
)

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
