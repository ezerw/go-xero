package xero

type TrackingCategory struct {
	Name               string            `json:"Name"`
	Status             string            `json:"Status"`
	TrackingCategoryID string            `json:"TrackingCategoryID"`
	Options            *[]TrackingOption `json:"Options"`
}