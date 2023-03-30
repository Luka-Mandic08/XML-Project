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
	userRepository   *repositories.UserRepository
}

// Injecting the logger makes this code much more testable.
func NewFlightHandler(l *log.Logger, r *repositories.FlightRepository, rUser *repositories.UserRepository) *FlightHandler {
	return &FlightHandler{l, r, rUser}
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

func (flightHandler *FlightHandler) GetSearchedFlights(rw http.ResponseWriter, req *http.Request) {
	flightsearchDTO := req.Context().Value(KeyProduct{}).(*model.FlightSearchDTO)
	flights, err := flightHandler.flightRepository.GetSearched(flightsearchDTO)
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
	/*vars := mux.Vars(req)
	id := vars["id"]

	amount := req.Header.Get("amount")

	amount_int, _ := strconv.ParseInt(amount, 10, 64)

	if amount_int < 0 {
		http.Error(rw, "Negative amount of cards.", http.StatusBadRequest)
		flightHandler.logger.Fatal("Negative amount of cards: ", amount_int)
		return
	}*/

	//flightHandler.flightRepository.UpdateFlightRemainingTickets(id, amount_int)
	userID := "6424c72be0d1136f9b01a438"
	flightID := "6424c733e0d1136f9b01a439"
	ticketCount := int64(4)

	_, err := flightHandler.flightRepository.GetById(flightID)

	if err == nil {
		err := flightHandler.userRepository.AddFlight(userID, flightID, ticketCount) //treba promeniti id i id
		if err != nil {
			http.Error(rw, "Adding tickets unsuccessful!", http.StatusBadRequest)
			return
		}

		rw.WriteHeader(http.StatusOK)
	} else {
		http.Error(rw, "Flight not found!", http.StatusBadRequest)
		flightHandler.logger.Fatal("Flight not found! ID: ", flightID)
		return
	}

	return
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

func (f *FlightHandler) MiddlewareFlightSearchDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		flight := &model.FlightSearchDTO{}
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
