package server

import (
	"codeid.revampacademy/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, controllerMgr *controllers.ControllersManager) *gin.Engine {

	JobHireRoute := router.Group("/job")
	{
		// Router endpoint
		// jobCategoryRepository := repositories.NewJobCategoryRepository(dbHandler)
		// 	jobCategoryService := services.NewJobCategoryService(jobCategoryRepository)
		// 	jobCategoryController := controllers.NewJobCategoryController(jobCategoryService)

		// jobClientRepository := repositories.NewJobClientRepository(dbHandler)
		// jobClientService := services.NewJobClientService(jobClientRepository)
		// jobClientController := controllers.NewJobClientController(jobClientService)

		// jobPostRepository := repositories.NewJobPostRepository(dbHandler)
		// jobPostService := services.NewJobPostService(jobPostRepository)
		// jobPostController := controllers.NewJobPostController(jobPostService)

		// masterIndustryRepository := repositories.NewMasterIndustryRepository(dbHandler)
		// masterIndustryService := services.NewMasterIndustryService(masterIndustryRepository)
		// masterIndustryController := controllers.NewMasterIndustryController(masterIndustryService)

		JobHireRoute.GET("/jobCategory/:id", controllerMgr.JobPostController.GetJobPostHttp)
		JobHireRoute.GET("/jobCategory", controllerMgr.JobPostController.ListJobPostHttp)
		//         router.POST("/createJobCategory", jobCategoryController.CreateJobCategoryHttp)
		//             router.PUT("/jobCategory/:id", jobCategoryController.UpdateJobCategoryHttp)
		//                 router.DELETE("/category/:id", jobCategoryController.DeleteJobCategoryHttp)

		// jobClient

		// router.GET("/jobClient/:id", jobClientController.GetJobClientHttp)
		//     router.GET("/jobClientList/", jobClientController.ListJobClientHttp)
		//         router.POST("/jobClientCreate", jobClientController.CreateJobClientHttp)
		//             router.PUT("/jobClientUpdate/:id", jobClientController.UpdateJobClientHttp)
		//                 router.DELETE("/jobClientDelete/:id", jobClientController.DeleteJobClientHttp)

		// jobPost

		// categoryRoute.GET("/jobPost/:id", controllerMgr.jobp.GetJobPostHttp)
		//     router.GET("/jobPostList/", jobPostController.ListJobPostHttp)
		//         router.POST("/jobPostCreate", jobPostController.CreateJobPostHttp)
		//             router.PUT("/jobPostUpdate/:id", jobPostController.UpdateJobPostHttp)
		//                 router.DELETE("/jobPostDelete/:id", jobPostController.DeleteJobPostHttp)

		// masterIndustry
		JobHireRoute.GET("/masterIndustry/:id", controllerMgr.MasterIndustryController.GetMasterIndustryHttp)
		JobHireRoute.GET("/masterIndustry", controllerMgr.MasterIndustryController.ListMasterIndustryHttp)

		JobHireRoute.GET("/jobposting", controllerMgr.JobPostingController.GetJobPostingHttp)
		JobHireRoute.GET("/jobpostingList", controllerMgr.JobPostingController.ListJobPostingHttp)

		JobHireRoute.GET("/applyProf", controllerMgr.ApplyProfController.ListApplyProfHttp)

	}
	return router
}
