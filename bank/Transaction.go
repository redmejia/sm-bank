package bank

type Transaction struct {
	Card        string  `json:"card"` // change to json card_number
	CvNumber    uint8   `json:"cv_number"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Balance     float64 `json:"balance"`
}
