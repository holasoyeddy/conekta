package conekta

// ShippingLine ....
type ShippingLine struct {
	ID             string `json:"id"`
	Amount         int    `json:"amount"`
	TrackingNumber string `json:"tracking_number"`
	Carrier        string `json:"carrier"`
	Method         string `json:"method"`
	ParentID       string `json:"parent_id"`
}
