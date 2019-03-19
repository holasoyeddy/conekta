package resources

// LineItem ...
type LineItem struct {
	ID            string `json:"id,omitempty"`
	Object        string `json:"object,omitempty"`
	Name          string `json:"name,omitempty"`
	UnitPrice     int    `json:"unit_price,omitempty"`
	Quantity      int    `json:"quantity,omitempty"`
	ParentID      string `json:"parent_id,omitempty"`
	AntifraudInfo struct {
	} `json:"antifraud_info,omitempty"`
	Metadata struct {
	} `json:"metadata,omitempty"`
}
