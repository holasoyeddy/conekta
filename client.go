package conekta

import (
	. "github.com/holasoyeddy/conekta/models"
)

// Client ...
type Client struct {
	BaseURI   string
	Locale    string
	Key       string
	Version   string
	requestor *requestor
}

// NewClient sets up a new client with the API key provided. The other values are set by default, but can be changed using the appropriate setters.
func NewClient(key string) *Client {

	client := &Client{
		BaseURI:   "https://api.conekta.io/",
		Locale:    "es",
		Key:       key,
		Version:   "2.0.0",
		requestor: newRequestor(),
	}

	return client
}

// CreateCustomer ...
func (client *Client) CreateCustomer(customer Customer) {}

// UpdateCustomer ...
func (client *Client) UpdateCustomer(customer Customer) {}

// DeleteCustomer ...
func (client *Client) DeleteCustomer(customer Customer) {}

// CreateOrder ...
func (client *Client) CreateOrder(order Order) {}

// UpdateOrder ...
func (client *Client) UpdateOrder(order Order) {}

// CaptureOrder ...
func (client *Client) CaptureOrder(order Order) {}

// RefundOrder ...
func (client *Client) RefundOrder(order Order) {}

// CreatePlan ...
func (client *Client) CreatePlan(plan Plan) {}

// UpdatePlan ...
func (client *Client) UpdatePlan(plan Plan) {}

// DeletePlan ...
func (client *Client) DeletePlan(plan Plan) {}

// CreateSubscription ...
func (client *Client) CreateSubscription(subscription Subscription) {}

// UpdateSubscription ...
func (client *Client) UpdateSubscription(subscription Subscription) {}

// PauseSubscription ...
func (client *Client) PauseSubscription(subscription Subscription) {}

// ResumeSubscription ...
func (client *Client) ResumeSubscription(subscription Subscription) {}

// CancelSubscription ...
func (client *Client) CancelSubscription(subscription Subscription) {}

// CreatePaymentSource ...
func (client *Client) CreatePaymentSource(paymentSource PaymentSource) {}

// UpdatePaymentSource ...
func (client *Client) UpdatePaymentSource(paymentSource PaymentSource) {}

// DeletePaymentSource ...
func (client *Client) DeletePaymentSource(paymentSource PaymentSource) {}

// CreateCharges ...
func (client *Client) CreateCharges(charge Charge) {}

// CreateLineItem ...
func (client *Client) CreateLineItem(lineItem LineItem) {}

// UpdateLineItem ...
func (client *Client) UpdateLineItem(lineItem LineItem) {}

// DeleteLineItem ...
func (client *Client) DeleteLineItem(lineItem LineItem) {}

// CreateShippingLine ...
func (client *Client) CreateShippingLine(shippingLine ShippingLine) {}

// UpdateShippingLine ...
func (client *Client) UpdateShippingLine(shippingLine ShippingLine) {}

// DeleteShippingLine ...
func (client *Client) DeleteShippingLine(shippingLine ShippingLine) {}

// CreateCustomerShippingContact ...
func (client *Client) CreateCustomerShippingContact(shippingContact ShippingContact) {}

// CreateOrderShippingContact ...
func (client *Client) CreateOrderShippingContact(shippingContact ShippingContact) {}

// UpdateShippingContact ...
func (client *Client) UpdateShippingContact(shippingContact ShippingContact) {}

// DeleteShippingContact ...
func (client *Client) DeleteShippingContact(shippingContact ShippingContact) {}

// CreateDiscountLine ...
func (client *Client) CreateDiscountLine(discountLine DiscountLine) {}

// UpdateDiscountLine ...
func (client *Client) UpdateDiscountLine(discountLine DiscountLine) {}

// DeleteDiscountLine ...
func (client *Client) DeleteDiscountLine(discountLine DiscountLine) {}

// CreateTaxLine ...
func (client *Client) CreateTaxLine(taxLine TaxLine) {}

// UpdateTaxLine ...
func (client *Client) UpdateTaxLine(taxLine TaxLine) {}

// DeleteTaxLine ...
func (client *Client) DeleteTaxLine(taxLine TaxLine) {}
