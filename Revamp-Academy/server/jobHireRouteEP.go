package server

import (
	"database/sql"

	"codeid.revampacademy/controller"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func JobHireEndpointRoute(cfg *viper.Viper, dbHandler *sql.DB) *gin.Engine {
	categoryRepo := repositories.NewCategoryRepo(dbHandler)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryControl := controller.NewCategoryController(categoryService)

	router := gin.Default()

	//buat router Endpoint
	router.GET("/listJobCategory", categoryControl.GetListCategoryControl)

	return router
}
