package bank

import (
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

	row := tx.QueryRow(retriveQuery, t.Card)

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
			_, err := tx.Exec(updateQuery, newBalance, clientInfo.Card)
			logr.CheckDBErr(err)

			// var withdrawReport = struct {
			// 	Withdraw        bool    `json:"withdraw"`
			// 	WithdrawFromAcc string  `json:"withdraw_from_acc"`
			// 	Amount          float64 `json:"amount"`
			// 	NewBalance      float64 `json:"new_balance"`
			// }{
			// 	Withdraw:        true,
			// 	WithdrawFromAcc: t.AccountType,
			// 	Amount:          t.Amount,
			// 	NewBalance:      newBalance,
			// }

			w.Header().Add("Content-Type", "application/json")

			// data := json.NewEncoder(w)
			// err = data.Encode(withdrawReport)
			// logr.CheckErr(err)
			logr.LogSuccess("Transaction was made")
		}

	}

	err = tx.Commit()
	logr.CheckDBErr(err)

}

func (p Purchase) Withdraw(w http.ResponseWriter) {
	fmt.Println("pruchase here", p.Card)
	w.Header().Add("Content-Type", "application/json")
}

func (p Purchase) Deposit(w http.ResponseWriter) {}
func (p Purchase) Save()                         {}

func (p Purchase) GetInfo() {
	fmt.Println("purchase", p.PurchaseAmount)
}
