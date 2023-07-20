package servers

import (
	"database/sql"
	"log"

	usersCt "codeid.revampacademy/controllers/usersController"
	usersRepo "codeid.revampacademy/repositories/usersRepository"
	usersServ "codeid.revampacademy/services/userServices"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
	// userController      *usersCt.UserEmailController
	userEmailController *usersCt.UserEmailController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	// >> Users Email <<
	userEmailRepository := usersRepo.NewUserEmailRepository(dbHandler)
	userEmailService := usersServ.NewUserEmailService(userEmailRepository)
	userEmailController := usersCt.NewUserEmailController(userEmailService)

	router := gin.Default()

	// Router endpoint (url) http Users Email
	router.GET("/usersemail", userEmailController.GetListUsersEmail)
	router.GET("/usersemail/:id", userEmailController.GetUserEmail)
	router.POST("/usersemail", userEmailController.CreateUserEmail)
	router.PUT("/usersemail/:id", userEmailController.UpdateEmail)
	router.DELETE("/usersemail/:id", userEmailController.DeleteEmail)

	return HttpServer{
		config:              config,
		router:              router,
		userEmailController: userEmailController,
	}
}

// Running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error While Starting HTTP Server : %v", err)
	}
}
