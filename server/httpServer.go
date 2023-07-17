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
	// batchTraineeController           *controllers.BatchTraineeController
	// batchTraineeEvaluationController *controllers.BatchTraineeEvaluationController
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

	// BATCH TRAINEE
	batchTraineeRepository := repositories.NewBatchTraineeRepository(dbHandler)
	batchTraineeService := services.NewBatchTraineeService(batchTraineeRepository)
	batchTraineeController := controllers.NewBatchTraineeController(batchTraineeService)

	router.GET("/batch_trainee", batchTraineeController.GetListBatchTrainee)

	// BATCH TRAINEE EVALUATION
	batchTraineeEvaluationRepository := repositories.NewBatchTraineeEvaluationRepository(dbHandler)
	batchTraineeEvaluationService := services.NewBatchTraineeEvaluationService(batchTraineeEvaluationRepository)
	batchTraineeEvaluationController := controllers.NewBatchTraineeEvaluationController(batchTraineeEvaluationService)

	router.GET("/batch_trainee_evaluation", batchTraineeEvaluationController.GetListBatchTraineeEvaluation)

	// INSTRUCTOR PROGRAM
	instructorProgramRepository := repositories.NewInstructorProgramRepository(dbHandler)
	instructorProgramService := services.NewInstructorProgramService(instructorProgramRepository)
	instructorProgramController := controllers.NewInstructorProgramController(instructorProgramService)

	router.GET("/instructor_programs", instructorProgramController.GetListInstructorProgram)

	// PROGRAM APPLY
	programApplyRepository := repositories.NewProgramApplyRepository(dbHandler)
	programApplyService := services.NewProgramApplyService(programApplyRepository)
	programApplyController := controllers.NewProgramApplyController(programApplyService)

	router.GET("/program_apply", programApplyController.GetListProgramApply)

	// PROGRAM APPLY PROGRESS
	programApplyProgressRepository := repositories.NewProgramApplyProgressRepository(dbHandler)
	programApplyProgressService := services.NewProgramApplyProgressService(programApplyProgressRepository)
	programApplyProgressController := controllers.NewProgramApplyProgressController(programApplyProgressService)

	router.GET("/program_apply_progress", programApplyProgressController.GetListProgramApplyProgress)

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
