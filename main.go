package main

import (
	"log"
	"net/http"

	"github.com/smbank/database"
	"github.com/smbank/logers"
	"github.com/smbank/routes"
)

func main() {
	loger := logers.NewLogers()
	db, err := database.DbConnect()
	loger.CheckDBErr(err)
	defer db.Close()

	srv := &http.Server{
		Addr:    ":8081", // bank serv port
		Handler: routes.Routes(),
	}

	log.Println("Server running on port :8081")
	err = srv.ListenAndServe()
	loger.CheckErr(err)
}
