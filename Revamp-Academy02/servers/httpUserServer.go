package servers

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers/usersCt"
	"codeid.revampacademy/repositories/users"
	"codeid.revampacademy/services/usersService"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	userController *usersCt.UserController
	userEmailController *usersCt.UserEmailController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer{


					// >> Users <<
	userRepository := users.NewUserRepository(dbHandler)
	userService := usersService.NewUserService(userRepository)
	userController := usersCt.NewUserController(userService)

					// >> Users Email <<
	userEmailRepository := users.NewUserEmailRepository(dbHandler)
	userEmailService := usersService.NewUserEmailService(userEmailRepository)
	userEmailController := usersCt.NewUserEmailController(userEmailService)

	router := gin.Default()

	// Router endpoint (url) http Users
	router.GET("/users", userController.GetListUser)
	router.GET("/users/:id", userController.GetUser)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	// Router endpoint (url) http Users Email
	router.GET("/usersemail", userEmailController.GetListUsersEmail)
	router.GET("/usersemail/:id", userEmailController.GetEmail)
	router.POST("/usersemail", userEmailController.CreateEmail)
	// router.PUT("/users/:id", userController.UpdateUser)
	// router.DELETE("/users/:id", userController.DeleteUser)


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