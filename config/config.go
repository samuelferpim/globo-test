package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string
	RedisURL    string
	RabbitMQURL string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("REDIS_URL", "redis://localhost:6379")
	viper.SetDefault("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		log.Println("Using default values")
	}

	config := &Config{
		ServerPort:  viper.GetString("SERVER_PORT"),
		RedisURL:    viper.GetString("REDIS_URL"),
		RabbitMQURL: viper.GetString("RABBITMQ_URL"),
	}

	log.Printf("Loaded configuration:")
	log.Printf("Server Port: %s", config.ServerPort)
	log.Printf("Redis URL: %s", config.RedisURL)
	log.Printf("RabbitMQ URL: %s", config.RabbitMQURL)

	return config, nil
}
