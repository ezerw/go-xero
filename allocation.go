package xero

// Allocation allocated an overpayment or Prepayment to an Invoice.
type Allocation struct {
	Amount  float64 `json:"Amount,omitempty"`
	Invoice Invoice `json:"Invoice,omitempty"`
	Date    string  `json:"Date,omitempty"`
}

// Allocations is a collection of Allocations.
type Allocations struct {
	Allocations []Allocation `json:"Allocations"`
}
