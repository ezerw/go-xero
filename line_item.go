package xero

// LineItem is a line containing detail on an Invoice.
type LineItem struct {
	Description    string             `json:"Description,omitempty"`
	Quantity       float64            `json:"Quantity,omitempty"`
	UnitAmount     float64            `json:"UnitAmount,omitempty"`
	ItemCode       string             `json:"ItemCode,omitempty"`
	AccountCode    string             `json:"AccountCode,omitempty"`
	LineItemID     string             `json:"LineItemID,omitempty"`
	TaxType        string             `json:"TaxType,omitempty"`
	TaxAmount      float64            `json:"TaxAmount,omitempty"`
	LineAmount     float64            `json:"LineAmount,omitempty"`
	DiscountRate   string             `json:"DiscountRate,omitempty"`
	DiscountAmount string             `json:"DiscountAmount,omitempty"`
	Tracking       []TrackingCategory `json:"Tracking,omitempty"`
}
