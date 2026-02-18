# Mailer - Powerful ESMTP Server with Postmark Integration

A clean, scalable Go application for sending emails via Postmark using modern architecture patterns.

## Project Structure

```
mailer/
├── cmd/
│   └── sender/
│       └── main.go               # entry point – only wires things together
├── internal/
│   ├── config/
│   │   └── config.go             # loads env vars, validates required keys
│   ├── email/
│   │   ├── email.go              # domain models (Email struct, etc.)
│   │   └── service.go            # business logic – SendEmail, SendBatch, etc.
│   └── postmark/
│       └── client.go             # wrapper / adapter around the 3rd-party library
├── pkg/                          # (optional) reusable non-business code
├── .env                          # local development (never commit!)
├── .env.example                  # template for others
├── go.mod
├── go.sum
└── README.md
```

## Architecture

This project follows Clean Architecture principles with clear separation of concerns:

- **cmd/sender/main.go**: CLI entry point with minimal logic
- **internal/email/service.go**: Business logic and use cases
- **internal/postmark/client.go**: Infrastructure adapter for Postmark API
- **internal/config/config.go**: Environment configuration management

## Setup

1. Copy the environment template:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` with your Postmark credentials:
   ```
   POSTMARK_SERVER_TOKEN=your_actual_token
   FROM_EMAIL=your-email@domain.com
   FROM_NAME=Your App Name
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run cmd/sender/main.go
   ```

## Usage

The service currently includes a welcome email example. You can extend it with additional email types like:

- Password reset emails
- Order confirmations
- Batch email sending
- Template-based emails

## Dependencies

- `github.com/joho/godotenv` - Environment variable management
- `github.com/mrz1836/postmark` - Postmark API client

## Extending the Service

To add new email types, simply add methods to `internal/email/service.go` following the same pattern as `SendWelcomeEmail`.
