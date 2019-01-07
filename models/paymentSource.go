package conekta

// PaymentSource ...
type PaymentSource struct {
	ID       string `json:"id"`
	TokenID  string `json:"token_id"`
	Type     string `json:"type"`
	Deleted  bool   `json:"deleted"`
	ParentID string `json:"parent_id"`
}
