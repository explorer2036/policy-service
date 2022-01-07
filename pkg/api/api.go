package api

import (
	"context"
	"fmt"
	"policy-service/pkg/config"
	"policy-service/pkg/db"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Server for license
type Server struct {
	e        *echo.Echo          // a fast web framework
	settings *config.Config      // the server configuration
	handler  db.Repository       // handler for db operation
	validate *validator.Validate // validate the request fields
	done     chan struct{}       // server is done
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

	s.e.POST("/benchmark", s.CreateBenchmark)
	s.e.DELETE("/benchmark", s.DeleteBenchmark)
	s.e.PUT("/benchmark", s.UpdateBenchmark)
	s.e.GET("/benchmark", s.QueryBenchmark)
	s.e.GET("/benchmarks", s.QueryBenchmarks)

	s.e.POST("/policy", s.CreatePolicy)
	s.e.DELETE("/policy", s.DeletePolicy)
	s.e.PUT("/policy", s.UpdatePolicy)
	s.e.GET("/policy", s.QueryPolicy)
	s.e.GET("/policies", s.QueryPolicies)

	s.e.POST("/tag", s.CreateTag)
	s.e.GET("/tag", s.QueryTag)
	s.e.GET("/tags", s.QueryTags)

	s.e.POST("/provider", s.CreateProvider)
	s.e.PUT("/provider", s.UpdateProvider)
	s.e.GET("/provider", s.QueryProvider)
	s.e.GET("/providers", s.QueryProviders)

	s.e.POST("/provider_type", s.CreateProviderType)
	s.e.PUT("/provider_type", s.UpdateProviderType)
	s.e.GET("/provider_type", s.QueryProviderType)
	s.e.GET("/provider_types", s.QueryProviderTypes)

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
