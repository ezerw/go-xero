package xero

type ContactGroup struct {
	Name           string    `json:"Name"`
	Status         string    `json:"Status"`
	ContactGroupID string    `json:"ContactGroupID"`
	Contacts       []Contact `json:"Contacts,omitempty"`
}
