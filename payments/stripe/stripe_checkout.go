package stripe

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/netlify/gocommerce/models"
	"github.com/netlify/gocommerce/payments"
	"github.com/sirupsen/logrus"
	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/webhook"
)

func (s *stripePaymentProvider) NewCheckouter(ctx context.Context, r *http.Request, log logrus.FieldLogger) (payments.Checkouter, error) {
	// Read body manually so we can extract raw params and parse what we need, or let the user pass standard params
	var params stripe.CheckoutSessionParams
	bod, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if len(bod) > 0 {
		var rawParams map[string]interface{}
		err = json.Unmarshal(bod, &rawParams)
		if err != nil {
			return nil, err
		}

		if successURL, ok := rawParams["success_url"].(string); ok {
			params.SuccessURL = stripe.String(successURL)
		}
		if cancelURL, ok := rawParams["cancel_url"].(string); ok {
			params.CancelURL = stripe.String(cancelURL)
		}
	}

	return func(amount uint64, currency string, order *models.Order) (string, string, error) {
		stripe.Key = s.config.SecretKey

		if len(params.LineItems) == 0 {
			params.LineItems = []*stripe.CheckoutSessionLineItemParams{
				{
					PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
						Currency: stripe.String(currency),
						ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
							Name: stripe.String(fmt.Sprintf("Order %s", order.ID)),
						},
						UnitAmount: stripe.Int64(int64(amount)),
					},
					Quantity: stripe.Int64(1),
				},
			}
		}

		if params.Mode == nil {
			params.Mode = stripe.String(string(stripe.CheckoutSessionModePayment))
		}

		if params.ClientReferenceID == nil {
			params.ClientReferenceID = stripe.String(order.ID)
		}

		if params.CustomerEmail == nil && order.Email != "" {
			params.CustomerEmail = stripe.String(order.Email)
		}

		if params.Metadata == nil {
			params.Metadata = map[string]string{}
		}
		params.Metadata["order_id"] = order.ID
		params.Metadata["instance_id"] = order.InstanceID

		checkoutSession, err := session.New(&params)
		if err != nil {
			return "", "", err
		}
		return checkoutSession.ID, checkoutSession.URL, nil
	}, nil
}

func (s *stripePaymentProvider) WebhookHandler(ctx context.Context, r *http.Request, log logrus.FieldLogger) (*payments.WebhookResult, error) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), s.config.WebhookSecret)
	if err != nil {
		log.WithError(err).Error("Error verifying webhook signature")
		return nil, err
	}

	switch event.Type {
	case "checkout.session.completed", "checkout.session.async_payment_succeeded":
		var sess stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &sess)
		if err != nil {
			return nil, err
		}

		if sess.PaymentStatus == stripe.CheckoutSessionPaymentStatusPaid {
			orderID := sess.Metadata["order_id"]
			instanceID := sess.Metadata["instance_id"]
			if orderID == "" {
				log.Warn("Stripe checkout session missing order_id metadata")
				return nil, nil // Ignored event
			}

			return &payments.WebhookResult{
				OrderID:       orderID,
				InstanceID:    instanceID,
				TransactionID: sess.ID,
				Amount:        uint64(sess.AmountTotal),
				Currency:      string(sess.Currency),
				Status:        models.PaidState,
			}, nil
		}
	}

	return nil, nil // Return nil for unhandled events
}
