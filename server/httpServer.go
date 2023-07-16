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
	config          *viper.Viper
	router          *gin.Engine
	batchController *controllers.BatchController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	// BATCH
	batchRepository := repositories.NewBatchRepository(dbHandler)
	batchService := services.NewBatchService(batchRepository)
	batchController := controllers.NewBatchController(batchService)

	router := gin.Default()

	router.GET("/batch", batchController.GetListBatch)
	router.GET("/batch/id", batchController.GetBatch)
	router.POST("/batch", batchController.CreateBatch)
	router.PUT("/batch/:id", batchController.UpdateBatch)
	router.DELETE("/batch/:id", batchController.DeleteBatch)

	// BATCH TRAINEE EVALUATION

	return HttpServer{
		config:          config,
		router:          router,
		batchController: batchController,
	}
}

// Running gin HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
