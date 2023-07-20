package server

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config         *viper.Viper
	router         *gin.Engine
	userController *controllers.UserController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	// >> Users <<
	userRepository := repositories.NewUserRepository(dbHandler)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	router := gin.Default()

	// Router endpoint (url) http Users
	router.GET("/users", userController.GetListUser)

	return HttpServer{
		config:         config,
		router:         router,
		userController: userController,
	}
}

// Running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error While Starting HTTP Server : %v", err)
	}
}
