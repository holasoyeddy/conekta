package conekta

// Return ...
type Return struct {
	ID        string     `json:"id"`
	LiveMode  bool       `json:"livemode"`
	Amount    string     `json:"amount"`
	Currency  string     `json:"currency"`
	Items     []LineItem `json:"items"`
	ChargeID  string     `json:"charge_id"`
	Reason    string     `json:"reason"`
	CreatedAt string     `json:"created_at"`
	ParentID  string     `json:"parent_id"`
}
