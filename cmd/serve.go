package cmd

import (
	"os"
	"os/signal"
	"policy-server/pkg/api"
	"policy-server/pkg/config"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server",
	Run: func(cmd *cobra.Command, args []string) {
		// start the http server
		serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func serve() {
	var wg sync.WaitGroup

	settings := config.New()
	// start http servers
	s, err := api.NewServer(settings)
	if err != nil {
		logrus.Fatalf("new api server: %v", err)
	}
	s.Start(&wg)

	sig := make(chan os.Signal, 1024)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for o := range sig {
		logrus.Infof("receive signal: %v", o)

		start := time.Now()
		// stop the server to release the allocated resources
		s.Stop()

		// wait for goroutines done
		wg.Wait()

		logrus.Info("server is stopped")

		logrus.Infof("shut down takes time: %v", time.Since(start))
		return
	}
}
