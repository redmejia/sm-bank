package logers

import (
	"log"
	"os"
)

func CheckErrLog(err error) {
	var srv BankServLog
	srv.Err = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)
	if err != nil {
		srv.Err.Fatal(err)
	}
}

func LogSuccess(msg string) {
	var srv BankServLog
	srv.Success = log.New(os.Stdout, "SUCCESS ", log.Ldate|log.Ltime)
	srv.Success.Println(msg)
}
