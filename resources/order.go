package resources

// Order represents a purchase. It contains all the details related to it, including payment method, shipment, charges, discounts, taxes and the products.
// For further information, visit https://developers.conekta.com/api?#order
type Order struct {
	ID              string          `json:"id,omitempty"`
	Object          string          `json:"object,omitempty"`
	Livemode        bool            `json:"livemode,omitempty"`
	Amount          int             `json:"amount,omitempty"`
	AmountRefunded  int             `json:"amount_refunded,omitempty"`
	PaymentStatus   string          `json:"payment_status,omitempty"`
	Currency        string          `json:"currency,omitempty"`
	CustomerInfo    Customer        `json:"customer_info,omitempty"`
	CreatedAt       int             `json:"created_at,omitempty"`
	UpdatedAt       int             `json:"updated_at,omitempty"`
	LineItems       []LineItem      `json:"line_items,omitempty"`
	Charges         []Charge        `json:"charges,omitempty"`
	ShippingLines   []ShippingLine  `json:"shipping_lines,omitempty"`
	ShippingContact ShippingContact `json:"shipping_contact,omitempty"`
	Metadata        struct{}        `json:"metadata,omitempty"`
}
