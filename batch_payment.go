package xero

type BatchPayment struct {
	Account        Account    `json:"Account"`
	Particulars    string     `json:"Particulars,omitempty"` // NZ Only
	Code           string     `json:"Code,omitempty"`        // NZ Only
	Reference      string     `json:"Reference,omitempty"`   // NZ Only
	Details        string     `json:"Details,omitempty"`     // NON-NZ Only
	Narrative      string     `json:"Narrative"`             // UK Only
	BatchPaymentID string     `json:"BatchPaymentID"`
	Date           string     `json:"Date"`
	Payments       []*Payment `json:"Payments"`
	Type           string     `json:"Type"`
	Status         string     `json:"Status"`
	TotalAmount    float64    `json:"TotalAmount"`
	IsReconciled   bool       `json:"IsReconciled"`
	UpdatedDateUTC string     `json:"UpdatedDateUTC"`
}
