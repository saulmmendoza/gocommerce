# Module 6: Closing the Loop (Webhooks & Emails)

### Teaching Arc
- **Metaphor:** A post-event wrap-up where the organizer sends thank you notes (Emails) and tells the caterers they can go home (Webhooks).
- **Opening hook:** The transaction is paid, the data is saved, but the customer doesn't know that yet!
- **Key insight:** The `mailer` package handles templating and sending emails asynchronously so it doesn't block the HTTP request.
- **"Why should I care?":** If you need to trigger a custom lambda function or send a different kind of email, this is where you hook in.

### Code Snippets (pre-extracted)

File: mailer/mailer.go (lines 40-50)
```go
func (m *mailer) OrderConfirmation(order *models.Order) error {
	m.log.Debugf("Sending order confirmation for %v", order.ID)

	subject := m.config.Subjects.OrderConfirmation
	if subject == "" {
		subject = "Order Confirmation"
	}

	return m.sendEmail(order.Email, subject, "order_confirmation", order)
}
```

### Interactive Elements

- [ ] **Code↔English translation** — The `OrderConfirmation` function above.

### Reference Files to Read
- `references/interactive-elements.md` -> "Code ↔ English Translations"

### Connections
- **Previous module:** Downloads & Coupons
- **Next module:** None (End of course)
