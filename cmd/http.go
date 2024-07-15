package main

import (
	log "github.com/sirupsen/logrus"
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
	globalApp, err := app.GetGlobalApp()
	if err != nil {
		log.Error(err.Error())
		return
	}
	err = globalApp.StartHTTPServer()
	if err != nil {
		log.Error(err.Error())
		return
	}
	//mux := http.NewServeMux()
	//mux.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
	//	q := request.URL.Query()
	//	ps := q.Get("passportNumber")
	//	pn := q.Get("passportSerie")
	//	/*res := UserInfo{
	//		Name:       ps,
	//		Surname:    pn,
	//		Patronymic: pn + " patronymic",
	//		Address:    request.URL.String(),
	//	}*/
	//	writer.Header().Set("Content-Type", "application/json")
	//	writer.WriteHeader(http.StatusOK)
	//	json.NewEncoder(writer).Encode(ps + pn)
	//	//fmt.Fprintf(writer, res)
	//})
	//log.Println("Listen serve success")
	//http.ListenAndServe(":8080", mux)
}
