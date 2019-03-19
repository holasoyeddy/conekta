package resources

// PaymentSource ...
type PaymentSource struct {
	ID       string `json:"id,omitempty"`
	TokenID  string `json:"token_id,omitempty"`
	Type     string `json:"type,omitempty"`
	Deleted  bool   `json:"deleted,omitempty"`
	ParentID string `json:"parent_id,omitempty"`
}
