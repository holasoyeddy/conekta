package resources

// Return ...
type Return struct {
	ID        string     `json:"id,omitempty"`
	LiveMode  bool       `json:"livemode,omitempty"`
	Amount    string     `json:"amount,omitempty"`
	Currency  string     `json:"currency,omitempty"`
	Items     []LineItem `json:"items,omitempty"`
	ChargeID  string     `json:"charge_id,omitempty"`
	Reason    string     `json:"reason,omitempty"`
	CreatedAt string     `json:"created_at,omitempty"`
	ParentID  string     `json:"parent_id,omitempty"`
}
