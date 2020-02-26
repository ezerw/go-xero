package xero

// ContactGroup is a group of Contacts
type ContactGroup struct {
	Name           string    `json:"Name,omitempty"`
	Status         string    `json:"Status,omitempty"`
	ContactGroupID string    `json:"ContactGroupID,omitempty"`
	Contacts       []Contact `json:"Contacts,omitempty"`
}

// ContactGroups is a collection of ContactGroups
type ContactGroups struct {
	ContactGroups []ContactGroup `json:"ContactGroups"`
}
