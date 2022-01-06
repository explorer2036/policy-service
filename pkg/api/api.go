package api

import (
	"context"
	"fmt"
	"policy-server/pkg/config"
	"policy-server/pkg/db"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Server for license
type Server struct {
	e        *echo.Echo     // a fast web framework
	settings *config.Config // the server configuration
	handler  db.Repository  // handler for db operation
	validate *validator.Validate
	done     chan struct{} // server is done
}

// NewServer returns a http server
func NewServer(settings *config.Config) (*Server, error) {
	s := &Server{
		e:        echo.New(),
		settings: settings,
		validate: validator.New(),
		done:     make(chan struct{}),
	}
	s.e.HideBanner = true
	s.e.HidePort = true

	// init the database handler
	handler, err := db.NewHandler(settings)
	if err != nil {
		return nil, fmt.Errorf("new db handler: %w", err)
	}
	s.handler = handler

	s.e.POST("/policy", s.CreatePolicy)
	s.e.DELETE("/policy", s.DeletePolicy)
	s.e.PUT("/policy", s.UpdatePolicy)
	s.e.GET("/policy", s.QueryPolicy)

	return s, nil
}

// Start the http server
func (s *Server) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	// start the http server
	go func() {
		wg.Done()

		logrus.Infof("http server started on %s", s.settings.Address)
		if err := s.e.Start(s.settings.Address); err != nil {
			if !strings.Contains(err.Error(), "Server closed") {
				logrus.Errorf("start http server: %v", err)
			}
		}
	}()
}

// Stop the http server
func (s *Server) Stop() {
	// stop the jobs for server
	if s.done != nil {
		close(s.done)
	}

	// shut down the http server
	if s.e != nil {
		if err := s.e.Shutdown(context.Background()); err != nil {
			logrus.Errorf("shutdown http server: %v", err)
		}
	}
}
