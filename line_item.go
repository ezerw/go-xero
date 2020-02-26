package xero

type LineItem struct {
	Description    string             `json:"Description,omitempty"`
	Quantity       float64            `json:"Quantity"`
	UnitAmount     float64            `json:"UnitAmount"`
	ItemCode       string             `json:"ItemCode,omitempty"`
	AccountCode    string             `json:"AccountCode"`
	LineItemID     string             `json:"LineItemID,omitempty"`
	TaxType        string             `json:"TaxType,omitempty"`
	TaxAmount      float64            `json:"TaxAmount,omitempty"`
	LineAmount     float64            `json:"LineAmount,omitempty"`
	DiscountRate   string             `json:"DiscountRate,omitempty"`
	DiscountAmount string             `json:"DiscountAmount,omitempty"`
	Tracking       []TrackingCategory `json:"Tracking,omitempty"`
}
