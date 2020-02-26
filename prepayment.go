package xero

// Prepayment are payments made before the associated document has been created.
type Prepayment struct {
	Type            string       `json:"Type,omitempty"`
	Contact         Contact      `json:"Contact,omitempty"`
	Date            string       `json:"Date,omitempty"`
	Status          string       `json:"Status,omitempty"`
	LineAmountTypes string       `json:"LineAmountTypes,omitempty"`
	LineItems       []LineItem   `json:"LineItems,omitempty"`
	SubTotal        float64      `json:"SubTotal,omitempty"`
	TotalTax        float64      `json:"TotalTax,omitempty"`
	Total           float64      `json:"Total,omitempty"`
	UpdatedDateUTC  string       `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode    string       `json:"CurrencyCode,omitempty"`
	PrepaymentID    string       `json:"PrepaymentID,omitempty"`
	CurrencyRate    string       `json:"CurrencyRate,omitempty"`
	Reference       string       `json:"Reference,omitempty"`
	RemainingCredit string       `json:"RemainingCredit,omitempty"`
	Allocations     []Allocation `json:"Allocations,omitempty"`
	Payments        []Payment    `json:"Payments,omitempty"`
	HasAttachments  string       `json:"HasAttachments,omitempty"`
}

// Prepayments is a collection of Prepayments.
type Prepayments struct {
	Prepayments []Prepayment `json:"Prepayments"`
}
