package usersServer

import (
	"codeid.revampacademy/controllers/usersController"
	"github.com/gin-gonic/gin"
)

func InitRouter(routers *gin.Engine, controllerMgr *usersController.ControllerManager) *gin.Engine {

	userRoute := routers.Group("/users")
	{
		// Router endpoint (url) http category
		userRoute.GET("/", controllerMgr.UserController.GetListUser)
		userRoute.GET("/:id", controllerMgr.UserController.GetUser)
		userRoute.POST("/", controllerMgr.UserController.CreateUser)
		userRoute.PUT("/:id", controllerMgr.UserController.UpdateUser)
		userRoute.DELETE("/:id", controllerMgr.UserController.DeleteUser)
	}

	userEmailRoute := routers.Group("/usersemail")
	{
		// Router endpoint (url) http category
		userEmailRoute.GET("/", controllerMgr.UserEmailController.GetListUsersEmail)
		userEmailRoute.GET("/:id", controllerMgr.UserEmailController.GetEmail)
		userEmailRoute.POST("/", controllerMgr.UserEmailController.CreateEmail)
		userEmailRoute.PUT("/:id", controllerMgr.UserEmailController.UpdateEmail)
		userEmailRoute.DELETE("/:id", controllerMgr.UserEmailController.DeleteEmail)
	}

	userPhoneRoute := routers.Group("/usersphone")
	{
		// Router endpoint (url) http category
		userPhoneRoute.GET("/", controllerMgr.UserPhoneController.GetListUsersPhone)
		userPhoneRoute.GET("/:id", controllerMgr.UserPhoneController.GetPhone)
		userPhoneRoute.POST("/", controllerMgr.UserPhoneController.CreatePhones)
		userPhoneRoute.PUT("/:id", controllerMgr.UserPhoneController.UpdatePhone)
		userPhoneRoute.DELETE("/:id", controllerMgr.UserPhoneController.DeletePhones)
	}

	userSignup := routers.Group("/userssignup")
	{
		// Router endpoint (url) http category
		// userSignup.GET("/", controllerMgr.SignUpController.GetListCategory)
		// userSignup.GET("/:id", controllerMgr.CategoryController.GetCategory)
		userSignup.POST("/", controllerMgr.SignUpController.CreateSignUp)
		// userSignup.PUT("/:id", controllerMgr.CategoryController.UpdateCategory)
		// userSignup.DELETE("/:id", controllerMgr.CategoryController.DeleteCategory)
	}

	userExperienceRoute := routers.Group("/usersexperience")
	{
		// Router endpoint (url) http category
		userExperienceRoute.GET("/", controllerMgr.UserExperienceController.GetListUserExperience)
		userExperienceRoute.GET("/:id", controllerMgr.UserExperienceController.GetExperience)
		userExperienceRoute.POST("/", controllerMgr.UserExperienceController.CreateExperience)
		userExperienceRoute.PUT("/:id", controllerMgr.UserExperienceController.UpdateExperience)
		userExperienceRoute.DELETE("/:id", controllerMgr.UserExperienceController.DeleteExperience)
	}

	userMedia := routers.Group("/usermedia")
	{
		// Router endpoint userMedia
		userMedia.GET("/", controllerMgr.UserMediaController.GetListUserMedia)
		userMedia.GET("/:id", controllerMgr.UserMediaController.GetUserMedia)
		userMedia.POST("/", controllerMgr.UserMediaController.CreateUserMedia)

		userMedia.PUT("/:id", controllerMgr.UserMediaController.UpdateMedia)
		userMedia.DELETE("/:id", controllerMgr.UserMediaController.DeleteMedia)
	}

	return routers
}
