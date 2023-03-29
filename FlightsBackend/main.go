package main

import (
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
	userLogger := log.New(os.Stdout, "[user-store] ", log.LstdFlags)
	flightLogger := log.New(os.Stdout, "[flight-store] ", log.LstdFlags)

	// NoSQL: Initialize Repositories

	userstore, err := repositories.NewUserRepository(timeoutContext, userLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer userstore.Disconnect(timeoutContext)

	userstore.Ping()

	flightstore, err := repositories.NewFlightRepository(timeoutContext, flightLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer flightstore.Disconnect(timeoutContext)

	flightstore.Ping()

	// Initialize the handler and inject said logger
	userHandler := handlers.NewUserHandler(logger, userstore)
	flightHandler := handlers.NewFlightHandler(logger, flightstore, userstore)

	// Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	router.Use(userHandler.MiddlewareContentTypeSet)

	// Users CREATE
	addUserRouter := router.Methods(http.MethodPost).Subrouter()
	addUserRouter.HandleFunc("/user/add", userHandler.InsertUser)
	addUserRouter.Use(userHandler.MiddlewareUserDeserialization)

	// Users READ
	getUsersRouter := router.Methods(http.MethodGet).Subrouter()
	getUsersRouter.HandleFunc("/users", userHandler.GetAllUsers)

	getUserByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getUserByIdRouter.HandleFunc("/user/id", userHandler.GetUserById)

	// http://localhost:8082/user/641d92e04d666e40fd539f77
	// getUserByIdRouter := router.Methods(http.MethodGet).Subrouter()
	// getUserByIdRouter.HandleFunc("/user/{id}", userHandler.GetUserById)

	getUserByNameRouter := router.Methods(http.MethodGet).Subrouter()
	getUserByNameRouter.HandleFunc("/user/name", userHandler.GetUsersByName)

	// Users UPDATE
	updateUserRouter := router.Methods(http.MethodPatch).Subrouter()
	updateUserRouter.HandleFunc("/user/update", userHandler.UpdateUser)
	updateUserRouter.Use(userHandler.MiddlewareUserDeserialization)

	updateUserAddressRouter := router.Methods(http.MethodPatch).Subrouter()
	updateUserAddressRouter.HandleFunc("/user/updateAddress", userHandler.UpdateAddress)
	updateUserAddressRouter.Use(userHandler.MiddlewareAddressDeserialization)

	updateUserCredentialsRouter := router.Methods(http.MethodPatch).Subrouter()
	updateUserCredentialsRouter.HandleFunc("/user/updateCredentials", userHandler.UpdateCredentials)
	updateUserCredentialsRouter.Use(userHandler.MiddlewareCredentialsDeserialization)

	// Users DELETE
	deleteUserRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteUserRouter.HandleFunc("/user/delete", userHandler.DeleteUser)

	// Users LOGIN/LOGOUT
	checkUserCredentialsRouter := router.Methods(http.MethodGet).Subrouter()
	checkUserCredentialsRouter.HandleFunc("/login", userHandler.LoginUser)
	checkUserCredentialsRouter.Use(userHandler.MiddlewareCredentialsDeserialization)

	logoutUserRouter := router.Methods(http.MethodGet).Subrouter()
	logoutUserRouter.HandleFunc("/logout", userHandler.LogoutUser)

	//Flights CRUD
	getAllFlightsRouter := router.Methods(http.MethodGet).Subrouter()
	getAllFlightsRouter.HandleFunc("/flight/all", flightHandler.GetAllFlights)

	getSearchedFlightsRouter := router.Methods(http.MethodGet).Subrouter()
	getSearchedFlightsRouter.HandleFunc("/flight/search", flightHandler.GetSearchedFlights)
	getSearchedFlightsRouter.Use(flightHandler.MiddlewareFlightSearchDeserialization)

	getFlightByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getFlightByIdRouter.HandleFunc("/flight/{id}", flightHandler.GetFlightById)

	postFlightRouter := router.Methods(http.MethodPost).Subrouter()
	postFlightRouter.HandleFunc("/flight/add", flightHandler.InsertFlight)
	postFlightRouter.Use(flightHandler.MiddlewareFlightDeserialization)

	updateFlightRouter := router.Methods(http.MethodPut).Subrouter()
	updateFlightRouter.HandleFunc("/flight/update/{id}", flightHandler.UpdateFlightRemainingTickets)

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
