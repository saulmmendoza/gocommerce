# Module 2: The API Layer (Routing & Middleware)

### Teaching Arc
- **Metaphor:** A busy reception desk where the receptionist checks your ID (Middleware) before handing you the room key (Context) and pointing you to your meeting (Handler).
- **Opening hook:** Before GoCommerce can process an order, it has to safely receive the request.
- **Key insight:** The router is wrapped in layers of middleware that inject useful context (like configuration and database connections) into the request.
- **"Why should I care?":** When building new features, you don't need to manually pass the database connection around; you grab it from the request context!

### Code Snippets (pre-extracted)

File: api/router.go (lines 60-70)
```go
func (m middlewareHandler) serve(next http.Handler, w http.ResponseWriter, r *http.Request) {
	ctx, err := m(w, r)
	if err != nil {
		handleError(err, w, r)
		return
	}
	if ctx != nil {
		r = r.WithContext(ctx)
	}
	next.ServeHTTP(w, r)
}
```

### Interactive Elements

- [ ] **Group chat animation** — actors: Client, Router, Middleware. Message flow: Client asks for `/settings`. Router says "Hold on". Middleware says "Checked ID, here is the DB Context". Router says "Proceed to Settings Handler".
- [ ] **Code↔English translation** — The `middlewareHandler.serve` function above.

### Reference Files to Read
- `references/interactive-elements.md` -> "Group Chat Animation", "Code ↔ English Translations"

### Connections
- **Previous module:** Intro & Architecture
- **Next module:** Cart & Models
