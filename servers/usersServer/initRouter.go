package usersServer

import (
	"codeid.revampacademy/controllers/usersController"
	"github.com/gin-gonic/gin"
)

func InitRouter(routers *gin.Engine, controllerMgr *usersController.ControllerManager) *gin.Engine {

	userRoute := routers.Group("/users")
	{
		// Router endpoint (url) http User
		userRoute.GET("/", controllerMgr.UserController.GetListUser)
		userRoute.GET("/:id", controllerMgr.UserController.GetUser)
		userRoute.POST("/", controllerMgr.UserController.CreateUser)
		userRoute.PUT("/:id", controllerMgr.UserController.UpdateUser)
		userRoute.DELETE("/:id", controllerMgr.UserController.DeleteUser)
	}

	userEmailRoute := routers.Group("/usersemail")
	{
		// Router endpoint (url) http User Email
		userEmailRoute.GET("/", controllerMgr.UserEmailController.GetListUsersEmail)
		userEmailRoute.GET("/:id", controllerMgr.UserEmailController.GetEmail)
		userEmailRoute.POST("/", controllerMgr.UserEmailController.CreateEmail)
		userEmailRoute.PUT("/:id", controllerMgr.UserEmailController.UpdateEmail)
		userEmailRoute.DELETE("/:id", controllerMgr.UserEmailController.DeleteEmail)
	}

	userPhoneRoute := routers.Group("/usersphone")
	{
		// Router endpoint (url) http User Phone
		userPhoneRoute.GET("/", controllerMgr.UserPhoneController.GetListUsersPhone)
		userPhoneRoute.GET("/:id", controllerMgr.UserPhoneController.GetPhone)
		userPhoneRoute.POST("/", controllerMgr.UserPhoneController.CreatePhones)
		userPhoneRoute.PUT("/:id", controllerMgr.UserPhoneController.UpdatePhone)
		userPhoneRoute.DELETE("/:id", controllerMgr.UserPhoneController.DeletePhones)
	}

	userSignup := routers.Group("/users/signup")
	{
		// Router endpoint (url) http User Sign Up
		// userSignup.GET("/", controllerMgr.SignUpController.GetListCategory)
		// userSignup.GET("/:id", controllerMgr.SignUpController.GetCategory)
		userSignup.POST("/", controllerMgr.SignUpController.CreateSignUp)
		// userSignup.PUT("/:id", controllerMgr.SignUpController.UpdateCategory)
		// userSignup.DELETE("/:id", controllerMgr.SignUpController.DeleteCategory)
	}

	userExperienceRoute := routers.Group("/usersexperience")
	{
		// Router endpoint (url) http User Experience
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

	userLicense := routers.Group("/userlicense")
	{
		// Router endpoint userMedia
		userLicense.GET("/", controllerMgr.UserLicenseController.GetListUserLicense)
		userLicense.GET("/:id", controllerMgr.UserLicenseController.GetUsersLicense)
		userLicense.POST("/", controllerMgr.UserLicenseController.CreateUserLicense)

		userLicense.PUT("/:id", controllerMgr.UserLicenseController.UpdateUserLicense)
		userLicense.DELETE("/:id", controllerMgr.UserLicenseController.DeleteLicense)
	}

	userAddressRoute := routers.Group("/usersaddress")
	{
		// Router endpoint (url) http User Address
		userAddressRoute.GET("/", controllerMgr.UserAddressController.GetListUserAddress)
		userAddressRoute.GET("/:id", controllerMgr.UserAddressController.GetAddress)
		userAddressRoute.POST("/", controllerMgr.UserAddressController.CreateAddrees)
		// userAddressRoute.PUT("/:id", controllerMgr.UserAddressController.UpdateExperience)
		// userAddressRoute.DELETE("/:id", controllerMgr.UserAddressController.DeleteExperience)
	}

	UserEducationRoute := routers.Group("/userseducation")
	{
		// Router endpoint (url) http User Address
		UserEducationRoute.GET("/", controllerMgr.UserEducationController.GetListUsersEducation)
		UserEducationRoute.GET("/:id", controllerMgr.UserEducationController.GetUserEducation)
		UserEducationRoute.POST("/", controllerMgr.UserEducationController.CreateUserEducation)
		UserEducationRoute.PUT("/:id", controllerMgr.UserEducationController.UpdateEducation)
		UserEducationRoute.DELETE("/:id", controllerMgr.UserEducationController.DeleteEducation)
	}

	UserSkillRoute := routers.Group("/userskill")
	{
		// Router endpoint (url) http User Address
		UserSkillRoute.GET("/", controllerMgr.UserSkillController.GetListUserSkill)
		UserSkillRoute.GET("/:id", controllerMgr.UserSkillController.GetUsersSkill)
		UserSkillRoute.POST("/", controllerMgr.UserSkillController.CreateUserSkill)
		UserSkillRoute.PUT("/:id", controllerMgr.UserSkillController.UpdateUserSkill)
		UserSkillRoute.DELETE("/:id", controllerMgr.UserSkillController.DeleteSkill)
	}

	masterAddressRoute := routers.Group("/master/address")
	{
		// Router endpoint (url) http master address
		masterAddressRoute.GET("/", controllerMgr.MasterAddressController.GetListMasterAddress)
		// masterAddressRoute.GET("/:id", controllerMgr.MasterAddressController.GetAddress)
		// masterAddressRoute.POST("/", controllerMgr.MasterAddressController.CreateAddrees)
		// masterAddressRoute.PUT("/:id", controllerMgr.MasterAddressController.UpdateExperience)
		// masterAddressRoute.DELETE("/:id", controllerMgr.MasterAddressController

		return routers
	}
}
