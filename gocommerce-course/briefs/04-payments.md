# Module 4: Taking the Money (Payments Integration)

### Teaching Arc
- **Metaphor:** A universal credit card terminal that accepts any card, regardless of whether the bank behind it is Stripe or PayPal.
- **Opening hook:** An order is just a wishlist until the money actually moves.
- **Key insight:** GoCommerce uses an interface for payments, allowing you to easily swap or add providers like Stripe or PayPal.
- **"Why should I care?":** When a payment fails, the interface normalizes the error so the rest of the application doesn't have to care *which* provider failed.

### Code Snippets (pre-extracted)

File: payments/payments.go (lines 14-20)
```go
// Provider is an interface for payment providers
type Provider interface {
	Charge(*models.Order, *models.Transaction) error
	Refund(*models.Transaction) error
	Name() string
}
```

### Interactive Elements

- [ ] **Icon-Label Rows** — Showing the Payments interface vs the concrete Stripe and PayPal implementations.
- [ ] **Multiple-Choice Quiz** — Question about what happens if `Charge` fails. Options: order is deleted, transaction is marked failed, application crashes.

### Reference Files to Read
- `references/interactive-elements.md` -> "Icon-Label Rows", "Multiple-Choice Quizzes"

### Connections
- **Previous module:** Cart & Models
- **Next module:** Downloads & Coupons
