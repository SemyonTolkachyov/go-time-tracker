package cmd

import (
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/app"
)

func RunHTTP() error {
	globalApp, err := app.GetGlobalApp()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err = globalApp.StartHTTPServer()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
