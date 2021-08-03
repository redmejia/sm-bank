package logers

import (
	"log"
	"os"
)

type BankServLog struct {
	Info, Success, DBErr, Err *log.Logger
}

func NewLogers() BankServLog {
	var l BankServLog
	l.Info = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	l.DBErr = log.New(os.Stdout, "DATABASE ERROR ", log.Ldate|log.Ltime)
	l.Err = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)
	l.Success = log.New(os.Stdout, "SUCCESS ", log.Ldate|log.Ltime)
	return l
}
