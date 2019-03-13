package resources

// TaxLine ...
type TaxLine struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Amount      string `json:"tracking_number"`
	ParentID    string `json:"parent_id"`
}
