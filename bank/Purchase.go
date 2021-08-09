package bank

type Purchase struct {
	Card            string  `json:"card_number"`
	CvNumber        uint8   `json:"cv_number"`
	TransactionType string  `json:"transaction_type"`
	PurchaseAmount  float64 `json:"purchase_amount"`
	Balance         float64 `json:"balance"`
	Refound         float64 `json:"refound"`
}
