package xero

// TrackingOption is an option from within a Tracking category.
type TrackingOption struct {
	TrackingOptionID    string `json:"TrackingOptionID,omitempty"`
	Name                string `json:"Name,omitempty"`
	Status              string `json:"Status,omitempty"`
	HasValidationErrors bool   `json:"HasValidationErrors,omitempty"`
	IsDeleted           bool   `json:"IsDeleted,omitempty"`
	IsArchived          bool   `json:"IsArchived,omitempty"`
	IsActive            bool   `json:"IsActive,omitempty"`
}

// TrackingCategory is used to segment data within a Xero organisation.
type TrackingCategory struct {
	TrackingCategoryID string           `json:"TrackingCategoryID,omitempty"`
	Name               string           `json:"Name,omitempty"`
	Status             string           `json:"Status,omitempty"`
	TrackingOptions    []TrackingOption `json:"Options,omitempty"`
}

// TrackingCategories is a collection of TrackingCategories.
type TrackingCategories struct {
	TrackingCategories []TrackingCategory `json:"TrackingCategories"`
}
