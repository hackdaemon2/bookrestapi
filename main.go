package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	loaderInstance := GetInstanceConfigurationLoader()
	configuration, configurationError := loaderInstance.LoadConfiguration()

	if configurationError != nil {
		log.Fatal("Unable to load config file")
		return
	}

	db, dbError := loaderInstance.ConnectDatabase(configuration)

	if dbError != nil {
		log.Fatal("Unable to connect to database")
		return
	}

	routeHandler := loaderInstance.GetRouteHandler(db, configuration)

	server := &http.Server{
		Addr:           ":3000",
		Handler:        routeHandler,
		ReadTimeout:    time.Duration(configuration.ReadTimeout()) * time.Second,
		WriteTimeout:   time.Duration(configuration.WriteTimeout()) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
