package main

import (
	"log"

	"github.com/smbank/bank"
	"github.com/smbank/database"
)

func main() {

	db, err := database.DbConnect()
	if err != nil {
		log.Println("ERROR DB ", err)
	}
	defer db.Close()

	// http.HandleFunc("/", handlers.HandleAccount)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	var acc bank.IBank = &bank.Account{ClID: 53, FirstName: "me", LastName: "yo ", Address: "1212 main"}
	acc.Save()
	acc.GetInfo()

}
