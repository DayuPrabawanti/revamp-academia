package usersService

import (
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/services/masterService"
)

type ServiceManager struct {
	UserService
	UserEmailService
	UserPhoneService
	SignUpService
	UserExperienceService
	UserMediaService
	UserAddressService
	masterService.MasterAddressService
}

// constructor
func NewServiceManager(repoMgr *usersRepository.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		UserService:           *NewUserService(&repoMgr.UserRepository),
		UserEmailService:      *NewUserEmailService(&repoMgr.UserEmailRepository),
		UserPhoneService:      *NewUserPhoneService(&repoMgr.UserPhoneRepository),
		SignUpService:         *NewSignUpService(&repoMgr.SignUpRepository),
		UserExperienceService: *NewUserExperienceService(&repoMgr.UserExperienceRepository),
		UserMediaService:      *NewUserMediaService(&repoMgr.UserMediaRepository),
		UserAddressService: 	*NewUserAddressService(&repoMgr.UserAddressRepository),
		MasterAddressService: *masterService.NewMasterAddressService(&repoMgr.MasterAddressRepository),
	}
}
