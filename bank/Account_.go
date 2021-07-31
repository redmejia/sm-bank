package bank

import (
	"fmt"

	"github.com/smbank/database"
	"github.com/smbank/logers"
)

func (a *Account) Save() {
	tx, err := database.DB.Begin()
	logers.CheckErrLog(err)
	defer tx.Rollback()

	// client basic information
	_, err = tx.Exec(`
		INSERT INTO 
			clients (first_name, last_name, address)
		VALUES 
			($1, $2, $3)`,
		a.FirstName,
		a.LastName,
		a.Address,
	)
	logers.CheckErrLog(err)

	newCard := createCard()

	// new card client and basic information
	_, err = tx.Exec(`
		INSERT INTO
			clients_cards (first_name, last_name, address, card_number, card_cv)
		VALUES
			($1, $2, $3, $4, $5)
	`,
		a.FirstName,
		a.LastName,
		a.Address,
		newCard.cardNumber,
		newCard.cvNumber,
	)
	logers.CheckErrLog(err)

	var initBalance float64 = 0
	// new checking account and basic information
	_, err = tx.Exec(`
		INSERT INTO
			checking_acc_type (balance, first_name, last_name, address, card_number, card_cv)
		VALUES
			($1, $2, $3, $4, $5, $6)
	`,
		initBalance,
		a.FirstName,
		a.LastName,
		a.Address,
		newCard.cardNumber,
		newCard.cvNumber,
	)
	logers.CheckErrLog(err)

	// new savin account and basic information
	_, err = tx.Exec(`
		INSERT INTO
			saving_acc_type (balance, first_name, last_name, address, card_number, card_cv)
		VALUES
			($1, $2, $3, $4, $5, $6)
	`,
		initBalance,
		a.FirstName,
		a.LastName,
		a.Address,
		newCard.cardNumber,
		newCard.cvNumber,
	)
	logers.CheckErrLog(err)

	err = tx.Commit()
	logers.CheckErrLog(err)
}

func (a *Account) GetInfo() {
	fmt.Println("Name ", a.FirstName)
	fmt.Println("Name ", a.LastName)
}
