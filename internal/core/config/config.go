package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Addr           string
	DatabaseURI    string
	AccrualAddress string
	Secret         string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.Addr, "a", "localhost:8080", "server address")
	flag.StringVar(&cfg.DatabaseURI, "d", "", "database address")
	flag.StringVar(&cfg.AccrualAddress, "r", "localhost:8081", "accrual system address")

	flag.Parse()

	if envAddr := os.Getenv("RUN_ADDRESS"); envAddr != "" {
		cfg.Addr = envAddr
	}

	if envDsn := os.Getenv("DATABASE_URI"); envDsn != "" {
		cfg.DatabaseURI = envDsn
	}

	if envAsa := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAsa != "" {
		cfg.AccrualAddress = envAsa
	}

	if cfg.DatabaseURI == "" {
		return nil, fmt.Errorf("empty Database URI")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		cfg.Secret = "secret"
	} else {
		cfg.Secret = secret
	}

	return cfg, nil
}
