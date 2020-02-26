package xero

// CreditNote an be raised directly against a customer or supplier,
// allowing the customer or supplier to be held in credit until a future invoice or bill is raised.
type CreditNote struct {
	Type             string       `json:"Type,omitempty"`
	Contact          Contact      `json:"Contact"`
	Date             string       `json:"Date,omitempty"`
	Status           string       `json:"Status,omitempty"`
	LineAmountTypes  string       `json:"LineAmountTypes,omitempty"`
	LineItems        []LineItem   `json:"LineItems,omitempty"`
	SubTotal         float64      `json:"SubTotal,omitempty"`
	TotalTax         float64      `json:"TotalTax,omitempty"`
	Total            float64      `json:"Total,omitempty"`
	CISDeduction     string       `json:"CISDeduction,omitempty"` // UK Only
	UpdatedDateUTC   string       `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode     string       `json:"CurrencyCode,omitempty"`
	FullyPaidOnDate  string       `json:"FullyPaidOnDate,omitempty"`
	CreditNoteID     string       `json:"CreditNoteID,omitempty"`
	CreditNoteNumber string       `json:"CreditNoteNumber,omitempty"`
	Reference        string       `json:"Reference,omitempty"`
	SentToContact    bool         `json:"SentToContact,omitempty"`
	CurrencyRate     float64      `json:"CurrencyRate,omitempty"`
	RemainingCredit  float64      `json:"RemainingCredit,omitempty"`
	Allocations      []Allocation `json:"Allocations,omitempty"`
	BrandingThemeID  string       `json:"BrandingThemeID,omitempty"`
	HasAttachments   bool         `json:"HasAttachments,omitempty"`
}
