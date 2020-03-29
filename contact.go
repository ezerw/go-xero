package xero

type Address struct {
	AddressType  string `json:"AddressType,omitempty"`
	AddressLine1 string `json:"AddressLine1,omitempty"`
	AddressLine2 string `json:"AddressLine2,omitempty"`
	AddressLine3 string `json:"AddressLine3,omitempty"`
	AddressLine4 string `json:"AddressLine4,omitempty"`
	City         string `json:"City,omitempty"`
	Region       string `json:"Region,omitempty"`
	PostalCode   string `json:"PostalCode,omitempty"`
	Country      string `json:"Country,omitempty"`
	AttentionTo  string `json:"AttentionTo,omitempty"`
}

type Phone struct {
	PhoneType        string `json:"PhoneType,omitempty"`
	PhoneNumber      string `json:"PhoneNumber,omitempty"`
	PhoneAreaCode    string `json:"PhoneAreaCode,omitempty"`
	PhoneCountryCode string `json:"PhoneCountryCode,omitempty"`
}

type ContactPerson struct {
	FirstName       string `json:"FirstName,omitempty"`
	LastName        string `json:"LastName,omitempty"`
	EmailAddress    string `json:"EmailAddress,omitempty"`
	IncludeInEmails bool   `json:"IncludeInEmails,omitempty"`
}

type Balance struct {
	Outstanding float64 `json:"Outstanding,omitempty"`
	Overdue     float64 `json:"Overdue,omitempty"`
}

type Balances struct {
	AccountsReceivable Balance `json:"AccountsReceivable,omitempty"`
	AccountsPayable    Balance `json:"AccountsPayable,omitempty"`
}

type Contact struct {
	ContactID                   string             `json:"ContactID,omitempty"`
	ContactNumber               string             `json:"ContactNumber,omitempty"`
	AccountNumber               string             `json:"AccountNumber,omitempty"`
	ContactStatus               string             `json:"ContactStatus,omitempty"`
	Name                        string             `json:"Name,omitempty"`
	FirstName                   string             `json:"FirstName,omitempty"`
	LastName                    string             `json:"LastName,omitempty"`
	EmailAddress                string             `json:"EmailAddress,omitempty"`
	SkypeUserName               string             `json:"SkypeUserName,omitempty"`
	BankAccountDetails          string             `json:"BankAccountDetails,omitempty"`
	TaxNumber                   string             `json:"TaxNumber,omitempty"`
	AccountsReceivableTaxType   string             `json:"AccountsReceivableTaxType,omitempty"`
	AccountsPayableTaxType      string             `json:"AccountsPayableTaxType,omitempty"`
	Addresses                   []Address          `json:"Addresses,omitempty"`
	Phones                      []Phone            `json:"Phones,omitempty"`
	IsSupplier                  bool               `json:"IsSupplier,omitempty"`
	IsCustomer                  bool               `json:"IsCustomer,omitempty"`
	DefaultCurrency             string             `json:"DefaultCurrency,omitempty"`
	UpdatedDateUTC              string             `json:"UpdatedDateUTC,omitempty"`
	ContactPersons              []ContactPerson    `json:"ContactPersons,omitempty"`
	XeroNetworkKey              string             `json:"XeroNetworkKey,omitempty"`
	SalesDefaultAccountCode     string             `json:"SalesDefaultAccountCode,omitempty"`
	PurchasesDefaultAccountCode string             `json:"PurchasesDefaultAccountCode,omitempty"`
	SalesTrackingCategories     []TrackingCategory `json:"SalesTrackingCategories,omitempty"`
	PurchasesTrackingCategories []TrackingCategory `json:"PurchasesTrackingCategories,omitempty"`
	TrackingCategoryName        string             `json:"TrackingCategoryName,omitempty"`
	TrackingOptionName          string             `json:"TrackingOptionName,omitempty"`
	PaymentTerms                string             `json:"PaymentTerms,omitempty"`
	ContactGroups               []ContactGroup     `json:"ContactGroups,omitempty"`
	Website                     string             `json:"Website,omitempty"`
	BrandingTheme               *BrandingTheme     `json:"BrandingTheme,omitempty"`
	BatchPayments               *BatchPayment      `json:"BatchPayments,omitempty"`
	Discount                    string             `json:"Discount,omitempty"`
	Balances                    *Balances          `json:"Balances,omitempty"`
	HasAttachments              bool               `json:"HasAttachments,omitempty"`
}

type Contacts struct {
	Contacts []Contact `json:"Contacts"`
}
