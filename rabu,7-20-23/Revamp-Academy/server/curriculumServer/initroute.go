package curriculumServer

import (
	controllers "codeid.revampacademy/controllers/curriculumControllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, controllerMrg *controllers.ControllerManager) *gin.Engine {

	progentityRoute := router.Group("/curriculum")
	{
		progentityRoute.GET("/progentity", controllerMrg.ProgEntityController.GetListProgEntity)
		progentityRoute.GET("/progentity/:id", controllerMrg.ProgEntityController.GetProgEntity)
		progentityRoute.POST("/progentity", controllerMrg.ProgEntityController.CreateProgEntity)
		progentityRoute.PUT("/progentity/:id", controllerMrg.ProgEntityController.UpdateProgEntity)
		progentityRoute.DELETE("/progentity/:id", controllerMrg.ProgEntityController.DeleteProgEntity)

		progentityRoute.GET("/sections", controllerMrg.ProgEntityController.GetListSection)
		progentityRoute.GET("/sections/:id", controllerMrg.ProgEntityController.GetSection)
		progentityRoute.POST("/section", controllerMrg.ProgEntityController.CreateSection)

		progentityRoute.GET("/sectiondetail", controllerMrg.ProgEntityController.GetListSectionDetail)

		progentityRoute.GET("/mastercategory", controllerMrg.ProgEntityController.GetListMasterCategory)

		progentityRoute.GET("/gabung", controllerMrg.ProgEntityController.GetListGabung)
		progentityRoute.GET("/gabung/:id", controllerMrg.ProgEntityController.GetGabung)
		progentityRoute.POST("/createallgabung", controllerMrg.ProgEntityController.CreateGabung)
	}
	return router

}
