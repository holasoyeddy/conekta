package resources

// PaymentMethod ...
type PaymentMethod struct {
	Object                 string `json:"_object"`
	Type                   string `json:"_type"`
	Name                   string `json:"name"`
	Bank                   string `json:"bank"`
	ReceivingAccountNumber string `json:"receiving_account_number"`
	ReceivingAccountBank   string `json:"receiving_account_bank"`
	CLABE                  string `json:"clabe"`
	ExpMonth               string `json:"exp_month"`
	ExpYear                string `json:"exp_year"`
	AuthCode               string `json:"auth_code"`
	Last4                  string `json:"last4"`
	Brand                  string `json:"brand"`
	Issuer                 string `json:"issuer"`
	AccountType            string `json:"account_type"`
	Country                string `json:"country"`
	ServiceName            string `json:"service_name"`
	ExpiresAt              string `json:"expires_at"`
	StoreName              string `json:"store_name"`
	Reference              string `json:"reference"`
}
