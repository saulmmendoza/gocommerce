# Module 1: The GoCommerce Checkout Flow (Intro & Architecture)

### Teaching Arc
- **Metaphor:** A checkout counter at a specialized boutique where one clerk (GoCommerce) coordinates between the customer, the bank, and the stockroom.
- **Opening hook:** Have you ever wondered what happens under the hood when a user clicks "Buy" on a static website?
- **Key insight:** GoCommerce bridges the gap between static pages and dynamic payments/fulfillment.
- **"Why should I care?":** Understanding the data flow makes debugging "failed to charge" or "didn't receive email" bugs much easier.

### Code Snippets (pre-extracted)

File: api/api.go (lines 40-49)
```go
// API is the main REST API
type API struct {
	handler http.Handler
	db      *gorm.DB
	config  *conf.Configuration
	mailer  mailer.Mailer
	version string
	ext     *extensions
}
```

### Interactive Elements

- [ ] **Data flow animation** — actors: Client, GoCommerce, Stripe, DB, Mailer. Steps: 1. Client sends order 2. GoCommerce calculates tax 3. GoCommerce calls Stripe 4. GoCommerce saves to DB 5. Mailer sends receipt.

### Reference Files to Read
- `references/interactive-elements.md` -> "Message Flow Visualization"

### Connections
- **Previous module:** None
- **Next module:** Routing & Middleware
