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
	config                  *viper.Viper
	router                  *gin.Engine
	programEntityController *controllers.ProgramEntityController
	// progEntityDescController *controllers.ProgEntityDescController
	// progReviewsController    *controllers.ProgReviewsController
	// sections                 *controllers.SectionController
	// sectionDetail            *controllers.SectionDetailController
	// sectionDetailMaterial    *controllers.SectionDetailMaterialController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	// PROGRAM ENTITY

	programEntityRepository := repositories.NewProgramEntityRepository(dbHandler)

	programEntityService := services.NewProgramEntityService(programEntityRepository)

	programEntityController := controllers.NewProgramEntityController(programEntityService)

	router := gin.Default()

	router.GET("/programEntity", programEntityController.GetListProgramEntity)

	router.GET("/programEntity/:id", programEntityController.GetProgramEntity)

	router.POST("/programEntity", programEntityController.CreateProgramEntity)

	router.PUT("/programEntity/:id", programEntityController.UpdateProgramEntity)

	router.DELETE("/programEntity/:id", programEntityController.DeleteProgramEntity)

	// SECTIONS

	sectionRepository := repositories.NewSectionRepository(dbHandler)

	sectionService := services.NewSectionService(sectionRepository)

	sectionController := controllers.NewSectionController(sectionService)

	router.GET("/sections", sectionController.GetListSection)

	// SECTION DETAIL

	sectionDetailRepository := repositories.NewSectionDetailRepository(dbHandler)

	sectionDetailService := services.NewSectionDetailService(sectionDetailRepository)

	sectionDetailController := controllers.NewSectionDetailController(sectionDetailService)

	router.GET("/sectionDetail", sectionDetailController.GetListSectionDetail)

	// SECTION DETAIL MATERIAL

	sectionDetailMaterialRepository := repositories.NewSectionDetailMaterialRepository(dbHandler)

	sectionDetailMaterialService := services.NewSectionDetailMaterialService(sectionDetailMaterialRepository)

	sectionDetailMaterialController := controllers.NewSectionDetailMaterialController(sectionDetailMaterialService)

	router.GET("/sectionDetailMaterial", sectionDetailMaterialController.GetListSectionDetailMaterial)

	// PROGRAM ENTITY DESC

	progEntityDescRepository := repositories.NewProgEntityDescRepository(dbHandler)

	progEntityDescService := services.NewProgEntityDescService(progEntityDescRepository)

	progEntityDescController := controllers.NewProgEntityDescController(progEntityDescService)

	router.GET("/progEntityDesc", progEntityDescController.GetListProgEntityDesc)

	// PROGRAM REVIEWS

	progReviewsRepository := repositories.NewProgReviewsRepository(dbHandler)

	progReviewsService := services.NewProgReviewsService(progReviewsRepository)

	progReviewsController := controllers.NewProgReviewsController(progReviewsService)

	router.GET("/progReviews", progReviewsController.GetListProgReviews)

	// GROUP

	groupRepository := repositories.NewProgramEntityRepository(dbHandler)

	groupService := services.NewGroupService(groupRepository)

	groupController := controllers.NewGroupController(groupService)

	//router.GET("/group", groupController.GROUP)
	router.GET("/gabung", groupController.Gabung)

	return HttpServer{
		config:                  config,
		router:                  router,
		programEntityController: programEntityController,
		// progEntityDescController: progEntityDescController,
		// progReviewsController:    progReviewsController,
		// sections:                 sectionController,
		// sectionDetail:            sectionDetailController,
		// sectionDetailMaterial:    sectionDetailMaterialController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
