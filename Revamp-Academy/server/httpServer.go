package server

import (
	// "codeid.revampacademy/config"
	"database/sql"
	"log"

	"codeid.revampacademy/controller"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	categoryController *controller.CategoryController
}

func InitHttpServer(cfg *viper.Viper, dbHandler *sql.DB) HttpServer {
	categoryRepo := repositories.NewCategoryRepo(dbHandler)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryControl := controller.NewCategoryController(categoryService)

	router := gin.Default()

	//buat router Endpoint
	router.GET("/listJobCategory", categoryControl.GetListCategoryControl)

	return HttpServer{
		config:             cfg,
		router:             router,
		categoryController: categoryControl,
	}
}

// running for gin server
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error While starting HTTP Server : %v", err)
	}
}
