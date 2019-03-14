package resources

// Charge ...
type Charge struct {
	ID            string        `json:"id"`
	Object        string        `json:"object"`
	Status        string        `json:"status"`
	Amount        int           `json:"amount"`
	Fee           int           `json:"fee"`
	ReferenceID   string        `json:"reference_id"`
	OrderID       string        `json:"order_id"`
	Livemode      bool          `json:"livemode"`
	CreatedAt     int           `json:"created_at"`
	Currency      string        `json:"currency"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}

// type Charge struct {
// 	ID            string        `json:"id"`
// 	LiveMode      bool          `json:"livemode"`
// 	CreatedAt     string        `json:"created_at"`
// 	Currency      string        `json:"currency"`
// 	Amount        int           `json:"amount"`
// 	ParentID      string        `json:"parent_id"`
// 	PaymentMethod PaymentMethod `json:"payment_method"`
// }
