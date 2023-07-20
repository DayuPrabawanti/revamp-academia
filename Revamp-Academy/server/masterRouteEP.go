package server

import (
	"database/sql"

	"codeid.revampacademy/controller"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func MasterEndPointRoute(cfg *viper.Viper, dbHandler *sql.DB) *gin.Engine {
	masterRepo := repositories.NewMasterRepo(dbHandler)
	masterService := service.NewMasterService(masterRepo)
	masterController := controller.NewMasterController(masterService)

	//set router from gin first
	router := gin.Default()

	//make endpoint route
	router.GET("/listaddress", masterController.GetListAddressControl)

	return router
}
