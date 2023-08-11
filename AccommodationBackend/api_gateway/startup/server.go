package startup

import (
	handler "api_gateway/infrastructure/api"
	services "api_gateway/infrastructure/services"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"time"
)

type Server struct {
	config *Config
	mux    *runtime.ServeMux
}

func NewServer(config *Config) *http.Server {
	publicAddress := fmt.Sprintf("%s:%s", config.Host, config.Port)
	router := CreateRoutersAndSetRoutes(config)
	publicServer := &http.Server{
		Handler:           router,
		Addr:              publicAddress,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 100 * time.Millisecond,
		MaxHeaderBytes:    2048,
	}
	return publicServer
}

func CreateRoutersAndSetRoutes(config *Config) *gin.Engine {
	//MICROSERVICES
	userServiceAddress := fmt.Sprintf("%s:%s", config.UserHost, config.UserPort)
	userClient := services.NewUserClient(userServiceAddress)
	userHandler := handler.NewUserHandler(userClient)

	authServiceAddress := fmt.Sprintf("%s:%s", config.AuthHost, config.AuthPort)
	authClient := services.NewAuthClient(authServiceAddress)
	authHandler := handler.NewAuthHandler(authClient, userClient)

	accommodationServiceAddress := fmt.Sprintf("%s:%s", config.AccommodationHost, config.AccommodationPort)
	accommodationClient := services.NewAccommodationClient(accommodationServiceAddress)
	accommodationHandler := handler.NewAccommodationHandler(accommodationClient)

	reservationServiceAddress := fmt.Sprintf("%s:%s", config.ReservationHost, config.ReservationPort)
	reservationClient := services.NewReservationClient(reservationServiceAddress)
	reservationHandler := handler.NewReservationHandler(reservationClient)

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	})

	router := gin.Default()
	router.Use(corsMiddleware)

	authGroup := router.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)
	authGroup.Use(services.ValidateToken()) //Login i register su navedeni pre ove linije, pa se za njih ne koristi validateToken
	authGroup.PUT("/update", authHandler.Update)
	authGroup.DELETE("/delete/:id", authHandler.Delete)

	userGroup := router.Group("/users")
	userGroup.Use(services.ValidateToken())
	userGroup.GET("/:id", services.AuthorizeRole("Host"), userHandler.Get)
	userGroup.PUT("/update", userHandler.Update)

	accommodationGroup := router.Group("/accommodation")
	accommodationGroup.Use(services.ValidateToken())
	accommodationGroup.POST("/create", services.AuthorizeRole("Host"), accommodationHandler.Create)
	/*{
	    "name":"Vila detelinara",
	    "address":{
	        "street":"Moja ulica",
	        "city":"Ns",
	        "country":"Srbija"
	    },
	    "amenities":["Klima","Bazen"],
	    "images":["a","b"],
	    "minGuests":2,
	    "maxGuests":5,
	    "hostId":"64d4fdddddf5b55946ce909f",
	    "priceIsPerGuest":true,
	    "hasAutomaticReservations":false
	}*/
	accommodationGroup.POST("/updateAvailability", services.AuthorizeRole("Host"), accommodationHandler.UpdateAvailability)
	accommodationGroup.POST("/checkAvailability", accommodationHandler.CheckAvailability)
	accommodationGroup.POST("/search", accommodationHandler.Search)

	reservationGroup := router.Group("/reservation")
	reservationGroup.Use(services.ValidateToken())
	reservationGroup.GET("/test", reservationHandler.Test)
	reservationGroup.GET("/get/:id", services.AuthorizeRole("Host"), reservationHandler.Get)
	reservationGroup.POST("/create", services.AuthorizeRole("Host"), reservationHandler.Create)

	return router
}
