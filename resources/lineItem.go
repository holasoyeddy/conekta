package resources

// LineItem ...
type LineItem struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	UnitPrice   int      `json:"unit_price"`
	Quantity    int      `json:"quantity"`
	SKU         string   `json:"sku"`
	Shippable   bool     `json:"shippable"`
	Tags        []string `json:"tags"`
	Brand       string   `json:"brand"`
	Type        string   `json:"type"`
	ParentID    string   `json:"parent_id"`
}
