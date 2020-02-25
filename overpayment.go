package xero

type OverPayment struct {
	Type            string       `json:"Type"`
	Contact         Contact      `json:"Contact"`
	Date            string       `json:"Date"`
	DateString      string       `json:"DateString,omitempty"`
	Status          string       `json:"Status"`
	LineAmountTypes string       `json:"LineAmountTypes,omitempty"`
	LineItems       []LineItem   `json:"LineItems"`
	SubTotal        string       `json:"SubTotal"`
	TotalTax        string       `json:"TotalTax"`
	Total           string       `json:"Total"`
	UpdatedDateUTC  string       `json:"UpdatedDateUTC"`
	CurrencyCode    string       `json:"CurrencyCode"`
	OverpaymentID   string       `json:"OverpaymentID"`
	CurrencyRate    string       `json:"CurrencyRate"`
	RemainingCredit string       `json:"RemainingCredit"`
	Allocations     []Allocation `json:"Allocations"`
	Payments        []Payment    `json:"Payments"`
	HasAttachments  string       `json:"HasAttachments"`
}
