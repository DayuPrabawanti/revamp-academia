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

	batchRepository := repositories.NewBatchRepository(dbHandler)

	batchService := services.NewBatchService(batchRepository)

	batchController := controllers.NewBatchController(batchService)

	router := gin.Default()

	//router endpoint
	router.GET("/api/bootcamp/batch/", batchController.GetListBatch)
	// router.GET("/api/bootcamp/batch/:id", batchController.GetBatch)
	router.GET("/api/bootcamp/batch/batchid", batchController.GetBatch) // Menggunakan query parameter id
	router.POST("/api/bootcamp/batch/create", batchController.CreateBatch)

	router.PUT("/api/bootcamp/batch/update/:id", batchController.UpdateBatch)
	router.DELETE("/api/bootcamp/batch/delete/:id", batchController.DeleteBatch)

	// router.GET("/api/bootcamp/batch/search?/batch=#3 & status=Open", batchController.SearchBatch)
	router.GET("/api/bootcamp/batch/search", batchController.SearchBatch)

	// router.GET("/api/bootcamp/batch/paging? page=1 & pageSize=10", batchController.PagingBatch)
	// router.GET("/api/bootcamp/batch/paging", batchController.PagingBatch) // Paging

	return HttpServer{
		config:          config,
		router:          router,
		batchController: batchController,
	}
}

// running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP Server : %v", err)
	}
}
