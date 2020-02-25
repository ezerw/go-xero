package xero

type TrackingCategory struct {
	TrackingCategoryID string           `json:"TrackingCategoryID"`
	Name               string           `json:"Name"`
	Status             string           `json:"Status"`
	TrackingOptions    []TrackingOption `json:"Options"`
}
