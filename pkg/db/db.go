package db

import (
	"fmt"
	"policy-server/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	// DefaultMaxOpenConns - max open connections for database
	DefaultMaxOpenConns = 20
	// DefaultMaxIdleConns - max idle connections for database
	DefaultMaxIdleConns = 5
)

// Handler for the db operations
type Handler struct {
	settings *config.Config // database connection config
	db       *gorm.DB       // postgres database object
}

// NewHandler returns a new database operation handler
func NewHandler(settings *config.Config) (Repository, error) {
	s := &Handler{settings: settings}

	// init the database connections
	source := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		settings.DBHost,
		settings.DBPort,
		settings.DBUser,
		settings.DBName,
		settings.DBPasswd,
	)
	db, err := gorm.Open("postgres", source)
	if err != nil {
		return nil, fmt.Errorf("gorm open: %w", err)
	}
	db.DB().SetMaxOpenConns(DefaultMaxOpenConns)
	db.DB().SetMaxIdleConns(DefaultMaxIdleConns)
	if err := db.DB().Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}
	// init the db for handler
	s.db = db

	return s, nil
}

// Close make sure all the database connections are released
func (s *Handler) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
