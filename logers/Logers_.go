package logers

// CheckDBErr for checking database errors exit from running server
func (s *BankServLog) CheckDBErr(err error) {
	if err != nil {
		s.DBErr.Println(err)
		return
	}

}

// CheckErr check errors
func (s *BankServLog) CheckErr(err error) {
	if err != nil {
		s.Err.Println(err)
		return
	}
}

// LogSuccess diplay any success request
func (s *BankServLog) LogSuccess(msg string) {
	s.Success.Println(msg)
}
