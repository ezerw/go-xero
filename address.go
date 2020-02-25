package xero

type Address struct {
	AddressType  string `json:"AddressType"`
	AddressLine1 string `json:"AddressLine1"`
	AddressLine2 string `json:"AddressLine2"`
	AddressLine3 string `json:"AddressLine3"`
	AddressLine4 string `json:"AddressLine4"`
	City         string `json:"City"`
	Region       string `json:"Region"`
	PostalCode   string `json:"PostalCode"`
	Country      string `json:"Country"`
	AttentionTo  string `json:"AttentionTo"`
}
