package xero

// Overpayment is used when a debtor overpays an invoice.
type OverPayment struct {
	Type            string       `json:"Type,omitempty"`
	Contact         Contact      `json:"Contact"`
	Date            string       `json:"Date,omitempty"`
	DateString      string       `json:"DateString,omitempty"`
	Status          string       `json:"Status,omitempty"`
	LineAmountTypes string       `json:"LineAmountTypes,omitempty"`
	LineItems       []LineItem   `json:"LineItems,omitempty"`
	SubTotal        string       `json:"SubTotal,omitempty"`
	TotalTax        string       `json:"TotalTax,omitempty"`
	Total           string       `json:"Total,omitempty"`
	UpdatedDateUTC  string       `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode    string       `json:"CurrencyCode,omitempty"`
	OverpaymentID   string       `json:"OverpaymentID,omitempty"`
	CurrencyRate    string       `json:"CurrencyRate,omitempty"`
	RemainingCredit string       `json:"RemainingCredit,omitempty"`
	Allocations     []Allocation `json:"Allocations,omitempty"`
	Payments        []Payment    `json:"Payments,omitempty"`
	HasAttachments  string       `json:"HasAttachments,omitempty"`
}

// OverPayments is a collection of OverPayments.
type OverPayments struct {
	OverPayments []OverPayment `json:"Overpayments"`
}
