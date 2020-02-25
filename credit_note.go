package xero

type CreditNote struct {
	Type             string        `json:"Type"`
	Contact          Contact      `json:"Contact"`
	Date             string        `json:"Date"`
	Status           string        `json:"Status"`
	LineAmountTypes  string        `json:"LineAmountTypes"`
	LineItems        []LineItem   `json:"LineItems"`
	SubTotal         float64       `json:"SubTotal"`
	TotalTax         float64       `json:"TotalTax"`
	Total            float64       `json:"Total"`
	CISDeduction     string        `json:"CISDeduction"` // UK Only
	UpdatedDateUTC   string        `json:"UpdatedDateUTC"`
	CurrencyCode     string        `json:"CurrencyCode"`
	FullyPaidOnDate  string        `json:"FullyPaidOnDate"`
	CreditNoteID     string        `json:"CreditNoteID"`
	CreditNoteNumber string        `json:"CreditNoteNumber"`
	Reference        string        `json:"Reference"`
	SentToContact    bool          `json:"SentToContact"`
	CurrencyRate     float64       `json:"CurrencyRate"`
	RemainingCredit  float64       `json:"RemainingCredit"`
	Allocations      []Allocation `json:"Allocations"`
	BrandingThemeID  string        `json:"BrandingThemeID"`
	HasAttachments   bool          `json:"HasAttachments"`
}
