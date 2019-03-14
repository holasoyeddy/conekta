package resources

// LineItem ...
type LineItem struct {
	ID            string `json:"id"`
	Object        string `json:"object"`
	Name          string `json:"name"`
	UnitPrice     int    `json:"unit_price"`
	Quantity      int    `json:"quantity"`
	ParentID      string `json:"parent_id"`
	AntifraudInfo struct {
	} `json:"antifraud_info"`
	Metadata struct {
	} `json:"metadata"`
}

// type LineItem struct {
// 	ID          string   `json:"id"`
// 	Name        string   `json:"name"`
// 	Description string   `json:"description"`
// 	UnitPrice   int      `json:"unit_price"`
// 	Quantity    int      `json:"quantity"`
// 	SKU         string   `json:"sku"`
// 	Shippable   bool     `json:"shippable"`
// 	Tags        []string `json:"tags"`
// 	Brand       string   `json:"brand"`
// 	Type        string   `json:"type"`
// 	ParentID    string   `json:"parent_id"`
// }
