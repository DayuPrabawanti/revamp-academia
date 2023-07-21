package server

import (
	"database/sql"

	// "codeid.revampacademy/controller"
	"codeid.revampacademy/controller/jobhireController"
	"codeid.revampacademy/service/jobhireService"

	// "codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/jobhireRepositories"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func MasterEndPointRoute(cfg *viper.Viper, dbHandler *sql.DB) *gin.Engine {
	masterRepo := jobhireRepositories.NewMasterRepo(dbHandler)
	masterService := jobhireService.NewMasterService(masterRepo)
	masterController := jobhireController.NewMasterController(masterService)

	//set router from gin first
	router := gin.Default()

	//make endpoint route
	router.GET("/listaddress", masterController.GetListAddressControl)

	return router
}
