package xero

type Phone struct {
	PhoneType        string `json:"PhoneType"`
	PhoneNumber      string `json:"PhoneNumber"`
	PhoneAreaCode    string `json:"PhoneAreaCode"`
	PhoneCountryCode string `json:"PhoneCountryCode"`
}
