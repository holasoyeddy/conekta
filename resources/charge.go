package resources

// Charge ...
type Charge struct {
	ID            string        `json:"id,omitempty"`
	Object        string        `json:"object,omitempty"`
	Status        string        `json:"status,omitempty"`
	Amount        int           `json:"amount,omitempty"`
	Fee           int           `json:"fee,omitempty"`
	ReferenceID   string        `json:"reference_id,omitempty"`
	OrderID       string        `json:"order_id,omitempty"`
	Livemode      bool          `json:"livemode,omitempty"`
	CreatedAt     int           `json:"created_at,omitempty"`
	Currency      string        `json:"currency,omitempty"`
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
}
