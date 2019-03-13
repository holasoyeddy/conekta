package resources

// Customer ...
type Customer struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	Email            string            `json:"email"`
	Phone            string            `json:"phone"`
	Corporate        bool              `json:"corporate"`
	PaymentSources   []PaymentSource   `json:"payment_sources"`
	FiscalEntities   []FiscalEntity    `json:"fiscal_entities"`
	ShippingContacts []ShippingContact `json:"shipping_contacts"`
	Subscription     Subscription      `json:"subscription"`
	AccountAge       int               `json:"account_age"`
	FirstPaidAt      int               `json:"first_paid_at"`
	CreatedAt        string            `json:"created_at"`
}
