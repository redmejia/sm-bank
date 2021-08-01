package logers

// CheckDBErr for checking database errors
func (s *BankServLog) CheckDBErr(err error) {
	if err != nil {
		s.DBErr.Fatal(err)
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
