package resources

// PaymentMethod ...
type PaymentMethod struct {
	ServiceName string `json:"service_name,omitempty"`
	TokenID     string `json:"token_id,omitempty"`
	Object      string `json:"object,omitempty"`
	Type        string `json:"type,omitempty"`
	ExpiresAt   int    `json:"expires_at,omitempty"`
	StoreName   string `json:"store_name,omitempty"`
	Reference   string `json:"reference,omitempty"`
}
