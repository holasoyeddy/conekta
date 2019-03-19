package resources

// Customer ...
type Customer struct {
	ID               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Email            string            `json:"email,omitempty"`
	Phone            string            `json:"phone,omitempty"`
	Corporate        bool              `json:"corporate,omitempty"`
	PaymentSources   []PaymentSource   `json:"payment_sources,omitempty"`
	FiscalEntities   []FiscalEntity    `json:"fiscal_entities,omitempty"`
	ShippingContacts []ShippingContact `json:"shipping_contacts,omitempty"`
	Subscription     Subscription      `json:"subscription,omitempty"`
	AccountAge       int               `json:"account_age,omitempty"`
	FirstPaidAt      int               `json:"first_paid_at,omitempty"`
	CreatedAt        string            `json:"created_at,omitempty"`
}
