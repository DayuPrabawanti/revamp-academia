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

	userEmailRoute := routers.Group("/users/email")
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

	userSignup := routers.Group("/api/users")
	{
		// Router endpoint (url) http User Sign Up
		userSignup.POST("/signup", controllerMgr.SignUpController.CreateSignUpUser)
		userSignup.POST("/signupEmployee", controllerMgr.SignUpEmployeeController.CreateSignUpEmployee)
	}

	userProfile := routers.Group("/api/users/profile")
	{
		// Router endpoint (url) http User Show Profile
		// userProfile.GET("/view/:id", controllerMgr.UserListProfileController.GetProfileUser)
		userProfile.GET("/view/:id", controllerMgr.UserListProfileController.GetProfile)
	}

	editUser := routers.Group("/api/users/profile")
	{
		// Router endpoint (url) http User Edit Username
		editUser.GET("/edit/:id", controllerMgr.EditUsernameController.GetUsername)
		editUser.PUT("/edit/:id", controllerMgr.EditUsernameController.EditUsername)
		editUser.GET("/password/:id", controllerMgr.EditUserPasswordController.GetPassword)
		editUser.PUT("/password/:id", controllerMgr.EditUserPasswordController.UpdatePassword)
		editUser.POST("/email/:id", controllerMgr.EditEmailController.AddEmail)
		editUser.POST("/phone/:id", controllerMgr.EditPhoneController.AddPhone)
		editUser.POST("/address/:id", controllerMgr.EditAddressController.AddAddress)
		editUser.POST("/education/:id", controllerMgr.EditEducationController.AddEducation)
		editUser.POST("/experience/:id", controllerMgr.EditExperienceController.AddExperience)
		editUser.POST("/skill/:id", controllerMgr.EditSkillController.AddSkill)

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

	UserEducationRoute := routers.Group("/users/education")
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
