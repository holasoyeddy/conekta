package conekta

// DiscountLine ...
type DiscountLine struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Type     string `json:"type"`
	Amount   string `json:"amount"`
	ParentID string `json:"parent_id"`
}
