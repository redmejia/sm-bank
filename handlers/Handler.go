package handlers

import (
	"github.com/smbank/bank"
	"github.com/smbank/logers"
)

var logr = logers.NewLogers()
var client bank.IBank
