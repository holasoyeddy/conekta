package resources

// Charge ...
type Charge struct {
	ID            string        `json:"id"`
	LiveMode      bool          `json:"livemode"`
	CreatedAt     string        `json:"created_at"`
	Currency      string        `json:"currency"`
	Amount        int           `json:"amount"`
	ParentID      string        `json:"parent_id"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}
