package resources

// ShippingContact ...
type ShippingContact struct {
	ID             string  `json:"id"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	Receiver       string  `json:"receiver"`
	BetweenStreets string  `json:"between_streets"`
	Address        Address `json:"address"`
	ParentID       string  `json:"parent_id"`
}
