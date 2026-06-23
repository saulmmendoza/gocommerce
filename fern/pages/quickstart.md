# Quickstart

This guide will help you get GoCommerce up and running locally for development and testing.

## Prerequisites

- **Go**: Version 1.11 or higher (GoCommerce uses Go modules).
- **Database**: SQLite3 (easiest for local dev), MySQL, or PostgreSQL.

## 1. Installation

First, clone the repository:

```bash
git clone https://github.com/netlify/gocommerce.git
cd gocommerce
```

Build the binary:

```bash
go build ./...
```

## 2. Configuration

GoCommerce is configured using an `.env` file or environment variables. The easiest way to get started is to use the provided example.

Copy the example configuration file:

```bash
cp example.env .env
```

Open `.env` in your text editor and ensure the following essential variables are set:

```env
GOCOMMERCE_SITE_URL=http://localhost:3000/ # The URL of your frontend
GOCOMMERCE_JWT_SECRET=supersecretvalue # Change this in production
GOCOMMERCE_DB_DRIVER=sqlite3
DATABASE_URL=gotrue.db
GOCOMMERCE_DB_AUTOMIGRATE=true
```

## 3. Running the Server

Once configured, you can run the GoCommerce backend:

```bash
go run main.go
```

By default, the API will listen on `localhost:8080`.

## 4. Frontend Integration

To interact with your GoCommerce backend from your static site, you can use the [commerce-js](https://github.com/netlify/netlify-commerce-js) client library.

Include it in your static site project and configure it to point to your local GoCommerce instance.

## Next Steps

Now that you have GoCommerce running, you can:
- Read the **Deep Dive** to understand VAT and advanced configurations.
- Explore the **API Reference** to integrate with the backend endpoints.
