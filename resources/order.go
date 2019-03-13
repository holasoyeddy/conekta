package resources

// Order ..
type Order struct {
	ID                string          `json:"id"`
	Amount            string          `json:"amount"`
	Currency          string          `json:"currency"`
	PaymentStatus     string          `json:"payment_status"`
	LineItems         []LineItem      `json:"line_items"`
	ShippingLines     []ShippingLine  `json:"shipping_lines"`
	TaxLines          []TaxLine       `json:"tax_lines"`
	DiscountLines     []DiscountLine  `json:"discount_lines"`
	CustomerID        string          `json:"customer_id"`
	Customer          Customer        `json:"customer_info"`
	ShippingContactID string          `json:"shipping_contact_id"`
	ShippingContact   ShippingContact `json:"shipping_contact"`
	FiscalEntity      FiscalEntity    `json:"fiscal_entity"`
	Charges           []Charge        `json:"charges"`
	Returns           []Return        `json:"returns"`
	Metadata          string          `json:"metadata"`
}
