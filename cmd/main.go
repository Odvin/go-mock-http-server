package main

import (
	"log"

	"github.com/Odvin/go-mock-http-server/config"
	"github.com/Odvin/go-mock-http-server/internal/application"
	"github.com/Odvin/go-mock-http-server/internal/services/store"
	"github.com/Odvin/go-mock-http-server/internal/services/web"
)

func main() {
	cfg := config.InitConfig()

	storeAdapter := store.Init(100)

	appAdapter := application.Init(storeAdapter)

	webService := web.Init(appAdapter, cfg.Port, "1.0.0", cfg.Env)

	log.Printf("starting server on port :%d (env: %s)", cfg.Port, cfg.Env)

	err := webService.Serve()
	if err != nil {
		log.Panic(err)
	}
}
