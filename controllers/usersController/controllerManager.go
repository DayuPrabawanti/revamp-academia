package usersController

import (
	"codeid.revampacademy/controllers/masterController"
	"codeid.revampacademy/services/usersService"
)

type ControllerManager struct {
	UserController
	UserEmailController
	UserPhoneController
	SignUpController
	SignUpEmployeeController
	UserExperienceController
	UserMediaController
	UserAddressController
	masterController.MasterAddressController
	UserEducationController
	UserLicenseController
	UserSkillController
	UserListProfileController
}

// constructor
func NewControllerManager(serviceMgr *usersService.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewUserController(&serviceMgr.UserService),
		*NewUserEmailController(&serviceMgr.UserEmailService),
		*NewUserPhoneController(&serviceMgr.UserPhoneService),
		*NewSignUpController(&serviceMgr.SignUpService),
		*NewSignUpEmployeeController(&serviceMgr.SignUpEmployeeService),
		*NewUserExperienceController(&serviceMgr.UserExperienceService),
		*NewUserMediaController(&serviceMgr.UserMediaService),
		*NewUseraddressController(&serviceMgr.UserAddressService),
		*masterController.NewMasterAddressController(&serviceMgr.MasterAddressService),
		*NewUserEducationController(&serviceMgr.UserEducationService),
		*NewUserLicenseController(&serviceMgr.UserLicenseService),
		*NewUserSkillController(&serviceMgr.UserSkillService),
		*NewUserListProfileController(&serviceMgr.UserListProfileService),
	}
}
