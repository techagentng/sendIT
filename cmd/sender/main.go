package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"mailer/internal/config"
	"mailer/internal/email"
	"mailer/internal/postmark"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	pmClient := postmark.NewClient(cfg.PostmarkServerToken, cfg.FromName, cfg.FromEmail)

	emailService := email.NewService(pmClient)

	// Example usage
	err = emailService.SendWelcomeEmail(ctx, "test@residencelaplaya.com", "Alex")
	if err != nil {
		log.Printf("failed to send welcome email: %v", err)
		os.Exit(1)
	}

	log.Println("Welcome email sent successfully")
}
