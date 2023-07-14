package server

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers"
	repo "codeid.revampacademy/repositories/curriculum"
	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config               *viper.Viper
	router               *gin.Engine
	progentityController *controllers.ProgEntityController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	progentityRepository := repo.NewProgEntityRepository(dbHandler)
	progentityServices := services.NewProgEntityService(progentityRepository)
	progentityController := controllers.NewProgEntityController(progentityServices)

	router := gin.Default()
	router.GET("/progentity", progentityController.GetListProgEntity)
	router.GET("/progentity/:id", progentityController.GetProgEntity)
	router.POST("/progentity", progentityController.CreateProgEntity)
	router.PUT("/progentity/:id", progentityController.UpdateProgEntity)
	router.DELETE("/progentity/:id", progentityController.DeleteProgEntity)

	router.GET("/sections", progentityController.GetListSection)
	router.GET("/sectiondetail", progentityController.GetListSectionDetail)
	router.GET("/gabung", progentityController.GetListGabung)

	return HttpServer{
		config:               config,
		router:               router,
		progentityController: progentityController,
	}
}

func (hs HttpServer) GetStart() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Eror di: %v", err)
	}
}
