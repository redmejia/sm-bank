package bank

import (
	"fmt"

	"github.com/smbank/database"
)

func (a *Account) Save() {
	err := database.DB.Ping()
	if err != nil {
		fmt.Println("ERROR PING ACCOUNT ", err)
	}

	fmt.Println("name ", a.FirstName)
	fmt.Println("last name ", a.LastName)
}
