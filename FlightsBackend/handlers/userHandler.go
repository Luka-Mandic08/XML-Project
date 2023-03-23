package handlers

import (
	"Rest/model"
	"Rest/repositories"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type UsersHandler struct {
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *repositories.UserRepository
}

// Injecting the logger makes this code much more testable.
func NewUserHandler(l *log.Logger, r *repositories.UserRepository) *UsersHandler {
	return &UsersHandler{l, r}
}

// CREATE
func (p *UsersHandler) InsertUser(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyProduct{}).(*model.User)
	p.repo.Insert(user)
	rw.WriteHeader(http.StatusCreated)
}

// READ
func (p *UsersHandler) GetAllUsers(rw http.ResponseWriter, h *http.Request) {
	users, err := p.repo.GetAll()
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *UsersHandler) GetUserById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	patient, err := p.repo.GetById(id)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if patient == nil {
		http.Error(rw, "Patient with given id not found", http.StatusNotFound)
		p.logger.Printf("Patient with id: '%s' not found", id)
		return
	}

	err = patient.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *UsersHandler) GetUsersByName(rw http.ResponseWriter, h *http.Request) {
	name := h.URL.Query().Get("name")

	users, err := p.repo.GetByName(name)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

// UPDATE
func (p *UsersHandler) UpdateUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	user := h.Context().Value(KeyProduct{}).(*model.User)

	p.repo.Update(id, user)
	rw.WriteHeader(http.StatusOK)
}

func (p *UsersHandler) UpdateAddress(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	address := h.Context().Value(KeyProduct{}).(*model.Address)

	p.repo.UpdateAddress(id, address)
	rw.WriteHeader(http.StatusOK)
}

func (p *UsersHandler) UpdatePhone(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	index, err := strconv.Atoi(vars["index"])
	if err != nil {
		http.Error(rw, "Unable to decode index", http.StatusBadRequest)
		p.logger.Fatal(err)
		return
	}

	var phoneNumber string
	d := json.NewDecoder(h.Body)
	d.Decode(&phoneNumber)

	p.repo.ChangePhone(id, index, phoneNumber)
	rw.WriteHeader(http.StatusOK)
}

// DELETE
func (p *UsersHandler) DeleteUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	p.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

func (p *UsersHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			p.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (p *UsersHandler) MiddlewareAddressDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		address := &model.Address{}
		err := address.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			p.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, address)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (p *UsersHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
