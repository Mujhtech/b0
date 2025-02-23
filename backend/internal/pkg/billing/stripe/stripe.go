package stripe

import (
	"fmt"

	"github.com/mujhtech/b0/config"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/client"
)

type Stripe struct {
	client *client.API
	cfg    config.Stripe
}

func New(cfg *config.Config) Stripe {
	stripe := &client.API{}
	stripe.Init(cfg.Stripe.ApiKey, nil)

	return Stripe{
		client: stripe,
		cfg:    cfg.Stripe,
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

func (s Stripe) CreateCustomerSubscription(customerID, plan string) (string, error) {

	var priceId string

	switch plan {
	case "starter":
		priceId = s.cfg.StarterSubscriptionPriceID
	case "pro":
		priceId = s.cfg.ProSubscriptionPriceID
	case "scale":
		priceId = s.cfg.ScaleSubscriptionPriceID
	default:
		return "", fmt.Errorf("invalid plan")
	}

	createSubscriptionOptions := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(priceId),
			},
		},
	}

	sub, err := s.client.Subscriptions.New(createSubscriptionOptions)

	if err != nil {
		return "", err
	}

	return sub.ID, nil
}

func (s *Stripe) Portal(customerId string) (string, error) {

	portal, err := s.client.BillingPortalSessions.New(&stripe.BillingPortalSessionParams{
		Customer: stripe.String(customerId),
	})
	if err != nil {

		return "", err
	}

	return portal.URL, nil
}
