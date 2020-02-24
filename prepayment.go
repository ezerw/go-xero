package xero

type Prepayment struct {
	Type            string        `json:"Type"`
	Contact         *Contact      `json:"Contact"`
	Date            string        `json:"Date"`
	Status          string        `json:"Status"`
	LineAmountTypes string        `json:"LineAmountTypes"`
	LineItems       []*LineItem   `json:"LineItems"`
	SubTotal        float64       `json:"SubTotal"`
	TotalTax        float64       `json:"TotalTax"`
	Total           float64       `json:"Total"`
	UpdatedDateUTC  string        `json:"UpdatedDateUTC"`
	CurrencyCode    string        `json:"CurrencyCode"`
	PrepaymentID    string        `json:"PrepaymentID"`
	CurrencyRate    string        `json:"CurrencyRate"`
	Reference       string        `json:"Reference"`
	RemainingCredit string        `json:"RemainingCredit"`
	Allocations     []*Allocation `json:"Allocations"`
	Payments        []*Payment    `json:"Payments"`
	HasAttachments  string        `json:"HasAttachments"`
}
