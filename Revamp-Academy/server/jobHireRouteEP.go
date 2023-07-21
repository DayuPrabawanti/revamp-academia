package server

import (
	"database/sql"

	// "codeid.revampacademy/controller"
	"codeid.revampacademy/controller/jobhireController"
	"codeid.revampacademy/service/jobhireService"

	// "codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/jobhireRepositories"
	// "codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func JobHireEndpointRoute(cfg *viper.Viper, dbHandler *sql.DB) *gin.Engine {
	categoryRepo := jobhireRepositories.NewCategoryRepo(dbHandler)
	categoryService := jobhireService.NewCategoryService(categoryRepo)
	categoryControl := jobhireController.NewCategoryController(categoryService)

	router := gin.Default()

	//buat router Endpoint
	router.GET("/listJobCategory", categoryControl.GetListCategoryControl)

	return router
}
