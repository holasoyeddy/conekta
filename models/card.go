package conekta

// Card ...
type Card struct {
	TokenID   string  `json:"token_id"`
	Name      string  `json:"name"`
	Number    string  `json:"number"`
	ExpMonth  string  `json:"exp_month"`
	ExpYear   string  `json:"exp_year"`
	Reference string  `json:"reference"`
	CVC       string  `json:"cvc"`
	Last4     string  `json:"last4"`
	Bin       string  `json:"bin"`
	Brand     string  `json:"brand"`
	Address   Address `json:"address"`
}
