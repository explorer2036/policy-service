package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Config structure for server
type Config struct {
	Address  string `json:"address"`
	DBHost   string `json:"db_host"`
	DBPort   string `json:"db_port"`
	DBUser   string `json:"db_user"`
	DBPasswd string `json:"db_passwd"`
	DBName   string `json:"db_name"`
}

// New returns the server config
func New() *Config {
	s := Config{
		Address:  ":8000",
		DBHost:   "localhost",
		DBPort:   "5432",
		DBUser:   "postgres",
		DBPasswd: "localdb",
		DBName:   "policies",
	}
	if address := os.Getenv("ADDRESS"); address != "" {
		s.Address = address
	}
	if host := os.Getenv("DB_HOST"); host != "" {
		s.DBHost = host
	}
	if port := os.Getenv("DB_PORT"); port != "" {
		s.DBPort = port
	}
	if user := os.Getenv("DB_USER"); user != "" {
		s.DBUser = user
	}
	if passwd := os.Getenv("DB_PASSWORD"); passwd != "" {
		s.DBPasswd = passwd
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		s.DBName = name
	}

	logrus.Infof("server config: %s", s)
	return &s
}
