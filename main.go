package main

import (
	"Rest/data"
	"Rest/handlers"
	"Rest/repositories"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//Reading from environment, if not set we will default it to 8080.
	//This allows flexibility in different environments (for eg. when running multiple docker api's and want to override the default port)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[patient-store] ", log.LstdFlags)
	flightLogger := log.New(os.Stdout, "[flight-store] ", log.LstdFlags)

	// NoSQL: Initialize Patient Repository store
	store, err := data.New(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	store.Ping()

	flightstore, err := repositories.New(timeoutContext, flightLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer flightstore.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	flightstore.Ping()

	//Initialize the handler and inject said logger
	patientsHandler := handlers.NewPatientsHandler(logger, store)
	flightHandler := handlers.NewFlightHandler(logger, flightstore)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	router.Use(patientsHandler.MiddlewareContentTypeSet)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", patientsHandler.GetAllPatients)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", patientsHandler.PostPatient)
	postRouter.Use(patientsHandler.MiddlewarePatientDeserialization)

	getByNameRouter := router.Methods(http.MethodGet).Subrouter()
	getByNameRouter.HandleFunc("/filter", patientsHandler.GetPatientsByName)

	receiptRouter := router.Methods(http.MethodGet).Subrouter()
	receiptRouter.HandleFunc("/receipt/{id}", patientsHandler.Receipt)

	reportRouter := router.Methods(http.MethodGet).Subrouter()
	reportRouter.HandleFunc("/report", patientsHandler.Report)

	getByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getByIdRouter.HandleFunc("/{id}", patientsHandler.GetPatientById)

	patchRouter := router.Methods(http.MethodPatch).Subrouter()
	patchRouter.HandleFunc("/{id}", patientsHandler.PatchPatient)
	patchRouter.Use(patientsHandler.MiddlewarePatientDeserialization)

	changePhoneRouter := router.Methods(http.MethodPatch).Subrouter()
	changePhoneRouter.HandleFunc("/phone/{id}/{index}", patientsHandler.ChangePhone)

	pushPhoneRouter := router.Methods(http.MethodPatch).Subrouter()
	pushPhoneRouter.HandleFunc("/phone/{id}", patientsHandler.AddPhoneNumber)

	addAnamnesisRouter := router.Methods(http.MethodPatch).Subrouter()
	addAnamnesisRouter.HandleFunc("/anamnesis/{id}", patientsHandler.AddAnamnesis)

	addTherapyRouter := router.Methods(http.MethodPatch).Subrouter()
	addTherapyRouter.HandleFunc("/therapy/{id}", patientsHandler.AddTherapy)

	changeAddressRouter := router.Methods(http.MethodPatch).Subrouter()
	changeAddressRouter.HandleFunc("/address/{id}", patientsHandler.ChangeAddress)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id}", patientsHandler.DeletePatient)

	//Flights CRUD
	getFlightByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getFlightByIdRouter.HandleFunc("/flight/{id}", flightHandler.GetFlightById)

	getAllFlightsRouter := router.Methods(http.MethodGet).Subrouter()
	getAllFlightsRouter.HandleFunc("/flights", flightHandler.GetAllFlights)

	postFlightRouter := router.Methods(http.MethodPost).Subrouter()
	postFlightRouter.HandleFunc("/addflight", flightHandler.InsertFlight)
	postFlightRouter.Use(flightHandler.MiddlewareFlightDeserialization)

	updateFlightRouter := router.Methods(http.MethodPatch).Subrouter()
	updateFlightRouter.HandleFunc("/update/flight/{id}", flightHandler.UpdateFlight)
	updateFlightRouter.Use(flightHandler.MiddlewareFlightDeserialization)

	deleteFlightRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteFlightRouter.HandleFunc("/flight/{id}", flightHandler.DeleteFlight)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
