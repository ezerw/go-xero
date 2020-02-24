package xero

type ContactPerson struct {
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	EmailAddress    string `json:"EmailAddress"`
	IncludeInEmails bool   `json:"IncludeInEmails"`
}
