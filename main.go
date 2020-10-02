package main

import (
	"context"
	"os/signal"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	"github.com/saravase/golang_mux_swagger/plant-api/data"
	"github.com/saravase/golang_mux_swagger/plant-api/handlers"

	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// New creates a new plant-api Logger.
	logger := log.New(os.Stdout, "product-plant-api", log.LstdFlags)

	// Initialize the Validation struct properties
	validation := data.NewValidation()

	// Initialize the plant struct properties
	plantHandler := handlers.NewPlant(logger, validation)

	// NewRouter returns a new gorilla mux router instance
	gorillaMux := mux.NewRouter()

	/*
		Subrouter creates a subrouter for the route
		It will test the inner routes only if the parent route matched
	*/

	// Get subrouter
	getRouter := gorillaMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/plant", plantHandler.GetPlants)
	getRouter.HandleFunc("/plant/{id:[0-9]+}", plantHandler.GetPlant)

	// Post subrouter
	postRouter := gorillaMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/plant", plantHandler.AddPlant)
	postRouter.Use(plantHandler.PlantValidationMiddleware)

	// Put subrouter
	putRouter := gorillaMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/plant/{id:[0-9]+}", plantHandler.UpdatePlant)
	putRouter.Use(plantHandler.PlantValidationMiddleware)

	// Delete subrouter
	deleteRouter := gorillaMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/plant/{id:[0-9]+}", plantHandler.DeletePlant)

	reDocOptions := middleware.RedocOpts{
		SpecURL: "/swagger.yaml",
	}
	swaggerMiddleware := middleware.Redoc(reDocOptions, nil)

	// Swagger docs API
	getRouter.Handle("/docs", swaggerMiddleware)

	// Download swagger.yaml file
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Initialize the plant-api server properties
	server := http.Server{
		Addr:         ":9090",
		Handler:      gorillaMux,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Initialize the go-routine function
	go func() {

		// ListenAndServe listens on the TCP network address specified in the server property
		listenAndServeError := server.ListenAndServe()

		if listenAndServeError != nil {
			logger.Fatal(listenAndServeError)
		}
		logger.Printf("Server running on port %s\n", server.Addr)
	}()

	// Make the channel with type os.Signal
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// Read the channel value
	sig := <-signalChannel

	logger.Println("Received os signal, graceful timeout", sig)

	//Canceling this context releases resources associated with it
	terminateContext, terminateContextError := context.WithTimeout(context.Background(), 30*time.Second)

	if terminateContextError != nil {
		logger.Fatal(terminateContextError)
	}

	// Shutdown gracefully shuts down the server without interrupting any active connections
	server.Shutdown(terminateContext)

}
