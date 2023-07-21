package bootcampServer

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers/bootcampController"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config          *viper.Viper
	router          *gin.Engine
	batchController *bootcampController.BatchController
	// batchTraineeController           *controllers.BatchTraineeController
	// batchTraineeEvaluationController *controllers.BatchTraineeEvaluationController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	// BATCH
	batchRepository := bootcampRepository.NewBatchRepository(dbHandler)
	batchService := bootcampService.NewBatchService(batchRepository)
	batchController := bootcampController.NewBatchController(batchService)

	router := gin.Default()

	router.GET("/batch", batchController.GetListBatch)
	router.GET("/batch/:id", batchController.GetBatch)
	router.POST("/batch", batchController.CreateBatch)
	router.PUT("/batch/:id", batchController.UpdateBatch)
	router.DELETE("/batch/:id", batchController.DeleteBatch)

	// GROUP
	bootcampBatchEvaluationRepository := bootcampRepository.NewBootcampBatchEvaluationRepository(dbHandler)
	bootcampBatchEvaluationService := bootcampService.NewBootcampBatchEvaluationService(bootcampBatchEvaluationRepository)
	bootcampBatchEvaluationController := bootcampController.NewBootcampBatchEvaluationController(bootcampBatchEvaluationService)

	router.GET("/group", bootcampBatchEvaluationController.GetListBootcampBatchEvaluation)

	// BATCH TRAINEE EVALUATION
	// batchTraineeEvaluationRepository := repositories.NewBatchTraineeEvaluationRepository(dbHandler)
	// batchTraineeEvaluationService := services.NewBatchTraineeEvaluationService(batchTraineeEvaluationRepository)
	// batchTraineeEvaluationController := controllers.NewBatchTraineeEvaluationController(batchTraineeEvaluationService)

	// router.GET("/batch_trainee_evaluation", batchTraineeEvaluationController.GetListBatchTraineeEvaluation)

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
