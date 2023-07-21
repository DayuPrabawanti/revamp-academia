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

	router.GET("/sections/:id", sectionController.GetSections)

	router.POST("/sections", sectionController.Createsections)

	router.PUT("/sections/:id", sectionController.UpdateSections)

	router.DELETE("/sections/:id", sectionController.DeleteSections)

	// SECTION DETAIL

	sectionDetailRepository := repositories.NewSectionDetailRepository(dbHandler)

	sectionDetailService := services.NewSectionDetailService(sectionDetailRepository)

	sectionDetailController := controllers.NewSectionDetailController(sectionDetailService)

	router.GET("/sectionDetail", sectionDetailController.GetListSectionDetail)

	router.GET("/sectionDetail/:id", sectionDetailController.GetSectionDetail)

	router.POST("/sectionDetail", sectionDetailController.CreateSectionDetail)

	router.PUT("/sectionDetail/:id", sectionDetailController.UpdateSectionDetail)

	router.DELETE("/sectionDetail/:id", sectionDetailController.DeleteSectionDetail)

	// SECTION DETAIL MATERIAL

	sectionDetailMaterialRepository := repositories.NewSectionDetailMaterialRepository(dbHandler)

	sectionDetailMaterialService := services.NewSectionDetailMaterialService(sectionDetailMaterialRepository)

	sectionDetailMaterialController := controllers.NewSectionDetailMaterialController(sectionDetailMaterialService)

	router.GET("/sectionDetailMaterial", sectionDetailMaterialController.GetListSectionDetailMaterial)

	router.GET("/sectionDetailMaterial/:id", sectionDetailMaterialController.GetSectionDetailMaterial)

	router.POST("/sectionDetailMaterial", sectionDetailMaterialController.CreatesectiondetailMaterial)

	router.PUT("/sectionDetailMaterial/:id", sectionDetailMaterialController.UpdateSectionDetailMaterial)

	router.DELETE("/sectionDetailMaterial/:id", sectionDetailMaterialController.DeleteSectionDetailMaterial)

	// PROGRAM ENTITY DESC

	progEntityDescRepository := repositories.NewProgEntityDescRepository(dbHandler)

	progEntityDescService := services.NewProgEntityDescService(progEntityDescRepository)

	progEntityDescController := controllers.NewProgEntityDescController(progEntityDescService)

	router.GET("/progEntityDesc", progEntityDescController.GetListProgEntityDesc)

	router.GET("/progEntityDesc/:id", progEntityDescController.GetProgEntityDesc)

	// router.POST("/sectionDetailMaterial", sectionDetailMaterialController.CreatesectiondetailMaterial)

	// router.PUT("/sectionDetailMaterial/:id", sectionDetailMaterialController.UpdateSectionDetailMaterial)

	// router.DELETE("/sectionDetailMaterial/:id", sectionDetailMaterialController.DeleteSectionDetailMaterial)

	// PROGRAM REVIEWS

	progReviewsRepository := repositories.NewProgReviewsRepository(dbHandler)

	progReviewsService := services.NewProgReviewsService(progReviewsRepository)

	progReviewsController := controllers.NewProgReviewsController(progReviewsService)

	router.GET("/progReviews", progReviewsController.GetListProgReviews)

	router.GET("/progReviews/:id", progReviewsController.GetProgramReviews)

	// GROUP

	groupRepository := repositories.NewProgramEntityRepository(dbHandler)

	groupService := services.NewProgramEntityService(groupRepository)

	groupController := controllers.NewProgramEntityController(groupService)

	router.GET("/group", groupController.GroupList)

	// MASTER.CATEGORY

	masterRepository := repositories.NewProgramEntityRepository(dbHandler)

	masterService := services.NewProgramEntityService(masterRepository)

	masterController := controllers.NewProgramEntityController(masterService)

	router.GET("/masterCategory", masterController.GetListMasterCategory)

	return HttpServer{
		config:                  config,
		router:                  router,
		programEntityController: programEntityController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
