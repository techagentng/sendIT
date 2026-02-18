package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    PostmarkServerToken string
    FromEmail           string // "app@yourdomain.com"
    FromName            string // "Your App Name"
}

func Load() (*Config, error) {
    _ = godotenv.Load() // ignore error if .env not present

    token := os.Getenv("POSTMARK_SERVER_TOKEN")
    if token == "" {
        return nil, fmt.Errorf("POSTMARK_SERVER_TOKEN is required")
    }

    return &Config{
        PostmarkServerToken: token,
        FromEmail:           getEnvOrDefault("FROM_EMAIL", "hello@yourdomain.com"),
        FromName:            getEnvOrDefault("FROM_NAME", "Your App"),
    }, nil
}

func getEnvOrDefault(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}
