package handlers

import (
	"Rest/model"
	"Rest/repositories"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type KeyProduct struct{}

type UserHandler struct {
	logger *log.Logger
	repo   *repositories.UserRepository
}

// Injecting the logger makes this code much more testable.
func NewUserHandler(l *log.Logger, r *repositories.UserRepository) *UserHandler {
	return &UserHandler{l, r}
}

// CREATE
func (userHandler *UserHandler) InsertUser(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyProduct{}).(*model.User)
	userHandler.repo.Insert(user)
	rw.WriteHeader(http.StatusCreated)
}

// READ
func (p *UserHandler) GetAllUsers(rw http.ResponseWriter, h *http.Request) {
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

func (p *UserHandler) GetUserById(rw http.ResponseWriter, h *http.Request) {
	//vars := mux.Vars(h)
	//id := vars["id"]
	id := h.URL.Query().Get("id")

	patient, err := p.repo.GetById(id)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if patient == nil {
		http.Error(rw, "User with given id not found", http.StatusNotFound)
		p.logger.Printf("User with id: '%s' not found", id)
		return
	}

	err = patient.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *UserHandler) GetUsersByName(rw http.ResponseWriter, h *http.Request) {
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
func (p *UserHandler) UpdateUser(rw http.ResponseWriter, h *http.Request) {
	id := h.URL.Query().Get("id")
	user := h.Context().Value(KeyProduct{}).(*model.User)

	p.repo.Update(id, user)
	rw.WriteHeader(http.StatusOK)
}

func (p *UserHandler) UpdateAddress(rw http.ResponseWriter, h *http.Request) {
	id := h.URL.Query().Get("id")
	address := h.Context().Value(KeyProduct{}).(*model.UserAddress)

	p.repo.UpdateAddress(id, address)
	rw.WriteHeader(http.StatusOK)
}

func (p *UserHandler) UpdateCredentials(rw http.ResponseWriter, h *http.Request) {
	id := h.URL.Query().Get("id")
	credentials := h.Context().Value(KeyProduct{}).(*model.UserCredentials)

	p.repo.UpdateCredentials(id, credentials)
	rw.WriteHeader(http.StatusOK)
}

/*
func (p *UserHandler) UpdatePhone(rw http.ResponseWriter, h *http.Request) {
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
*/

// DELETE
func (p *UserHandler) DeleteUser(rw http.ResponseWriter, h *http.Request) {
	id := h.URL.Query().Get("id")

	p.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

// LOGIN/LOGOUT
func (p *UserHandler) LoginUser(rw http.ResponseWriter, h *http.Request) {
	credentials := h.Context().Value(KeyProduct{}).(*model.UserCredentials)

	user, err := p.repo.Login(credentials)
	if err != nil {
		p.logger.Println(err)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	if user != nil {
		session := &model.Session{ID: user.ID, Role: user.Role}

		err = session.ToJSON(rw)
		if err != nil {
			http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
			p.logger.Fatal("Unable to convert to json :", err)

			return
		}
		rw.WriteHeader(http.StatusAccepted)
	} else {
		rw.WriteHeader(http.StatusBadRequest)
	}

	return
}

func (p *UserHandler) LogoutUser(rw http.ResponseWriter, h *http.Request) {
	e := json.NewEncoder(rw)
	err := e.Encode("")
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

// MANAGE FLIGHTS

// MIDDLEWARE

func (p *UserHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
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

func (p *UserHandler) MiddlewareAddressDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		address := &model.UserAddress{}
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

func (p *UserHandler) MiddlewareCredentialsDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		credentials := &model.UserCredentials{}
		err := credentials.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			p.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, credentials)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (p *UserHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
