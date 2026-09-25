package config

import (
	"os"
)

type Config struct {
	AppPort    string
	AWSRegion  string
	SQSQueueURL string
}

func Load() Config {
	return Config{
		AppPort:     getEnv("APP_PORT", "8080"),
		AWSRegion:   getEnv("AWS_REGION", "ap-south-1"),
		SQSQueueURL: getEnv("SQS_QUEUE_URL", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}