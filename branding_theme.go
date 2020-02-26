package xero

// BrandingTheme applies structure and visuals to an invoice when printed or sent.
type BrandingTheme struct {
	BrandingThemeID string  `json:"BrandingThemeID,omitempty"`
	Name            string  `json:"Name,omitempty"`
	LogoUrl         string  `json:"LogoUrl,omitempty"`
	Type            string  `json:"Type,omitempty"`
	SortOrder       float64 `json:"SortOrder,omitempty"`
	CreatedDateUTC  string  `json:"CreatedDateUTC,omitempty"`
}

// BrandingThemes contains a collection of BrandingThemes.
type BrandingThemes struct {
	BrandingThemes []BrandingTheme `json:"BrandingThemes"`
}
