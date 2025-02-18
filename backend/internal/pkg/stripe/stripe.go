package stripe

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/client"
)

type Stripe struct {
	client *client.API
}

func New() Stripe {
	stripe := &client.API{}
	stripe.Init("access_token", nil)

	return Stripe{
		client: stripe,
	}
}

func (s Stripe) CreateCustomer(email, name string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		Name:  stripe.String(name),
	}

	customer, err := s.client.Customers.New(params)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
