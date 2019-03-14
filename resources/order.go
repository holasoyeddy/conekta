package resources

// Order ..
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

// type Order struct {
// 	ID                string          `json:"id,omitempty"`
// 	Amount            string          `json:"amount,omitempty"`
// 	Currency          string          `json:"currency,omitempty"`
// 	PaymentStatus     string          `json:"payment_status,omitempty"`
// 	LineItems         []LineItem      `json:"line_items,omitempty"`
// 	ShippingLines     []ShippingLine  `json:"shipping_lines,omitempty"`
// 	TaxLines          []TaxLine       `json:"tax_lines,omitempty"`
// 	DiscountLines     []DiscountLine  `json:"discount_lines,omitempty"`
// 	CustomerID        string          `json:"customer_id,omitempty"`
// 	Customer          Customer        `json:"customer_info,omitempty"`
// 	ShippingContactID string          `json:"shipping_contact_id,omitempty"`
// 	ShippingContact   ShippingContact `json:"shipping_contact,omitempty"`
// 	FiscalEntity      FiscalEntity    `json:"fiscal_entity,omitempty"`
// 	Charges           []Charge        `json:"charges,omitempty"`
// 	Returns           []Return        `json:"returns,omitempty"`
// 	Metadata          string          `json:"metadata,omitempty"`
// }
