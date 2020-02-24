package xero

type Address struct {
	AddressType  string `json:"AddressType"`  //s
	AddressLine1 string `json:"AddressLine1"` //s
	AddressLine2 string `json:"AddressLine2"` //s
	AddressLine3 string `json:"AddressLine3"` //s
	AddressLine4 string `json:"AddressLine4"` //s
	City         string `json:"City"`         //s
	Region       string `json:"Region"`       //s
	PostalCode   string `json:"PostalCode"` //s
	Country      string `json:"Country"` //s
	AttentionTo  string `json:"AttentionTo"` //s
}
