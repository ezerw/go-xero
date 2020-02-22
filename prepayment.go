package xero

type Prepayment struct {
	Date            string        `json:"Date"`
	Status          string        `json:"Status"`
	LineAmountTypes string        `json:"LineAmountTypes"`
	SubTotal        string        `json:"SubTotal"`
	TotalTax        string        `json:"TotalTax"`
	Total           string        `json:"Total"`
	UpdatedDateUTC  string        `json:"UpdatedDateUTC"`
	CurrencyCode    string        `json:"CurrencyCode"`
	FullyPaidOnDate string        `json:"FullyPaidOnDate"`
	Type            string        `json:"Type"`
	PrepaymentID    string        `json:"PrepaymentID"`
	CurrencyRate    string        `json:"CurrencyRate"`
	RemainingCredit string        `json:"RemainingCredit"`
	Allocations     *[]Allocation `json:"Allocations"`
	HasAttachments  string        `json:"HasAttachments"`
}
