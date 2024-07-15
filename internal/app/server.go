package app

import (
	"github.com/hanagantig/gracy"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/handler/http/api"
	v1 "go-time-tracker/internal/handler/http/api/v1"
)

func (a *App) StartHTTPServer() error {
	go func() {
		a.startHTTPServer()
	}()

	err := gracy.Wait()
	if err != nil {
		log.Error("Failed to gracefully shutdown server", err.Error())
		return err
	}
	log.Info("Server gracefully stopped")
	return nil
}

func (a *App) startHTTPServer() {
	handler := v1.NewHandler(a.c.GetService())

	router := api.NewRouter()
	router.
		WithSwagger().
		WithHandler(handler)

	srv := api.NewServer(a.cfg.HTTP)
	srv.RegisterRoutes(router)

	gracy.AddCallback(func() error {
		return srv.Stop()
	})

	log.Infof("Starting HTTP server at %s:%s", a.cfg.HTTP.Host, a.cfg.HTTP.Port)
	err := srv.Start()
	if err != nil {
		log.Fatalf("Fail to start %s http server: %s", a.cfg.App.Name, err.Error())
	}
}
