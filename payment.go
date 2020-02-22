package xero

type Payment struct {
	PaymentID           string   `json:"PaymentID"`
	Date                string   `json:"Date"`
	BankAmount          float32  `json:"BankAmount"`
	Amount              float32  `json:"Amount"`
	Reference           string   `json:"Reference"`
	CurrencyRate        float64  `json:"CurrencyRate"`
	PaymentType         string   `json:"PaymentType"`
	Status              string   `json:"Status"`
	UpdatedDateUTC      string   `json:"UpdatedDateUTC"`
	HasAccount          bool     `json:"HasAccount"`
	IsReconciled        bool     `json:"IsReconciled"`
	Account             *Account `json:"Account"`
	Invoice             *Invoice `json:"Invoice"`
	HasValidationErrors bool     `json:"HasValidationErrors"`
}
