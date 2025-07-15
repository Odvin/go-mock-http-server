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

	"github.com/Odvin/go-mock-http-server/internal/application"
)

type Web struct {
	api application.API
	adr int
	ver string
	env string
}

func Init(api application.API, adr int, ver, env string) *Web {
	return &Web{
		api: api,
		adr: adr,
		ver: ver,
		env: env,
	}
}

func (web *Web) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", web.adr),
		Handler:      web.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		log.Printf("stopping the server (signal: %s)", s.String())

		web.api.StopCompanyUpdates()

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	log.Printf("stopped server on port %s", srv.Addr)

	return nil
}
