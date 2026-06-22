# Module 3: Managing the Cart (Orders & Models)

### Teaching Arc
- **Metaphor:** A detailed filing cabinet where every folder (Order) has specific sheets inside (Line Items, Addresses).
- **Opening hook:** An order is more than just a total amount; it's a collection of addresses, products, and calculations.
- **Key insight:** The core entities represent exactly how the data is structured in the database via GORM.
- **"Why should I care?":** Changing a JSON payload in a request means you likely need to update the Go struct model to match.

### Code Snippets (pre-extracted)

File: models/order.go (lines 20-30)
```go
// Order struct
type Order struct {
	ID              string       `json:"id"`
	LineItems       []LineItem   `json:"line_items"`
	SubTotal        uint64       `json:"sub_total"`
	Taxes           uint64       `json:"taxes"`
	Total           uint64       `json:"total"`
	Email           string       `json:"email"`
	ShippingAddress *Address     `json:"shipping_address,omitempty"`
	BillingAddress  *Address     `json:"billing_address,omitempty"`
}
```

### Interactive Elements

- [ ] **Visual File Tree** — Show `models/` directory with `order.go`, `line_item.go`, and `address.go`.
- [ ] **Drag-and-drop Quiz** — Items: "line_items", "shipping_address", "taxes". Targets: The matching struct field descriptions.

### Reference Files to Read
- `references/interactive-elements.md` -> "Visual File Tree", "Drag-and-Drop Matching"

### Connections
- **Previous module:** Routing & Middleware
- **Next module:** Payments
