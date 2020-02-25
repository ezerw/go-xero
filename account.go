package xero

type Account struct {
	AccountID               string `json:"AccountID"`
	Code                    string `json:"Code"`
	Name                    string `json:"Name"`
	Status                  string `json:"Status"`
	Type                    string `json:"Type"`
	TaxType                 string `json:"TaxType"`
	Class                   string `json:"Class"`
	SystemAccount           string `json:"SystemAccount,omitempty"`
	EnablePaymentsToAccount bool   `json:"EnablePaymentsToAccount"`
	ShowInExpenseClaims     bool   `json:"ShowInExpenseClaims"`
	Description             string `json:"Description"`
	BankAccountNumber       string `json:"BankAccountNumber"`
	BankAccountType         string `json:"BankAccountType"`
	CurrencyCode            string `json:"CurrencyCode"`
	ReportingCode           string `json:"ReportingCode"`
	ReportingCodeName       string `json:"ReportingCodeName"`
	HasAttachments          bool   `json:"HasAttachments"`
	UpdatedDateUTC          string `json:"UpdatedDateUTC"`
	AddToWatchlist          bool   `json:"AddToWatchlist"`
}
