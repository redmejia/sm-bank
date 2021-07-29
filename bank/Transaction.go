package bank

type Transaction struct {
	Card        string  `json:"card"`
	CvNumber    uint8   `json:"cv_number"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Balance     float64 `json:"balance"`
}
