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
func (handler *UserHandler) InsertUser(rw http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(KeyProduct{}).(*model.User)
	handler.repo.Insert(user)
	rw.WriteHeader(http.StatusCreated)
}

// READ
func (handler *UserHandler) GetAllUsers(rw http.ResponseWriter, request *http.Request) {
	users, err := handler.repo.GetAll()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *UserHandler) GetUserById(rw http.ResponseWriter, request *http.Request) {
	//vars := mux.Vars(request)
	//id := vars["id"]
	id := request.URL.Query().Get("id")

	patient, err := handler.repo.GetById(id)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if patient == nil {
		http.Error(rw, "User with given id not found", http.StatusNotFound)
		handler.logger.Printf("User with id: '%s' not found", id)
		return
	}

	err = patient.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *UserHandler) GetUsersByName(rw http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	users, err := handler.repo.GetByName(name)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

// UPDATE
func (handler *UserHandler) UpdateUser(rw http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	user := request.Context().Value(KeyProduct{}).(*model.User)

	handler.repo.Update(id, user)
	rw.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateAddress(rw http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	address := request.Context().Value(KeyProduct{}).(*model.UserAddress)

	handler.repo.UpdateAddress(id, address)
	rw.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateCredentials(rw http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	credentials := request.Context().Value(KeyProduct{}).(*model.UserCredentials)

	handler.repo.UpdateCredentials(id, credentials)
	rw.WriteHeader(http.StatusOK)
}

/*
func (handler *UserHandler) UpdatePhone(rw http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	index, err := strconv.Atoi(vars["index"])
	if err != nil {
		http.Error(rw, "Unable to decode index", http.StatusBadRequest)
		handler.logger.Fatal(err)
		return
	}

	var phoneNumber string
	d := json.NewDecoder(request.Body)
	d.Decode(&phoneNumber)

	handler.repo.ChangePhone(id, index, phoneNumber)
	rw.WriteHeader(http.StatusOK)
}
*/

// DELETE
func (handler *UserHandler) DeleteUser(rw http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	handler.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

// LOGIN/LOGOUT
func (handler *UserHandler) LoginUser(rw http.ResponseWriter, request *http.Request) {
	credentials := request.Context().Value(KeyProduct{}).(*model.UserCredentials)

	user, err := handler.repo.Login(credentials)
	if err != nil {
		handler.logger.Println(err)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	if user != nil {
		session := &model.Session{ID: user.ID, Role: user.Role}

		err = session.ToJSON(rw)
		if err != nil {
			http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
			handler.logger.Fatal("Unable to convert to json :", err)

			return
		}
		rw.WriteHeader(http.StatusAccepted)
	} else {
		rw.WriteHeader(http.StatusBadRequest)
	}

	return
}

func (handler *UserHandler) LogoutUser(rw http.ResponseWriter, request *http.Request) {
	e := json.NewEncoder(rw)
	err := e.Encode("")
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) LinkUserToBookingApp(rw http.ResponseWriter, request *http.Request) {
	dto := request.Context().Value(KeyProduct{}).(*model.LinkUserDTO)

	user, err := handler.repo.Login(&model.UserCredentials{Username: dto.Username, Password: dto.Password})
	if err != nil {
		handler.logger.Println(err)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	msg, err := handler.repo.LinkUserToBookingApp(user.ID, &dto.ApiKey)
	if err != nil {
		handler.logger.Println(err)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	handler.logger.Println(msg)
	rw.WriteHeader(http.StatusOK)
}

// MANAGE FLIGHTS

// MIDDLEWARE

func (handler *UserHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		user := &model.User{}
		err := user.FromJSON(request.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, user)
		request = request.WithContext(ctx)

		next.ServeHTTP(rw, request)
	})
}

func (handler *UserHandler) MiddlewareAddressDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		address := &model.UserAddress{}
		err := address.FromJSON(request.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, address)
		request = request.WithContext(ctx)

		next.ServeHTTP(rw, request)
	})
}

func (handler *UserHandler) MiddlewareCredentialsDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		credentials := &model.UserCredentials{}
		err := credentials.FromJSON(request.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, credentials)
		request = request.WithContext(ctx)

		next.ServeHTTP(rw, request)
	})
}

func (handler *UserHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		handler.logger.Println("Method [", request.Method, "] - Hit path :", request.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, request)
	})
}

func (handler *UserHandler) MiddlewareLinkUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		dto := &model.LinkUserDTO{}
		err := dto.FromJSON(request.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, dto)
		request = request.WithContext(ctx)

		next.ServeHTTP(rw, request)
	})
}
