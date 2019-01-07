package conekta

import (
	"errors"

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
func (client *Client) CreateCustomer(customer Customer) error {
	return errors.New("Not yet implemented")
}

// UpdateCustomer ...
func (client *Client) UpdateCustomer(customer Customer) error {
	return errors.New("Not yet implemented")
}

// DeleteCustomer ...
func (client *Client) DeleteCustomer(customer Customer) error {
	return errors.New("Not yet implemented")
}

// CreateOrder ...
func (client *Client) CreateOrder(order Order) error {
	return errors.New("Not yet implemented")
}

// UpdateOrder ...
func (client *Client) UpdateOrder(order Order) error {
	return errors.New("Not yet implemented")
}

// CaptureOrder ...
func (client *Client) CaptureOrder(order Order) error {
	return errors.New("Not yet implemented")
}

// RefundOrder ...
func (client *Client) RefundOrder(order Order) error {
	return errors.New("Not yet implemented")
}

// CreatePlan ...
func (client *Client) CreatePlan(plan Plan) error {
	return errors.New("Not yet implemented")
}

// UpdatePlan ...
func (client *Client) UpdatePlan(plan Plan) error {
	return errors.New("Not yet implemented")
}

// DeletePlan ...
func (client *Client) DeletePlan(plan Plan) error {
	return errors.New("Not yet implemented")
}

// CreateSubscription ...
func (client *Client) CreateSubscription(subscription Subscription) error {
	return errors.New("Not yet implemented")
}

// UpdateSubscription ...
func (client *Client) UpdateSubscription(subscription Subscription) error {
	return errors.New("Not yet implemented")
}

// PauseSubscription ...
func (client *Client) PauseSubscription(subscription Subscription) error {
	return errors.New("Not yet implemented")
}

// ResumeSubscription ...
func (client *Client) ResumeSubscription(subscription Subscription) error {
	return errors.New("Not yet implemented")
}

// CancelSubscription ...
func (client *Client) CancelSubscription(subscription Subscription) error {
	return errors.New("Not yet implemented")
}

// CreatePaymentSource ...
func (client *Client) CreatePaymentSource(paymentSource PaymentSource) error {
	return errors.New("Not yet implemented")
}

// UpdatePaymentSource ...
func (client *Client) UpdatePaymentSource(paymentSource PaymentSource) error {
	return errors.New("Not yet implemented")
}

// DeletePaymentSource ...
func (client *Client) DeletePaymentSource(paymentSource PaymentSource) error {
	return errors.New("Not yet implemented")
}

// CreateCharges ...
func (client *Client) CreateCharges(charge Charge) error { return errors.New("Not yet implemented") }

// CreateLineItem ...
func (client *Client) CreateLineItem(lineItem LineItem) error {
	return errors.New("Not yet implemented")
}

// UpdateLineItem ...
func (client *Client) UpdateLineItem(lineItem LineItem) error {
	return errors.New("Not yet implemented")
}

// DeleteLineItem ...
func (client *Client) DeleteLineItem(lineItem LineItem) error {
	return errors.New("Not yet implemented")
}

// CreateShippingLine ...
func (client *Client) CreateShippingLine(shippingLine ShippingLine) error {
	return errors.New("Not yet implemented")
}

// UpdateShippingLine ...
func (client *Client) UpdateShippingLine(shippingLine ShippingLine) error {
	return errors.New("Not yet implemented")
}

// DeleteShippingLine ...
func (client *Client) DeleteShippingLine(shippingLine ShippingLine) error {
	return errors.New("Not yet implemented")
}

// CreateCustomerShippingContact ...
func (client *Client) CreateCustomerShippingContact(shippingContact ShippingContact) error {
	return errors.New("Not yet implemented")
}

// CreateOrderShippingContact ...
func (client *Client) CreateOrderShippingContact(shippingContact ShippingContact) error {
	return errors.New("Not yet implemented")
}

// UpdateShippingContact ...
func (client *Client) UpdateShippingContact(shippingContact ShippingContact) error {
	return errors.New("Not yet implemented")
}

// DeleteShippingContact ...
func (client *Client) DeleteShippingContact(shippingContact ShippingContact) error {
	return errors.New("Not yet implemented")
}

// CreateDiscountLine ...
func (client *Client) CreateDiscountLine(discountLine DiscountLine) error {
	return errors.New("Not yet implemented")
}

// UpdateDiscountLine ...
func (client *Client) UpdateDiscountLine(discountLine DiscountLine) error {
	return errors.New("Not yet implemented")
}

// DeleteDiscountLine ...
func (client *Client) DeleteDiscountLine(discountLine DiscountLine) error {
	return errors.New("Not yet implemented")
}

// CreateTaxLine ...
func (client *Client) CreateTaxLine(taxLine TaxLine) error { return errors.New("Not yet implemented") }

// UpdateTaxLine ...
func (client *Client) UpdateTaxLine(taxLine TaxLine) error { return errors.New("Not yet implemented") }

// DeleteTaxLine ...
func (client *Client) DeleteTaxLine(taxLine TaxLine) error { return errors.New("Not yet implemented") }
