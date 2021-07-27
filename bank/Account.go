package bank

// Account ... Add more information such as ID, phone etc.
type Account struct {
	ClID      int    `json:"cl_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
}
