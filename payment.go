package xero

// Payment details payments against invoices and CreditNotes
type Payment struct {
	PaymentID      string       `json:"PaymentID,omitempty"`
	Date           string       `json:"Date,omitempty"`
	CurrencyRate   float64      `json:"CurrencyRate,omitempty"`
	BankAmount     float32      `json:"BankAmount,omitempty"`
	Amount         float32      `json:"Amount,omitempty"`
	HasAccount     bool         `json:"HasAccount,omitempty"`
	Reference      string       `json:"Reference,omitempty"`
	IsReconciled   bool         `json:"IsReconciled,omitempty"`
	Status         string       `json:"Status,omitempty"`
	PaymentType    string       `json:"PaymentType,omitempty"`
	UpdatedDateUTC string       `json:"UpdatedDateUTC,omitempty"`
	BatchPaymentID string       `json:"BatchPaymentID,omitempty"`
	Account        *Account      `json:"Account,omitempty"`
	Invoice        *Invoice      `json:"Invoice,omitempty"`
	CreditNote     *CreditNote   `json:"CreditNote,omitempty"`
	Prepayments    []Prepayment `json:"Prepayments,omitempty"`
	Overpayment    *OverPayment  `json:"Overpayment,omitempty"`
}

// Payments is a collection of Payments
type Payments struct {
	Payments []Payment `json:"Payments"`
}
