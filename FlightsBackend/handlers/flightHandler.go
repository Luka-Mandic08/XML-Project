package handlers

import (
	"Rest/model"
	"Rest/repositories"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FlightHandler struct {
	logger     *log.Logger
	repository *repositories.FlightRepository
}

// Injecting the logger makes this code much more testable.
func NewFlightHandler(l *log.Logger, r *repositories.FlightRepository) *FlightHandler {
	return &FlightHandler{l, r}
}

func (flightHandler *FlightHandler) InsertFlight(rw http.ResponseWriter, req *http.Request) {
	flight := req.Context().Value(KeyProduct{}).(*model.Flight)
	flightHandler.repository.Insert(flight)
	rw.WriteHeader(http.StatusCreated)
}

func (flightHandler *FlightHandler) GetFlightById(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	patient, err := flightHandler.repository.GetById(id)
	if err != nil {
		flightHandler.logger.Print("Database exception: ", err)
	}

	if patient == nil {
		http.Error(rw, "Patient with given id not foundd", http.StatusNotFound)
		flightHandler.logger.Printf("Patient with id: '%s' not found", id)
		return
	}

	err = patient.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		flightHandler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (flightHandler *FlightHandler) GetAllFlights(rw http.ResponseWriter, req *http.Request) {
	flgihts, err := flightHandler.repository.GetAll()
	if err != nil {
		flightHandler.logger.Print("Database exception: ", err)
	}

	if flgihts == nil {
		return
	}

	err = flgihts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		flightHandler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (flightHandler *FlightHandler) UpdateFlight(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	amount := req.Header.Get("amount")

	amount_int, _ := strconv.ParseInt(amount, 10, 64)

	flightHandler.repository.Update(id, amount_int)
	rw.WriteHeader(http.StatusOK)
}

func (flightHandler *FlightHandler) DeleteFlight(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	flightHandler.repository.Delete(id)
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
