package resources

// DiscountLine ...
type DiscountLine struct {
	ID       string `json:"id,omitempty"`
	Code     string `json:"code,omitempty"`
	Type     string `json:"type,omitempty"`
	Amount   string `json:"amount,omitempty"`
	ParentID string `json:"parent_id,omitempty"`
}
