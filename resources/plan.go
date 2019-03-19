package resources

// Plan is a template for a subscription. It allows you to define the amount and frequency with which you would like to bill your clients.
// For further information, visit https://developers.conekta.com/api?#plan
type Plan struct {
	ID              string `json:"id,omitempty"`
	Object          string `json:"object,omitempty"`
	LiveMode        string `json:"livemode,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	Name            string `json:"name,omitempty"`
	Amount          int    `json:"amount,omitempty"`
	Currency        string `json:"currency,omitempty"`
	Interval        string `json:"interval,omitempty"`
	Frequency       int    `json:"frquency,omitempty"`
	ExpiryCount     int    `json:"expiry_count,omitempty"`
	TrailPeriodDays int    `json:"trail_period_days,omitempty"`
}
