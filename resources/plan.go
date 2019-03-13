package resources

// Plan is a template for a subscription. It allows you to define the amount and frequency with which you would like to bill your clients.
// For further information, visit https://developers.conekta.com/api?#plan
type Plan struct {
	ID              string `json:"id"`
	Object          string `json:"object"`
	LiveMode        string `json:"livemode"`
	CreatedAt       string `json:"created_at"`
	Name            string `json:"name"`
	Amount          int    `json:"amount"`
	Currency        string `json:"currency"`
	Interval        string `json:"interval"`
	Frequency       int    `json:"frquency"`
	ExpiryCount     int    `json:"expiry_count"`
	TrailPeriodDays int    `json:"trail_period_days"`
}
