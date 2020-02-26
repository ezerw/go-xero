package xero

type (
	Address struct {
		AddressType  string `json:"AddressType"`
		AddressLine1 string `json:"AddressLine1"`
		AddressLine2 string `json:"AddressLine2"`
		AddressLine3 string `json:"AddressLine3"`
		AddressLine4 string `json:"AddressLine4"`
		City         string `json:"City"`
		Region       string `json:"Region"`
		PostalCode   string `json:"PostalCode"`
		Country      string `json:"Country"`
		AttentionTo  string `json:"AttentionTo"`
	}

	Phone struct {
		PhoneType        string `json:"PhoneType"`
		PhoneNumber      string `json:"PhoneNumber"`
		PhoneAreaCode    string `json:"PhoneAreaCode"`
		PhoneCountryCode string `json:"PhoneCountryCode"`
	}
)

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
	Balances                    string             `json:"Balances,omitempty"`
	HasAttachments              bool               `json:"HasAttachments,omitempty"`
}
