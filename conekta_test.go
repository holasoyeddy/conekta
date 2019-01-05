package conekta

import "testing"

func TestNewClient(t *testing.T) {

	client := NewClient("key_eYvWV7gSDkNYXsmr")

	if client == nil {
		t.Error("Client is null")
	}
}
