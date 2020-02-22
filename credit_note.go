package xero

type CreditNote struct {
	CreditNoteID     string        `json:"CreditNoteID"`
	CreditNoteNumber string        `json:"CreditNoteNumber"`
	SentToContact    bool          `json:"SentToContact"`
	Payments         *[]Payment    `json:"Payments"`
	ID               string        `json:"ID"`
	HasErrors        bool          `json:"HasErrors"`
	CurrencyRate     float64       `json:"CurrencyRate"`
	Type             string        `json:"Type"`
	Reference        string        `json:"Reference"`
	RemainingCredit  float64       `json:"RemainingCredit"`
	Allocations      *[]Allocation `json:"Allocations"`
	HasAttachments   bool          `json:"HasAttachments"`
	Contact          *Contact      `json:"Contact"`
	DateString       string        `json:"DateString"`
	Date             string        `json:"Date"`
	BrandingThemeID  string        `json:"BrandingThemeID"`
	Status           string        `json:"Status"`
	LineAmountTypes  string        `json:"LineAmountTypes"`
	LineItems        *[]LineItem   `json:"LineItems"`
	SubTotal         float64       `json:"SubTotal"`
	TotalTax         float64       `json:"TotalTax"`
	Total            float64       `json:"Total"`
	UpdatedDateUTC   string        `json:"UpdatedDateUTC"`
	CurrencyCode     string        `json:"CurrencyCode"`
	FullyPaidOnDate  string        `json:"FullyPaidOnDate"`
}


