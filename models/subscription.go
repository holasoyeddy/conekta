package conekta

// Subscription ...
type Subscription struct {
	ID                string `json:"id"`
	Status            string `json:"status"`
	Plan              string `json:"plan"`
	PlanID            string `json:"plan_id"`
	CardID            string `json:"card_id"`
	BillingCycleStart int    `json:"billing_cycle_start"`
	BillingCycleEnd   int    `json:"billing_cycle_end"`
	CreatedAt         string `json:"created_at"`
	CustomerID        string `json:"customer_id"`
}
