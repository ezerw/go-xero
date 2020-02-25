package xero

type LineItem struct {
	Description    string             `json:"Description"`
	Quantity       float64            `json:"Quantity"`
	UnitAmount     float64            `json:"UnitAmount"`
	ItemCode       string             `json:"ItemCode,omitempty"`
	AccountCode    string             `json:"AccountCode"`
	LineItemID     string             `json:"LineItemID"`
	TaxType        string             `json:"TaxType"`
	TaxAmount      float64            `json:"TaxAmount"`
	LineAmount     float64            `json:"LineAmount"`
	DiscountRate   string             `json:"DiscountRate,omitempty"`
	DiscountAmount string             `json:"DiscountAmount,omitempty"`
	Tracking       []TrackingCategory `json:"Tracking"`
}
