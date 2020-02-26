package xero

type BatchPayment struct {
	Account        Account   `json:"Account,omitempty"`
	Particulars    string    `json:"Particulars,omitempty"` // NZ Only
	Code           string    `json:"Code,omitempty"`        // NZ Only
	Reference      string    `json:"Reference,omitempty"`   // NZ Only
	Details        string    `json:"Details,omitempty"`     // NON-NZ Only
	Narrative      string    `json:"Narrative,omitempty"`   // UK Only
	BatchPaymentID string    `json:"BatchPaymentID,omitempty"`
	Date           string    `json:"Date,omitempty"`
	Payments       []Payment `json:"Payments,omitempty"`
	Type           string    `json:"Type,omitempty"`
	Status         string    `json:"Status,omitempty"`
	TotalAmount    float64   `json:"TotalAmount,omitempty"`
	IsReconciled   bool      `json:"IsReconciled,omitempty"`
	UpdatedDateUTC string    `json:"UpdatedDateUTC,omitempty"`
}
