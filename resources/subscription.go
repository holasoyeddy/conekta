package resources

// Subscription ...
type Subscription struct {
	ID                string `json:"id,omitempty"`
	Status            string `json:"status,omitempty"`
	Plan              string `json:"plan,omitempty"`
	PlanID            string `json:"plan_id,omitempty"`
	CardID            string `json:"card_id,omitempty"`
	BillingCycleStart int    `json:"billing_cycle_start,omitempty"`
	BillingCycleEnd   int    `json:"billing_cycle_end,omitempty"`
	CreatedAt         string `json:"created_at,omitempty"`
	CustomerID        string `json:"customer_id,omitempty"`
}
