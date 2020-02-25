package xero

type Contact struct {
	ContactID                   string             `json:"ContactID"`
	ContactNumber               string             `json:"ContactNumber"`
	AccountNumber               string             `json:"AccountNumber"`
	ContactStatus               string             `json:"ContactStatus"`
	Name                        string             `json:"Name"`
	FirstName                   string             `json:"FirstName"`
	LastName                    string             `json:"LastName"`
	EmailAddress                string             `json:"EmailAddress"`
	SkypeUserName               string             `json:"SkypeUserName"`
	BankAccountDetails          string             `json:"BankAccountDetails"`
	TaxNumber                   string             `json:"TaxNumber"`
	AccountsReceivableTaxType   string             `json:"AccountsReceivableTaxType"`
	AccountsPayableTaxType      string             `json:"AccountsPayableTaxType"`
	Addresses                   []Address          `json:"Addresses"`
	Phones                      []Phone            `json:"Phones"`
	IsSupplier                  bool               `json:"IsSupplier"`
	IsCustomer                  bool               `json:"IsCustomer"`
	DefaultCurrency             string             `json:"DefaultCurrency"`
	UpdatedDateUTC              string             `json:"UpdatedDateUTC"`
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
	BrandingTheme               BrandingTheme      `json:"BrandingTheme,omitempty"`
	BatchPayments               BatchPayment       `json:"BatchPayments,omitempty"`
	Discount                    string             `json:"Discount,omitempty"`
	Balances                    string             `json:"Balances,omitempty"`
	HasAttachments              bool               `json:"HasAttachments,omitempty"`
}
