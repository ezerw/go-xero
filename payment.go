package xero

type Payment struct {
	PaymentID      string       `json:"PaymentID"`
	Date           string       `json:"Date"`
	CurrencyRate   float64      `json:"CurrencyRate"`
	BankAmount     float32      `json:"BankAmount,omitempty"`
	Amount         float32      `json:"Amount"`
	HasAccount     bool         `json:"HasAccount,omitempty"`
	Reference      string       `json:"Reference,omitempty"`
	IsReconciled   bool         `json:"IsReconciled"`
	Status         string       `json:"Status"`
	PaymentType    string       `json:"PaymentType"`
	UpdatedDateUTC string       `json:"UpdatedDateUTC"`
	BatchPaymentID string       `json:"BatchPaymentID"`
	Account        Account      `json:"Account"`
	Invoice        Invoice      `json:"Invoice"`
	CreditNote     CreditNote   `json:"CreditNote"`
	Prepayments    []Prepayment `json:"Prepayments"`
	Overpayment    OverPayment  `json:"Overpayment"`
}
