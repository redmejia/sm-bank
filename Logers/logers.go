package logers

import "log"

type BankServLog struct {
	Success *log.Logger
	Err     *log.Logger
}
