package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"mailer/internal/config"
	"mailer/internal/email"
	"mailer/internal/handlers"
	"mailer/internal/postmark"

	"github.com/gin-gonic/gin"
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
	emailHandler := handlers.NewEmailHandler(emailService)

	// Setup Gin router
	router := gin.Default()

	// Routes
	api := router.Group("/api/v1")
	{
		api.POST("/send", emailHandler.SendEmail)
		api.POST("/welcome", emailHandler.SendWelcome)
		api.GET("/health", emailHandler.HealthCheck)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")
	srv.Shutdown(context.Background())
	log.Println("Server stopped")
}
