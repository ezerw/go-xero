package xero

type BrandingTheme struct {
	BrandingThemeID string `json:"BrandingThemeID"`
	Name            string `json:"Name"`
	LogoUrl         string `json:"LogoUrl"`
	Type            string `json:"Type"`
	SortOrder       int64  `json:"SortOrder"`
	CreatedDateUTC  string `json:"CreatedDateUTC"`
}
