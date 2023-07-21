package usersService

import (
	"codeid.revampacademy/repositories/usersRepository"
)

type ServiceManager struct {
	UserService
	UserEmailService
	UserPhoneService
	SignUpService
	UserExperienceService
}

// constructor
func NewServiceManager(repoMgr *usersRepository.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		UserService: *NewUserService(&repoMgr.UserRepository),
		UserEmailService: *NewUserEmailService(&repoMgr.UserEmailRepository),
		UserPhoneService: *NewUserPhoneService(&repoMgr.UserPhoneRepository),
		SignUpService: *NewSignUpService(&repoMgr.SignUpRepository),
		UserExperienceService: *NewUserExperienceService(&repoMgr.UserExperienceRepository),
	}
}