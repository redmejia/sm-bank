package logers

import (
	"log"
	"os"
)

type BankServLog struct {
	Success *log.Logger
	DBErr   *log.Logger
	Err     *log.Logger
}

func NewLogers() BankServLog {
	var l BankServLog
	l.DBErr = log.New(os.Stdout, "DATABASE ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Err = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)
	l.Success = log.New(os.Stdout, "SUCCESS ", log.Ldate|log.Ltime)
	return l
}
