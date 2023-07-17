package userssv

import (
	"database/sql"
	"log"

	"codeid.revamptwo/controllers/usersct"
	"codeid.revamptwo/repositories/users"
	"codeid.revamptwo/services/usersc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	userController *usersct.UserController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer{

	userRepository := users.NewUserRepository(dbHandler)
	userService := usersc.NewUserService(userRepository)

	userController := usersct.NewUserController(userService)

	router := gin.Default()

	// Router endpoint (url) http Users
	router.GET("/users", userController.GetListUser)
	router.GET("/users/:id", userController.GetUser)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)


	return HttpServer{
		config: config,
		router: router,
		userController: userController,
	}
}

// Running gin httpserver
func (hs HttpServer)Start(){
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err!= nil {
		log.Fatalf("Error While Starting HTTP Server : %v", err)
	}
}