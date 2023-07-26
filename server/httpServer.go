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
	config             *viper.Viper
	router             *gin.Engine
	ControllersManager controllers.ControllersManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
 
    repositoriesManager := repositories.NewRepositoriesManager(dbHandler)

    servicesManager := services.NewServiceManager(repositoriesManager)

    controllerManager := controllers.NewControllersManager(servicesManager)



    router := gin.Default()

    InitRouter(router, controllerManager)




     
   return HttpServer{
		config:            config,
		router:            router,
		ControllersManager: *controllerManager,
	}
}

func (hs HttpServer) Start() {
    err := hs.router.Run(hs.config.GetString("http.server_address"))

    if err != nil {
        log.Fatalf("Error while starting HTTP Server: %v", err)
    }
}



    // Jobhire

	


    // Master



    // groupRepository := repositories.NewJobPostRepository(dbHandler)

	// groupService := services.NewJobPostService(groupRepository)

	// groupController := controllers.NewJobPostController(groupService)

	


    // jobCategory

	

    // JOB POSTING

       


// Running gin HttpServer
