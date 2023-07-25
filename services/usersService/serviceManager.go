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
	UserMediaService
	UserAddressService
	UserLicenseService
	UserSkillService
	UserEducationService
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
		UserAddressService:    *NewUserAddressService(&repoMgr.UserAddressRepository),
		UserLicenseService:    *NewUserLicenseService(&repoMgr.UserLicenseRepository),
		UserSkillService:      *NewUserSkillService(&repoMgr.UserSkillRepository),
		UserEducationService:  *NewUserEducationService(&repoMgr.UserEducationRepository),
	}
}
