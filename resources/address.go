package resources

// Address ...
type Address struct {
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	Street3     string `json:"street3"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	PostalCode  string `json:"postal_code"`
	Residential bool   `json:"residential"`
	ParentID    string `json:"parent_id"`
}
