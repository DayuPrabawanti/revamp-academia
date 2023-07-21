package server

import (
	"codeid.revampacademy/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, controllerMgr *controllers.ControllersManager) *gin.Engine {

	categoryRoute := router.Group("/category")
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

	categoryRoute.GET("/jobCategory/:id", controllerMgr.JobPostController.GetJobPostHttp)
        categoryRoute.GET("/jobCategory", controllerMgr.JobPostController.ListJobPostHttp)
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
    categoryRoute.GET("/masterIndustry/:id", controllerMgr.MasterIndustryController.GetMasterIndustryHttp)
        categoryRoute.GET("/masterIndustry", controllerMgr.MasterIndustryController.ListMasterIndustryHttp)


		categoryRoute.GET("/jobposting", controllerMgr.JobPostingController.ListJobPostingHttp)
	}
	return router
}
