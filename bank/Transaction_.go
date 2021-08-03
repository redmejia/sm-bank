package bank

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/smbank/database"
	"github.com/smbank/logers"
)

// makeDeop make deposit to save or checking take a retrive bank account and make the update with the new balance
func makeDepo(t *Transaction, retriveQuery, updateQuery string, w http.ResponseWriter) {
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
		newBalance := clientInfo.Balance + t.Amount
		_, err := tx.Exec(updateQuery, newBalance, clientInfo.Card, clientInfo.CvNumber)

		logr.CheckDBErr(err)

		var depoReport = struct {
			Depo      bool    `json:"depo"`
			DepoToAcc string  `json:"depo_to_acc"`
			Amount    float64 `json:"amount"`
		}{
			Depo:      true,
			DepoToAcc: t.AccountType,
			Amount:    t.Amount,
		}

		w.Header().Add("Content-Type", "application/json")

		data := json.NewEncoder(w)
		err = data.Encode(depoReport)
		logr.CheckErr(err)
		logr.LogSuccess("Transaction was made")
	}

	err = tx.Commit()
	logr.CheckDBErr(err)
}

func (t Transaction) Deposit(w http.ResponseWriter) {
	switch t.AccountType {
	case "checking":
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
		makeDepo(&t, retriveStm, updateStm, w)

	case "saving":
		retriveStm := `
	 		SELECT
	 			balance,
	 			card_number,
	 			card_cv
	 		FROM
	 			saving_acc_type
	 		WHERE
	 			card_number = $1
	 		AND card_cv = $2
	 		`
		updateStm := `
	 		UPDATE
	 			saving_acc_type
	 		SET
	 			balance = $1
	 		WHERE
	 			card_number = $2 AND card_cv = $3
	 		`
		makeDepo(&t, retriveStm, updateStm, w)
	}
}

func makeWithdraw(t *Transaction, retriveQuery, updateQuery string, w http.ResponseWriter) {
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
		} else if clientInfo.Balance < t.Amount {
			log.Println("balance is less than amount")
		} else if clientInfo.Balance > t.Amount {
			fmt.Println("inside ", clientInfo)
			newBalance := clientInfo.Balance - t.Amount
			_, err := tx.Exec(updateQuery, newBalance, clientInfo.Card)

			logr.CheckDBErr(err)

			var withdrawReport = struct {
				Withdraw        bool    `json:"withdraw"`
				WithdrawFromAcc string  `json:"withdraw_from_acc"`
				Amount          float64 `json:"amount"`
				NewBalance      float64 `json:"new_balance"`
			}{
				Withdraw:        true,
				WithdrawFromAcc: t.AccountType,
				Amount:          t.Amount,
				NewBalance:      newBalance,
			}

			w.Header().Add("Content-Type", "application/json")

			data := json.NewEncoder(w)
			err = data.Encode(withdrawReport)
			logr.CheckErr(err)
			logr.LogSuccess("Transaction was made")
		}

	}

	err = tx.Commit()
	logr.CheckDBErr(err)

}

func (t Transaction) Withdraw(w http.ResponseWriter) {
	switch t.AccountType {
	case "checking":
		retriveStm := `
	  		SELECT
	  			balance,
	  			card_number,
	  			card_cv
	  		FROM
	  			checking_acc_type
	  		WHERE
	  			card_number = $1
	  		`
		updateStm := `
	  		UPDATE
	  			checking_acc_type
	  		SET
	  			balance = $1
	  		WHERE
	  			card_number = $2
	  		`
		makeWithdraw(&t, retriveStm, updateStm, w)
	}
}

func (t Transaction) Save() {}

func (t Transaction) GetInfo() {
	fmt.Println("Diplay transaction info")
	fmt.Println("Account type", t.AccountType)
	fmt.Println("Balance ", t.Balance)
	fmt.Println("Amount ", t.Amount)
}
