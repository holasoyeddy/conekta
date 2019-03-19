package conekta

// FiscalEntity ...
type FiscalEntity struct {
	ID          string  `json:"id,omitempty"`
	TaxID       string  `json:"tax_id,omitempty"`
	CompanyName string  `json:"company_name,omitempty"`
	Email       string  `json:"email,omitempty"`
	Phone       string  `json:"phone,omitempty"`
	Address     Address `json:"address,omitempty"`
	ParentID    string  `json:"parent_id,omitempty"`
}
