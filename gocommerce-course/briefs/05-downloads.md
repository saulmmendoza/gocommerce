# Module 5: Delivering the Goods (Downloads & Coupons)

### Teaching Arc
- **Metaphor:** A secure vault where a bouncer verifies your ticket (JWT) before handing you the digital product.
- **Opening hook:** How do you securely sell a PDF if your site is just static HTML?
- **Key insight:** GoCommerce uses JWTs (JSON Web Tokens) as secure, temporary tickets that grant access to downloads.
- **"Why should I care?":** When configuring digital products, you need to know how the signed URLs are generated and validated.

### Code Snippets (pre-extracted)

File: api/download.go (lines 80-92)
```go
// Create a new token object, specifying signing method and the claims
// you would like it to contain.
claims := jwt.StandardClaims{
	Id:        dl.ID,
	Subject:   "download",
	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// Sign and get the complete encoded token as a string using the secret
tokenString, err := token.SignedString([]byte(config.JWT.Secret))
```

### Interactive Elements

- [ ] **Numbered Step Cards** — Showing the process of claiming a download. 1. Buy product. 2. Get email with link. 3. Click link (contains JWT). 4. GoCommerce validates JWT. 5. File is streamed.

### Reference Files to Read
- `references/interactive-elements.md` -> "Numbered Step Cards"

### Connections
- **Previous module:** Payments
- **Next module:** Webhooks & Emails
