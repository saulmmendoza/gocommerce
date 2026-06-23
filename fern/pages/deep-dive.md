# Deep Dive

This section covers more advanced configurations and features of GoCommerce, including tax calculations, payment integrations, and webhooks.

## VAT, Countries, and Regions

GoCommerce can handle VAT calculations based on a combination of product types and user locations (billing address).

To enable this, GoCommerce looks for a settings file hosted on your static site at: `https://your-site.com/gocommerce/settings.json`.

This file is optional, but it enables advanced tax verification features.

### Settings.json Example

```json
{
  "taxes": [{
    "percentage": 20,
    "product_types": ["ebook"],
    "countries": ["Austria", "Bulgaria", "Estonia", "France", "Gibraltar", "Slovakia", "United Kingdom"]
  }, {
    "percentage": 7,
    "product_types": ["book"],
    "countries": ["Austria", "Belgium", "Bulgaria", "Croatia", "Cyprus", "Denmark", "Estonia"]
  }]
}
```

### How Verification Works

1. A user attempts an order.
2. The user's cart contains an item with `type: "ebook"` and their billing address is in `"Austria"`.
3. GoCommerce reads your `settings.json` and sees that an `ebook` in `Austria` requires a 20% tax.
4. It verifies that a 20% tax was included in the product total before proceeding to charge the customer via the payment provider.

*Note: You must also implement these tax calculations on the client side so the user sees the correct total before checkout.*

## Payment Providers

GoCommerce supports multiple payment gateways. You configure them via environment variables.

### Stripe

To enable Stripe, you need to provide your secret key:

```env
PAYMENT_STRIPE_ENABLED=true
PAYMENT_STRIPE_SECRET_KEY=sk_test_your_secret_key
```

### PayPal

For PayPal, provide your OAuth credentials and specify the environment (`sandbox` or `production`):

```env
PAYMENT_PAYPAL_ENABLED=true
PAYMENT_PAYPAL_CLIENT_ID=your_client_id
PAYMENT_PAYPAL_SECRET=your_secret
PAYMENT_PAYPAL_ENV=sandbox
```

## Webhooks

GoCommerce can notify external services when specific events occur by sending an HTTP POST request (webhook).

You can configure webhook URLs for the following events:

```env
WEBHOOKS_ORDER=https://example.com/webhooks/order
WEBHOOKS_PAYMENT=https://example.com/webhooks/payment
WEBHOOKS_UPDATE=https://example.com/webhooks/update
WEBHOOKS_REFUND=https://example.com/webhooks/refund
```

### Webhook Security

To verify that the webhook came from GoCommerce, set a secret:

```env
WEBHOOKS_SECRET=my_webhook_secret
```

GoCommerce will sign the webhook payload with a JSON Web Token (JWT) using this secret and include it in the `X-Commerce-Signature` header. Your external service can decode and verify this JWT to ensure authenticity.

## Multi-Instance Mode

GoCommerce was originally designed to run on Netlify's infrastructure in a multi-tenant setup. While you generally don't need this for a single store, if you are acting as an operator (hosting multiple stores on a single GoCommerce deployment), you can enable Multi-Instance mode.

In this mode, GoCommerce relies on a `OPERATOR_TOKEN` to verify requests proxied from the operator. It also exposes `/instances` endpoints to manage configuration per-tenant.
