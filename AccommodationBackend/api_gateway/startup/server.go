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
	authGroup.Use(services.ValidateToken()) //Login i register do not use ValidateToken()
	authGroup.PUT("/update", authHandler.Update)
	authGroup.DELETE("/delete/:userId", authHandler.Delete)
	authGroup.GET("/get/:userId", authHandler.GetByUserId)

	userGroup := router.Group("/users")
	userGroup.Use(services.ValidateToken())
	userGroup.GET("/:id", userHandler.Get) //services.AuthorizeRole("Host") TODO: add for Guest
	userGroup.PUT("/update", userHandler.Update)

	accommodationGroup := router.Group("/accommodation")
	accommodationGroup.GET("/all", accommodationHandler.GetAll)
	accommodationGroup.GET("/:id", accommodationHandler.GetById)
	accommodationGroup.POST("/search", accommodationHandler.Search)
	accommodationGroup.Use(services.ValidateToken())
	accommodationGroup.POST("/create", services.AuthorizeRole("Host"), accommodationHandler.Create)
	accommodationGroup.POST("/updateAvailability", services.AuthorizeRole("Host"), accommodationHandler.UpdateAvailability)
	accommodationGroup.POST("/checkAvailability", accommodationHandler.CheckAvailability)
	accommodationGroup.GET("/all/host/:hostId", accommodationHandler.GetAllByHostId)
	accommodationGroup.PUT("/availability", accommodationHandler.GetAvailabilities)

	reservationGroup := router.Group("/reservation")
	reservationGroup.Use(services.ValidateToken())
	reservationGroup.GET("/getAllByUserId/:id", reservationHandler.GetAllByUserId)
	reservationGroup.POST("/request", services.AuthorizeRole("Guest"), reservationHandler.Request)

	return router
}