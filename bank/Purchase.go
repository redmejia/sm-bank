package bank

type Purchase struct {
	Card            string  `json:"card"`
	CvNumber        uint8   `json:"cv_number"`
	TransactionType string  `json:"transaction_type"`
	PurchaseAmount  float64 `json:"purchase_amount"`
}
