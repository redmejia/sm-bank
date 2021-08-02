package bank

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/smbank/database"
	"github.com/smbank/logers"
)

func (t *Transaction) Deposit(w http.ResponseWriter) {
	logr := logers.NewLogers()
	switch t.AccountType {
	case "checking":
		tx, err := database.DB.Begin()
		logr.CheckDBErr(err)
		defer tx.Rollback()

		row := tx.QueryRow(`
			SELECT 
				balance,
				card_number,
				card_cv
			FROM
				checking_acc_type
			WHERE
				card_number = $1
			AND card_cv = $2`,
			t.Card,
			t.CvNumber,
		)

		var clientInfo Transaction
		err = row.Scan(&clientInfo.Balance, &clientInfo.Card, &clientInfo.CvNumber)
		logr.CheckDBErr(err)

		if clientInfo.Card == "" || clientInfo.CvNumber == 0 {
			log.Println("Not found")
		} else {
			newBalance := clientInfo.Balance + t.Amount
			_, err := tx.Exec(`
				UPDATE 
					checking_acc_type
				SET
					balance = $1
				WHERE 
					card_number = $2 AND card_cv = $3
			`, newBalance, clientInfo.Card, clientInfo.CvNumber)

			logr.CheckDBErr(err)

			var depo = struct {
				Depo   bool    `json:"depo"`
				Amount float64 `json:"amount"`
			}{
				Depo:   true,
				Amount: t.Amount,
			}

			w.Header().Add("Content-Type", "application/json")

			data := json.NewEncoder(w)
			err = data.Encode(depo)
			logr.CheckErr(err)
		}

		err = tx.Commit()
		logr.CheckDBErr(err)
	case "saving":
		return
		// tx, err := database.DB.Begin()
		// if err != nil {
		// 	logr.Println(err)
		// }
		// defer tx.Rollback()

		// err = tx.Commit()
		// if err != nil {
		// 	logr.Println(err)
		// }
	}
}
