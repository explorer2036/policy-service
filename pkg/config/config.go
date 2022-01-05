package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Config structure for server
type Config struct {
	Address     string `yaml:"address"`
	PostgresURL string `yaml:"postgres_url"`
}

// New returns the server config
func New() *Config {
	s := Config{
		Address:     ":8000",
		PostgresURL: "postgres:123456@localhost:5432/policies",
	}
	if address := os.Getenv("ADDRESS"); address != "" {
		s.Address = address
	}
	if url := os.Getenv("POSTGRES_URL"); url != "" {
		s.PostgresURL = url
	}
	logrus.Infof("address: %s, postgres url: %s", s.Address, s.PostgresURL)
	return &s
}
