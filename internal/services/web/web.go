package web

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Odvin/go-mock-http-server/internal/app"
	"github.com/Odvin/go-mock-http-server/pkg/mediator"
)

var ps = mediator.GetPubSub()

type HttpServer struct {
	api      app.API
	adr      int
	ver      string
	env      string
	ctx      context.Context
	shutdown context.CancelFunc
}

func Init(api app.API, adr int, ver, env string) *HttpServer {
	ctx, cancel := context.WithCancel(context.Background())

	return &HttpServer{
		api:      api,
		adr:      adr,
		ver:      ver,
		env:      env,
		ctx:      ctx,
		shutdown: cancel,
	}
}

func (hs *HttpServer) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", hs.adr),
		Handler:      hs.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func(srv *http.Server) {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}(srv)

	s := <-stop
	log.Printf("Shutdown signal received (signal: %s)", s.String())

	hs.api.StopCompanyUpdates()
	hs.shutdown()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}
