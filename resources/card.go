package resources

// Card ...
type Card struct {
	TokenID   string  `json:"token_id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Number    string  `json:"number,omitempty"`
	ExpMonth  string  `json:"exp_month,omitempty"`
	ExpYear   string  `json:"exp_year,omitempty"`
	Reference string  `json:"reference,omitempty"`
	CVC       string  `json:"cvc,omitempty"`
	Last4     string  `json:"last4,omitempty"`
	Bin       string  `json:"bin,omitempty"`
	Brand     string  `json:"brand,omitempty"`
	Address   Address `json:"address,omitempty"`
}
