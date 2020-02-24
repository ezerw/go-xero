package xero

type Account struct {
	AccountID               string `json:"AccountID"`               //s
	Code                    string `json:"Code"`                    //s
	Name                    string `json:"Name"`                    //s
	Status                  string `json:"Status"`                  //s
	Type                    string `json:"Type"`                    //s
	TaxType                 string `json:"TaxType"`                 //s
	Class                   string `json:"Class"`                   //s
	SystemAccount           string `json:"SystemAccount,omitempty"` //s
	EnablePaymentsToAccount bool   `json:"EnablePaymentsToAccount"` //s
	ShowInExpenseClaims     bool   `json:"ShowInExpenseClaims"`     //s
	Description             string `json:"Description"`             //s
	BankAccountNumber       string `json:"BankAccountNumber"`       //s
	BankAccountType         string `json:"BankAccountType"`         //s
	CurrencyCode            string `json:"CurrencyCode"`            //s
	ReportingCode           string `json:"ReportingCode"`           //s
	ReportingCodeName       string `json:"ReportingCodeName"`       //s
	HasAttachments          bool   `json:"HasAttachments"`          //s
	UpdatedDateUTC          string `json:"UpdatedDateUTC"`          //s
	AddToWatchlist          bool   `json:"AddToWatchlist"`          //s
}
