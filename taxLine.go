package conekta

// TaxLine ...
type TaxLine struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      string `json:"tracking_number,omitempty"`
	ParentID    string `json:"parent_id,omitempty"`
}
