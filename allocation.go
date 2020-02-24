package xero

type Allocation struct {
	Amount  float64  `json:"Amount"`  //s
	Invoice *Invoice `json:"Invoice"` //s
	Date    string   `json:"Date"`    //s
}
