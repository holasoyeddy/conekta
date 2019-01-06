package conekta

// FiscalEntity ...
type FiscalEntity struct {
	ID          string  `json:"id"`
	TaxID       string  `json:"tax_id"`
	CompanyName string  `json:"company_name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     Address `json:"address"`
	ParentID    string  `json:"parent_id"`
}
