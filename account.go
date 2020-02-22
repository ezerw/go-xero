package xero

type Account struct {
	AccountID               string `json:"AccountID"`
	Code                    string `json:"Code"`
	Name                    string `json:"Name"`
	Type                    string `json:"Type"`
	TaxType                 string `json:"TaxType"`
	EnablePaymentsToAccount bool   `json:"EnablePaymentsToAccount"`
	BankAccountNumber       string `json:"BankAccountNumber"`
	BankAccountType         string `json:"BankAccountType"`
	CurrencyCode            string `json:"CurrencyCode"`
}
