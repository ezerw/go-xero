package xero

type Account struct {
	AccountID               string `json:"AccountID,omitempty"`
	Code                    string `json:"Code,omitempty"`
	Name                    string `json:"Name,omitempty"`
	Status                  string `json:"Status,omitempty"`
	Type                    string `json:"Type,omitempty"`
	TaxType                 string `json:"TaxType,omitempty"`
	Class                   string `json:"Class,omitempty"`
	SystemAccount           string `json:"SystemAccount,omitempty"`
	EnablePaymentsToAccount bool   `json:"EnablePaymentsToAccount,omitempty"`
	ShowInExpenseClaims     bool   `json:"ShowInExpenseClaims,omitempty"`
	Description             string `json:"Description,omitempty"`
	BankAccountNumber       string `json:"BankAccountNumber,omitempty"`
	BankAccountType         string `json:"BankAccountType,omitempty"`
	CurrencyCode            string `json:"CurrencyCode,omitempty"`
	ReportingCode           string `json:"ReportingCode,omitempty"`
	ReportingCodeName       string `json:"ReportingCodeName,omitempty"`
	HasAttachments          bool   `json:"HasAttachments,omitempty"`
	UpdatedDateUTC          string `json:"UpdatedDateUTC,omitempty"`
	AddToWatchlist          bool   `json:"AddToWatchlist,omitempty"`
}

// Accounts contains a collection of Accounts
type Accounts struct {
	Accounts []Account `json:"Accounts,omitempty"`
}
