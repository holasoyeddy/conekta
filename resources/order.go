package resources

// Order ..
type Order struct {
	ID              string          `json:"id"`
	Object          string          `json:"object"`
	Livemode        bool            `json:"livemode"`
	Amount          int             `json:"amount"`
	AmountRefunded  int             `json:"amount_refunded"`
	PaymentStatus   string          `json:"payment_status"`
	Currency        string          `json:"currency"`
	CustomerInfo    Customer        `json:"customer_info"`
	CreatedAt       int             `json:"created_at"`
	UpdatedAt       int             `json:"updated_at"`
	LineItems       []LineItem      `json:"line_items"`
	Charges         []Charge        `json:"charges"`
	ShippingLines   []ShippingLine  `json:"shipping_lines"`
	ShippingContact ShippingContact `json:"shipping_contact"`
	Metadata        struct {
	} `json:"metadata"`
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
