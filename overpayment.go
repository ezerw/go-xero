package xero

type OverPayment struct {
	Contact         *Contact      `json:"Contact"`
	DateString      string        `json:"DateString"`
	Date            string        `json:"Date"`
	Status          string        `json:"Status"`
	LineAmountTypes string        `json:"LineAmountTypes"`
	SubTotal        string        `json:"SubTotal"`
	TotalTax        string        `json:"TotalTax"`
	Total           string        `json:"Total"`
	UpdatedDateUTC  string        `json:"UpdatedDateUTC"`
	CurrencyCode    string        `json:"CurrencyCode"`
	Type            string        `json:"Type"`
	OverpaymentID   string        `json:"OverpaymentID"`
	CurrencyRate    string        `json:"CurrencyRate"`
	RemainingCredit string        `json:"RemainingCredit"`
	Allocations     *[]Allocation `json:"Allocations"`
	HasAttachments  string        `json:"HasAttachments"`
}
