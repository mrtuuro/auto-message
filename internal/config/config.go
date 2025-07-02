package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	MongoURI       string
	DatabaseName   string
	CollectionName string
	Ctx            context.Context
	SecretKey      string
	WebhookURL     string
	WebhookKey     string
}

func NewConfig() *Config {
	if err := loadEnv(); err != nil {
		log.Printf("Error loading env: %v", err)
		return nil
	}

	ctx := context.Background()
	cfg := &Config{
		Port:           getEnvWithDefault("PORT", ":3000"),
		MongoURI:       getEnvWithDefault("MONGO_URI", ""),
		DatabaseName:   getEnvWithDefault("DATABASE_NAME", ""),
		CollectionName: getEnvWithDefault("COLLECTION_NAME", ""),
		Ctx:            ctx,
		SecretKey:      getEnvWithDefault("SECRET_KEY", ""),
		WebhookURL:     getEnvWithDefault("WEBHOOK_URL", ""),
		WebhookKey:     getEnvWithDefault("WEBHOOK_KEY", ""),
	}

	return cfg
}

func getEnvWithDefault(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func loadEnv() error {
	return godotenv.Load()
}
