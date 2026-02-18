package postmark

import (
	"context"
	"fmt"

	"github.com/mrz1836/postmark"
)

type Email struct {
	To          string
	Subject     string
	HTMLBody    string
	TextBody    string
	Tag         string
	TrackOpens  bool
	Attachments []Attachment
}

type Attachment struct {
	Name        string
	Content     []byte
	ContentType string
}

type ClientInterface interface {
	Send(ctx context.Context, e Email) error
}

type Client struct {
	pm        *postmark.Client
	fromName  string
	fromEmail string
}

func NewClient(serverToken, fromName, fromEmail string) *Client {
	return &Client{
		pm:        postmark.NewClient(serverToken, ""),
		fromName:  fromName,
		fromEmail: fromEmail,
	}
}

func (c *Client) Send(ctx context.Context, e Email) error {
	pmEmail := postmark.Email{
		From:       fmt.Sprintf("%s <%s>", c.fromName, c.fromEmail),
		To:         e.To,
		Subject:    e.Subject,
		HTMLBody:   e.HTMLBody,
		TextBody:   e.TextBody,
		Tag:        e.Tag,
		TrackOpens: e.TrackOpens,
	}

	resp, err := c.pm.SendEmail(ctx, pmEmail)
	if err != nil {
		return fmt.Errorf("postmark send failed: %w", err)
	}

	if resp.ErrorCode != 0 {
		return fmt.Errorf("postmark error %d: %s", resp.ErrorCode, resp.Message)
	}

	return nil
}
