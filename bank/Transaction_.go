package bank

import (
	"log"

	"github.com/smbank/database"
)

func (t *Transaction) Deposit() {
	switch t.AccountType {
	case "checking":
		tx, err := database.DB.Begin()
		if err != nil {
			log.Println(err)
		}
		defer tx.Rollback()

		err = tx.Commit()
		if err != nil {
			log.Println(err)
		}
	case "saving":
		tx, err := database.DB.Begin()
		if err != nil {
			log.Println(err)
		}
		defer tx.Rollback()

		err = tx.Commit()
		if err != nil {
			log.Println(err)
		}
	}
}
