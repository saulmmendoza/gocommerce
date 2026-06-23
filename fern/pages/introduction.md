# Introduction

Welcome to GoCommerce!

GoCommerce is a small, Go-based API designed for static e-commerce sites. It handles the backend processes necessary for running a store, such as order management and payment processing.

It currently integrates with **Stripe** and **PayPal** for payments, and supports international pricing and VAT verification.

## Key Concepts

GoCommerce is built specifically for JAMstack or static sites. Since these sites don't have a traditional backend to handle database queries and business logic, GoCommerce acts as the missing piece for commerce capabilities.

### How It Works

1. **Product Metadata:** Each product you want to sell must have a unique URL where GoCommerce can find the metadata needed for calculating pricing and taxes.
2. **Metadata Format:** The metadata is embedded directly into your product pages using a script tag in JSON format:

```html
<script class="gocommerce-product" type="application/json">
{
  "sku": "my-product",
  "title": "My Product",
  "prices": [{"amount": "49.99", "currency": "USD"}],
  "type": "ebook"
}
</script>
```

The minimum required fields are `sku`, `title`, and at least one `price`. The default currency is USD if nothing else is specified.
