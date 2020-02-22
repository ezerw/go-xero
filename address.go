package xero

type Address struct {
	AddressType  string `json:"AddressType"`
	AddressLine1 string `json:"AddressLine1"`
	City         string `json:"City"`
	PostalCode   string `json:"PostalCode"`
	AttentionTo  string `json:"AttentionTo"`
}