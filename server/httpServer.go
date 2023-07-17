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
	jobcategoryController *controllers.JobCategoryController
    jobclientController *controllers.JobPostController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
 
	jobCategoryRepository := repositories.NewJobCategoryRepository(dbHandler)
	jobCategoryService := services.NewJobCategoryService(jobCategoryRepository)
	jobCategoryController := controllers.NewJobCategoryController(jobCategoryService)

    jobClientRepository := repositories.NewJobClientRepository(dbHandler)
	jobClientService := services.NewJobClientService(jobClientRepository)
	jobClientController := controllers.NewJobClientController(jobClientService)

    jobPostRepository := repositories.NewJobPostRepository(dbHandler)
	jobPostService := services.NewJobPostService(jobPostRepository)
	jobPostController := controllers.NewJobPostController(jobPostService)


    router := gin.Default()

    // jobCategory

	router.GET("/jobCategory/:id", jobCategoryController.GetJobCategoryHttp)
        router.GET("/jobCategory", jobCategoryController.ListJobCategoryHttp)
            router.POST("/createJobCategory", jobCategoryController.CreateJobCategoryHttp)
                router.PUT("/jobCategory/:id", jobCategoryController.UpdateJobCategoryHttp)
                    router.DELETE("/category/:id", jobCategoryController.DeleteJobCategoryHttp)

    // jobClient

    router.GET("/jobClient/:id", jobClientController.GetJobClientHttp)
        router.GET("/jobClientList/", jobClientController.ListJobClientHttp)
            router.POST("/jobClientCreate", jobClientController.CreateJobClientHttp)
                router.PUT("/jobClientUpdate/:id", jobClientController.UpdateJobClientHttp)
                    router.DELETE("/jobClientDelete/:id", jobClientController.DeleteJobClientHttp)


    // jobPost

    router.GET("/jobPost/:id", jobPostController.GetJobPostHttp)
        router.GET("/jobPostList/", jobPostController.ListJobPostHttp)




    return HttpServer{
        config:             config,
        router:             router,
        jobcategoryController: jobCategoryController,
        // clientController: jobClientController,
    }
}

// Running gin HttpServer
func (hs HttpServer) Start() {
    err := hs.router.Run(hs.config.GetString("http.server_address"))

    if err != nil {
        log.Fatalf("Error while starting HTTP Server: %v", err)
    }
}
