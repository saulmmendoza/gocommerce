---
title: Multi Instances
description: How to configure and run GoCommerce in Multi-Instance mode
---

# Multi-Instance Mode

GoCommerce supports a Multi-Instance mode designed specifically for operators (like Netlify or other hosting providers) who want to host multiple e-commerce stores on a single GoCommerce deployment. Instead of spinning up a separate binary for every store, one GoCommerce API handles requests for many distinct sites securely.

This mode enables configuration at both a global level and an instance (per-tenant) level.

## Getting Started with Multi-Instance Mode

To activate multi-instance mode, you must set up specific global configuration variables.

1. **Enable the mode** by providing an operator token:
   ```env
   GOCOMMERCE_OPERATOR_TOKEN=your_secure_operator_token
   ```
   *Note: Providing an `OPERATOR_TOKEN` automatically enables multi-instance mode.*

2. **Database:** You still configure the global database connection normally (e.g., using `DATABASE_URL`). GoCommerce stores configuration for all instances within a specific `instances` table in this database.

### How Requests are Verified

In this mode, GoCommerce assumes that a proxy or edge routing layer (controlled by the operator) sits in front of it.

When a user interacts with a store, the operator's proxy intercepts the request and forwards it to GoCommerce. The proxy must include a JSON Web Token (JWT) signed with the `OPERATOR_TOKEN`. This JWT proves that the request is legitimate and tells GoCommerce *which* instance the request belongs to.

## Instance Configuration

Each instance has its own configuration stored in the database. This allows different stores to have entirely separate:
- Stripe and PayPal credentials
- SMTP email settings
- JWT secrets for user authentication
- Webhook URLs

### Global Fallbacks

GoCommerce uses a fallback mechanism. If an instance does not provide a specific configuration value (for example, it doesn't specify an `SMTP_HOST`), GoCommerce will fall back to using the global configuration defined in your environment variables or `.env` file.

## Managing Instances via API

GoCommerce provides a set of API endpoints for the operator to manage instances. These endpoints require a valid JWT signed by the `OPERATOR_TOKEN`.

### 1. Create an Instance

To provision a new store, send a `POST` request to `/instances`.

**Request Body:**
```json
{
  "id": "unique-instance-id-123",
  "config": {
    "site_url": "https://store.example.com",
    "jwt_secret": "instance_specific_jwt_secret",
    "payment": {
      "stripe": {
        "enabled": true,
        "secret_key": "sk_live_..."
      }
    }
  }
}
```

### 2. Get an Instance

To retrieve the current configuration for an instance, send a `GET` request to `/instances/{instance_id}`.

### 3. Update an Instance

To update an instance's configuration, send a `PUT` request to `/instances/{instance_id}`.

You should provide the updated configuration object. GoCommerce will update the stored settings for that specific tenant.

**Request Body:**
```json
{
  "config": {
    "site_url": "https://store.example.com",
    "jwt_secret": "instance_specific_jwt_secret",
    "payment": {
      "stripe": {
        "enabled": false
      },
      "paypal": {
        "enabled": true,
        "client_id": "...",
        "secret": "...",
        "env": "production"
      }
    }
  }
}
```

### 4. Delete an Instance

To remove an instance and its configuration, send a `DELETE` request to `/instances/{instance_id}`.

## Operator App Manifest

GoCommerce also provides a `GET /` endpoint at the root (when operating in multi-instance mode) that returns an App Manifest. This is useful for operators to programmatically determine the capabilities and required configuration fields of the GoCommerce microservice.
