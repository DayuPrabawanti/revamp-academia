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
	SignUpEmployeeService
	UserExperienceService
	UserMediaService
	UserAddressService
	masterService.MasterAddressService
	UserEducationService
	UserLicenseService
	UserSkillService
	UserListProfileService
	EditUsernameService
	EditPasswordService
	EditEmailService
	EditPhoneService
	AddAddressService
	EditEducationService
	EditExperienceService
	EditSkillService
}

// constructor
func NewServiceManager(repoMgr *usersRepository.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		UserService:           *NewUserService(&repoMgr.UserRepository),
		UserEmailService:      *NewUserEmailService(&repoMgr.UserEmailRepository),
		UserPhoneService:      *NewUserPhoneService(&repoMgr.UserPhoneRepository),
		SignUpService:         *NewSignUpService(&repoMgr.SignUpRepository),
		SignUpEmployeeService: *NewSignUpEmployeeService(&repoMgr.SignUpEmployeeRepository),
		UserExperienceService: *NewUserExperienceService(&repoMgr.UserExperienceRepository),
		UserMediaService:      *NewUserMediaService(&repoMgr.UserMediaRepository),
		UserAddressService:    *NewUserAddressService(&repoMgr.UserAddressRepository),
		MasterAddressService:  *masterService.NewMasterAddressService(&repoMgr.MasterAddressRepository),
		UserEducationService:  *NewUserEducationService(&repoMgr.UserEducationRepository),
		UserLicenseService:    *NewUserLicenseService(&repoMgr.UserLicenseRepository),
		UserSkillService:      *NewUserSkillService(&repoMgr.UserSkillRepository),
		UserListProfileService: *NewUserListProfileService(&repoMgr.UserListProfileRepository),
		EditUsernameService:   *NewEditUsernameService(&repoMgr.EditUsernameRepository),
		EditPasswordService: *NewEditUserPasswordService(&repoMgr.EditPasswordRepository),
		EditEmailService: *NewEditUserEmailService(&repoMgr.EditEmailRepository),
		EditPhoneService: *NewEditUserPhoneService(&repoMgr.EditPhoneRepository),
		AddAddressService:     *NewAddAddressService(&repoMgr.EditAddressRepository),
		EditEducationService: *NewEditUserEducationService(&repoMgr.EditEducationRepository),
		EditExperienceService: *NewEditUserExperienceService(&repoMgr.EditExperienceRepository),
		EditSkillService: *NewEditUserSkillService(&repoMgr.EditSkillRepository),
	}
}
