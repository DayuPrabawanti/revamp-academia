package server

import (
	controllers "codeid.revampacademy/controllers/curriculumControllers"
	"github.com/gin-gonic/gin"
)

func InitRouterCurriculum(router *gin.Engine, controllerMrg *controllers.ControllerManager) *gin.Engine {

	progentityRoute := router.Group("/curriculum")
	{
		progentityRoute.GET("/progentity", controllerMrg.ProgEntityController.GetListProgEntity)
		progentityRoute.GET("/progentity/:id", controllerMrg.ProgEntityController.GetProgEntity)
		progentityRoute.POST("/progentity", controllerMrg.ProgEntityController.CreateProgEntity)
		progentityRoute.PUT("/progentity/:id", controllerMrg.ProgEntityController.UpdateProgEntity)
		progentityRoute.DELETE("/progentity/:id", controllerMrg.ProgEntityController.DeleteProgEntity)

		progentityRoute.GET("/sections", controllerMrg.ProgEntityController.GetListSection)
		progentityRoute.GET("/sections/:id", controllerMrg.ProgEntityController.GetSection)
		progentityRoute.POST("/sections", controllerMrg.ProgEntityController.CreateSection)
		//progentityRoute.PUT("/sections/:id", controllerMrg.ProgEntityController.UpdateSection)

		progentityRoute.GET("/sectiondetail", controllerMrg.ProgEntityController.GetListSectionDetail)

		progentityRoute.GET("/mastercategory", controllerMrg.ProgEntityController.GetListMasterCategory)

		progentityRoute.GET("/gabung", controllerMrg.ProgEntityController.GetListGabung)
		progentityRoute.GET("/gabung/:id", controllerMrg.ProgEntityController.GetGabung)
		progentityRoute.POST("/createallgabung", controllerMrg.ProgEntityController.CreateGabung)

	}

	progentitydescRoute := router.Group("/curriculum")
	{
		progentitydescRoute.GET("/progEntityDesc", controllerMrg.ProgEntityDescController.GetListProgEntityDesc)
		progentitydescRoute.GET("/progEntityDesc/:id", controllerMrg.ProgEntityDescController.GetProgEntityDesc)

	}
	progreviewsRoute := router.Group("/curriculum")
	{
		progreviewsRoute.GET("/progReviews", controllerMrg.ProgReviewsController.GetListProgReviews)
		progreviewsRoute.GET("/progReviews/:id", controllerMrg.ProgReviewsController.GetProgramReviews)
	}
	sectionsetailmaterialRoute := router.Group("/curriculum")
	{
		sectionsetailmaterialRoute.GET("/sectionDetailMaterial", controllerMrg.SectionDetailMaterialController.GetListSectionDetailMaterial)
		sectionsetailmaterialRoute.GET("/sectionDetailMaterial/:id", controllerMrg.SectionDetailMaterialController.GetSectionDetailMaterial)
		sectionsetailmaterialRoute.POST("/sectionDetailMaterial", controllerMrg.SectionDetailMaterialController.CreatesectiondetailMaterial)
		sectionsetailmaterialRoute.PUT("/sectionDetailMaterial/:id", controllerMrg.SectionDetailMaterialController.UpdateSectionDetailMaterial)
		sectionsetailmaterialRoute.DELETE("/sectionDetailMaterial/:id", controllerMrg.SectionDetailMaterialController.DeleteSectionDetailMaterial)

	}
	return router

}
