package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config holds the credentials to establish a connection with the db
type Config struct {
	DbName   string `env:"DBNAME,required"`
	User     string `env:"DBUSER,required"`
	Password string `env:"PASSWORD,required"`
	Host     string `env:"HOST,required"`
	Port     string `env:"PORT,required"`
}

// Loads the env and returns a db configuration
func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		log.Fatalf("Error parsing env to Config struct: %s", err)
		return nil, err
	}

	return cfg, nil
}
