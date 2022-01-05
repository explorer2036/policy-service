package api

import (
	"context"
	"policy-server/pkg/config"
	"strings"
	"sync"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Server for license
type Server struct {
	e        *echo.Echo     // a fast web framework
	settings *config.Config // the server configuration
	// handler  *db.Handler    // handler for db operation
	done chan struct{} // server is done
}

// NewServer returns a http server
func NewServer(settings *config.Config) *Server {
	s := &Server{
		e:        echo.New(),
		settings: settings,
		done:     make(chan struct{}),
	}
	s.e.HideBanner = true
	s.e.HidePort = true

	// init the database handler
	// s.handler = db.NewHandler(settings)

	// s.e.POST("/licenses/create", s.CreateHandler)
	// s.e.GET("/licenses/list", s.ListHandler)
	// s.e.GET("/licenses/:id/delete", s.DeleteHandler)
	// s.e.GET("/licenses/:id", s.GetHandler)
	// s.e.GET("/licenses/:id/activate", s.ActivateHandler)
	// s.e.GET("/licenses/:id/deactivate", s.DeactivateHandler)
	// s.e.POST("/licenses/verify", s.VerifyHandler)
	// s.e.GET("/license/ping", s.PingHandler)

	return s
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
