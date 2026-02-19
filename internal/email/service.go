package email

import (
	"context"
	"fmt"

	"mailer/internal/postmark"
)

type Service struct {
	client postmark.ClientInterface
}

func NewService(client postmark.ClientInterface) *Service {
	return &Service{client: client}
}

func (s *Service) SendWelcomeEmail(ctx context.Context, to, name string) error {
	body := fmt.Sprintf(`<h1>Welcome, %s!</h1><p>We're excited to have you.</p>`, name)

	email := postmark.Email{
		To:         to,
		Subject:    "Welcome to Our App!",
		HTMLBody:   body,
		TextBody:   "Welcome!\nWe're excited to have you.",
		Tag:        "welcome",
		TrackOpens: true,
	}

	return s.client.Send(ctx, email)
}

func (s *Service) SendEmail(ctx context.Context, email postmark.Email) error {
	return s.client.Send(ctx, email)
}

// Add more methods: SendPasswordReset, SendOrderConfirmation, SendBatch, etc.
