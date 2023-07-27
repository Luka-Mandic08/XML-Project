package startup

import (
	handler "api_gateway/infrastructure/api"
	grpc_client "api_gateway/infrastructure/services"
	cfg "api_gateway/startup/config"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"time"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *http.Server {
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

func CreateRoutersAndSetRoutes(config *cfg.Config) *gin.Engine {
	//MICROSERVICES
	userServiceAddress := fmt.Sprintf("%s:%s", config.UserHost, config.UserPort)
	userClient := grpc_client.NewUserClient(userServiceAddress)
	userHandler := handler.NewUserHandler(userClient)

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	})

	router := gin.Default()
	router.Use(corsMiddleware)
	userGroup := router.Group("/users")
	userGroup.GET("/:id", userHandler.Get)
	return router
}
