package bank

import (
	"log"
	"os"

	logers "github.com/smbank/Logers"
)

func checkErr(err error) {
	var ll logers.BankServLog
	ll.Err = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)
	if err != nil {
		ll.Err.Println(err)
	}
}

func success() {

}
