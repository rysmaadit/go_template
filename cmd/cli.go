package cmd

import (
	"context"
	"example.com/m/config"
	"example.com/m/router"
	"example.com/m/service"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (c *Cli) Run() {
	log.SetLevel(log.InfoLevel)
	log.StandardLogger()
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", config.GetAppHost(), config.GetAppPort()),
		Handler: router.NewRouter(service.InstantiateDependencies()),
	}

	log.Println(fmt.Sprintf("starting application: %v on port: %v", config.GetAppName(), config.GetAppPort()))

	go listenAndServe(srv)
	waitForShutdown(srv)
}

func listenAndServe(server *http.Server) {
	err := server.ListenAndServe()

	if err != nil {
		log.WithField("error", err.Error()).Fatal("unable to serve")
	}
}

func waitForShutdown(server *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Warn("shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}

	log.Warn("shutdown complete")
}
