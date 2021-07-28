package main

import (
	"log"
	"net/http"

	"github.com/smbank/database"
	"github.com/smbank/handlers"
)

func main() {

	db, err := database.DbConnect()
	if err != nil {
		log.Println("ERROR DB ", err)
	}
	defer db.Close()

	http.HandleFunc("/", handlers.HandleAccount)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
