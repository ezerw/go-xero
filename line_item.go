package xero

type LineItem struct {
	LineItemID  string              `json:"LineItemID"`
	LineAmount  float64             `json:"LineAmount"`
	Description string              `json:"Description"`
	Quantity    float64             `json:"Quantity"`
	UnitAmount  float64             `json:"UnitAmount"`
	TaxType     string              `json:"TaxType"`
	TaxAmount   float64             `json:"TaxAmount"`
	AccountCode string              `json:"AccountCode"`
	Tracking    *[]TrackingCategory `json:"Tracking"`
}
