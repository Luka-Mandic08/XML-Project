package handlers

import (
	"Rest/data"
	"Rest/flights"
	"context"
	"log"
	"net/http"
)

type FlightHandler struct {
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *flights.FlightRepository
}

// Injecting the logger makes this code much more testable.
func NewFlightHandler(l *log.Logger, r *flights.FlightRepository) *FlightHandler {
	return &FlightHandler{l, r}
}

func (f *FlightHandler) PostFlight(rw http.ResponseWriter, h *http.Request) {
	flight := h.Context().Value(KeyProduct{}).(*data.Flight)
	f.repo.Insert(flight)
	rw.WriteHeader(http.StatusCreated)
}

func (f *FlightHandler) MiddlewareFlightDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		flight := &data.Flight{}
		err := flight.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			f.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, flight)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
