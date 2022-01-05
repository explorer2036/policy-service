package api

import (
	"context"
	"fmt"
	"policy-server/pkg/config"
	"policy-server/pkg/db"
	"strings"
	"sync"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Server for license
type Server struct {
	e        *echo.Echo     // a fast web framework
	settings *config.Config // the server configuration
	handler  *db.Handler    // handler for db operation
	done     chan struct{}  // server is done
}

// NewServer returns a http server
func NewServer(settings *config.Config) (*Server, error) {
	s := &Server{
		e:        echo.New(),
		settings: settings,
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

	s.e.POST("/policy", s.createPolicy)
	s.e.DELETE("/policy", s.deletePolicy)
	s.e.PUT("/policy", s.updatePolicy)
	s.e.GET("/policy", s.queryPolicy)

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
				panic(err)
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
