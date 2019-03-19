package resources

// ShippingContact ...
type ShippingContact struct {
	ID             string  `json:"id,omitempty"`
	Email          string  `json:"email,omitempty"`
	Phone          string  `json:"phone,omitempty"`
	Receiver       string  `json:"receiver,omitempty"`
	BetweenStreets string  `json:"between_streets,omitempty"`
	Address        Address `json:"address,omitempty"`
	ParentID       string  `json:"parent_id,omitempty"`
}
