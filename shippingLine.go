package conekta

// ShippingLine ....
type ShippingLine struct {
	ID             string `json:"id,omitempty"`
	Amount         int    `json:"amount,omitempty"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	Carrier        string `json:"carrier,omitempty"`
	Method         string `json:"method,omitempty"`
	ParentID       string `json:"parent_id,omitempty"`
}
