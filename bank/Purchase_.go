package bank

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/smbank/database"
	"github.com/smbank/logers"
)

func makePurchase(t *Purchase, retriveQuery, updateQuery string, w http.ResponseWriter) {
	logr := logers.NewLogers()

	tx, err := database.DB.Begin()
	logr.CheckDBErr(err)
	defer tx.Rollback()

	row := tx.QueryRow(retriveQuery, t.Card, t.CvNumber)

	var clientInfo Transaction
	err = row.Scan(
		&clientInfo.Balance,
		&clientInfo.Card,
		&clientInfo.CvNumber,
	)
	logr.CheckDBErr(err)

	if clientInfo.Card == "" || clientInfo.CvNumber == 0 {
		log.Println("Not found")
	} else {

		if clientInfo.Balance == 0 {
			log.Panicln("Balance of cero ")
		} else if clientInfo.Balance < t.PurchaseAmount {
			log.Println("balance is less than amount")
		} else if clientInfo.Balance > t.PurchaseAmount {
			newBalance := clientInfo.Balance - t.PurchaseAmount
			_, err := tx.Exec(updateQuery, newBalance, clientInfo.Card, clientInfo.CvNumber)
			logr.CheckDBErr(err)

			var requestReport = struct {
				Withdraw       bool    `json:"withdraw"`
				TransationType string  `json:"transation_type"`
				Amount         float64 `json:"amount"`
				NewBalance     float64 `json:"new_balance"`
			}{
				Withdraw:       true,
				TransationType: t.TransactionType,
				Amount:         t.PurchaseAmount,
				NewBalance:     newBalance,
			}

			var purchaseStatus = struct {
				Status          string `json:"status"`
				TransactionCode uint8  `json:"transaction_code"`
			}{
				Status:          "APROVED",
				TransactionCode: 02,
			}

			log.Println("purchase Response", requestReport)
			w.Header().Add("Content-Type", "application/json")

			data := json.NewEncoder(w)
			err = data.Encode(purchaseStatus)
			logr.CheckErr(err)
			logr.LogSuccess("Transaction was made")
		}

	}

	err = tx.Commit()
	logr.CheckDBErr(err)

}

func (p Purchase) Withdraw(w http.ResponseWriter) {
	retriveStm := `
	 		SELECT
	 			balance,
	 			card_number,
	 			card_cv
	 		FROM
	 			checking_acc_type
	 		WHERE
	 			card_number = $1
	 		AND card_cv = $2
	 		`
	updateStm := `
	 		UPDATE
	 			checking_acc_type
	 		SET
	 			balance = $1
	 		WHERE
	 			card_number = $2 AND card_cv = $3
	 		`

	makePurchase(&p, retriveStm, updateStm, w)
}

func (p Purchase) Deposit(w http.ResponseWriter) {}
func (p Purchase) Save()                         {}

func (p Purchase) GetInfo() {
	fmt.Println("purchase", p.PurchaseAmount, p.Card, p.CvNumber, p.TransactionType)
}
