package xero

type Allocation struct {
	Amount  float64  `json:"Amount"`
	Invoice Invoice `json:"Invoice"`
	Date    string   `json:"Date"`
}
