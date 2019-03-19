package conekta_test

import (
	"testing"

	"github.com/holasoyeddy/conekta"
)

var client *conekta.Client

func TestCreateOneTimePayment(t *testing.T) {

	// Initialize your Conekta client.
	var err error

	client, err = conekta.NewClient("key_eYvWV7gSDkNYXsmr")

	// Check for client init error.
	if err != nil || client == nil {
		t.Fatal(err)
	}

	// Create an order and add the token.
	order := conekta.Order{
		LineItems: []conekta.LineItem{
			conekta.LineItem{
				Name:      "Tacos",
				UnitPrice: 1000,
				Quantity:  120,
			},
		},
		ShippingLines: []conekta.ShippingLine{
			conekta.ShippingLine{
				Amount:  1500,
				Carrier: "FEDEX",
			},
		},
		Currency: "MXN",
		CustomerInfo: conekta.Customer{
			Name:  "Fulanito Pérez",
			Email: "fulanito@conekta.com",
			Phone: "+52181818181",
		},
		ShippingContact: conekta.ShippingContact{
			Address: conekta.Address{
				Street1:    "Calle 123, Int 2",
				PostalCode: "06100",
				Country:    "MX",
			},
		},
		Charges: []conekta.Charge{
			conekta.Charge{
				PaymentMethod: conekta.PaymentMethod{
					Type:    "card",
					TokenID: "tok_test_visa_4242",
				},
			},
		},
	}

	// Attempt to create the order.
	if _, err := client.Create(order); err != nil {
		t.Fatal(err)
	}
}
