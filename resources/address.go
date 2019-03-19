package resources

// Address ...
type Address struct {
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	Street3     string `json:"street3,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	Residential bool   `json:"residential,omitempty"`
	ParentID    string `json:"parent_id,omitempty"`
}
