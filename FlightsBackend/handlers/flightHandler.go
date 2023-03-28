package handlers

import (
	"Rest/model"
	"Rest/repositories"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FlightHandler struct {
	logger           *log.Logger
	flightRepository *repositories.FlightRepository
}

// Injecting the logger makes this code much more testable.
func NewFlightHandler(l *log.Logger, r *repositories.FlightRepository) *FlightHandler {
	return &FlightHandler{l, r}
}

func (flightHandler *FlightHandler) InsertFlight(rw http.ResponseWriter, req *http.Request) {
	flight := req.Context().Value(KeyProduct{}).(*model.Flight)
	flightHandler.flightRepository.Insert(flight)
	rw.WriteHeader(http.StatusCreated)
}

func (flightHandler *FlightHandler) GetFlightById(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	flight, err := flightHandler.flightRepository.GetById(id)
	if err != nil {
		flightHandler.logger.Print("Database exception: ", err)
	}

	if flight == nil {
		http.Error(rw, "Flight with given id not found", http.StatusNotFound)
		flightHandler.logger.Printf("Flight with id: '%s' not found", id)
		return
	}

	err = flight.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		flightHandler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (flightHandler *FlightHandler) GetAllFlights(rw http.ResponseWriter, req *http.Request) {
	flights, err := flightHandler.flightRepository.GetAll()
	if err != nil {
		flightHandler.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		flightHandler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (flightHandler *FlightHandler) UpdateFlightRemainingTickets(rw http.ResponseWriter, req *http.Request) {

	buyTicketDto := req.Context().Value(KeyProduct{}).(*model.BuyTicketDto)
	flightHandler.logger.Println(buyTicketDto.Amount)
	flightHandler.logger.Println(buyTicketDto.FlightId)

	if buyTicketDto.Amount < 1 {
		http.Error(rw, "Negative or Zero amount of cards. Can not buy.", http.StatusBadRequest)
		flightHandler.logger.Fatal("Negative or Zero amount of cards: ", buyTicketDto.Amount)
		return
	}

	flightHandler.flightRepository.UpdateFlightRemainingTickets(buyTicketDto.FlightId, buyTicketDto.Amount)
	rw.WriteHeader(http.StatusOK)
}

func (flightHandler *FlightHandler) DeleteFlight(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	flightHandler.flightRepository.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

func (f *FlightHandler) MiddlewareFlightDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		flight := &model.Flight{}
		err := flight.FromJSON(req.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			f.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(req.Context(), KeyProduct{}, flight)
		req = req.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}

func (f *FlightHandler) MiddlewareBuyTicketsDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		f.logger.Println(req.Body)
		f.logger.Println(req.FormValue("flightId"))
		f.logger.Println(req.FormValue("amount"))

		buyTicketDto := &model.BuyTicketDto{}
		err := buyTicketDto.FromJSON(req.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			f.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(req.Context(), KeyProduct{}, buyTicketDto)
		req = req.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
