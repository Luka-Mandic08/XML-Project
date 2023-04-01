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
	buyTicketDto := req.Context().Value(KeyProduct{}).(*model.BuyTicketDto)
	flightHandler.logger.Println("FlightId: " + buyTicketDto.FlightId)
	flightHandler.logger.Printf("Amopunt: %d\n", buyTicketDto.Amount)
	flightHandler.logger.Println("UserId:" + buyTicketDto.UserId)

	if buyTicketDto.Amount < 1 {
		http.Error(rw, "Can not buy Negative or Zero amount of cards.", http.StatusBadRequest)
		flightHandler.logger.Fatal("Negative or Zero amount of cards: ", buyTicketDto.Amount)
		return
	}

	/*_, errUser := flightHandler.userRepository.GetById(buyTicketDto.UserId)

	if errUser != nil {*/
	err := flightHandler.flightRepository.UpdateFlightRemainingTickets(buyTicketDto.FlightId, buyTicketDto.Amount)

	if err == nil {
		err := flightHandler.userRepository.AddFlight(buyTicketDto.UserId, buyTicketDto.FlightId, buyTicketDto.Amount) //treba promeniti id i id
		if err != nil {
			http.Error(rw, "Adding tickets unsuccessful!", http.StatusBadRequest)
			return
		}

		rw.WriteHeader(http.StatusOK)
		return
	} else {
		http.Error(rw, "Flight not found!", http.StatusBadRequest)
		flightHandler.logger.Fatal("Flight not found! ID: ", buyTicketDto.FlightId)
		return
	}
	/*}

	http.Error(rw, "User not found", http.StatusBadRequest)
	flightHandler.logger.Fatal("User not found! ID: ", buyTicketDto.UserId)*/
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
