package usersController

import "codeid.revampacademy/services/usersService"

type ControllerManager struct {
	UserController
	UserEmailController
	UserPhoneController
	SignUpController
}

// constructor
func NewControllerManager(serviceMgr *usersService.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewUserController(&serviceMgr.UserService),
		*NewUserEmailController(&serviceMgr.UserEmailService),
		*NewUserPhoneController(&serviceMgr.UserPhoneService),
		*NewSignUpController(&serviceMgr.SignUpService),
	}
}