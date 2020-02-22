package xero

type Contact struct {
	ContactID                 string     `json:"ContactID"`
	ContactStatus             string     `json:"ContactStatus"`
	Name                      string     `json:"Name"`
	FirstName                 string     `json:"FirstName"`
	LastName                  string     `json:"LastName"`
	EmailAddress              string     `json:"EmailAddress"`
	SkypeUserName             string     `json:"SkypeUserName"`
	BankAccountDetails        string     `json:"BankAccountDetails"`
	TaxNumber                 string     `json:"TaxNumber"`
	AccountsReceivableTaxType string     `json:"AccountsReceivableTaxType"`
	AccountsPayableTaxType    string     `json:"AccountsPayableTaxType"`
	Addresses                 *[]Address `json:"Addresses"`
	Phones                    *[]Phone   `json:"Phones"`
	UpdatedDateUTC            string     `json:"UpdatedDateUTC"`
	IsSupplier                bool       `json:"IsSupplier"`
	IsCustomer                bool       `json:"IsCustomer"`
	DefaultCurrency           string     `json:"DefaultCurrency"`
}
