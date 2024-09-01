package main

import (
	log "github.com/sirupsen/logrus"
	"go-time-tracker/cmd"
	"go-time-tracker/internal/app"
)

// @title Go Time Tracker API
// @version 1.0
// @description API Server for time tracker Application

func main() {
	err := app.InitApp()
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("App initialized")
	err = cmd.RunHTTP()
	if err != nil {
		log.Error(err)
	}
}
