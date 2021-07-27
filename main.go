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

	acc := bank.Account{ClID: 53, FirstName: "Reynaldo", LastName: "The king", Address: "1 11 "}

	acc.Save()
}
