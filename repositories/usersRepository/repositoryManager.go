package usersRepository

import (
	"database/sql"

	"codeid.revampacademy/repositories/masterRepository"
)

type RepositoryManager struct {
	UserRepository
	UserEmailRepository
	UserPhoneRepository
	SignUpRepository
	SignUpEmployeeRepository
	UserExperienceRepository
	UserMediaRepository
	UserAddressRepository
	masterRepository.MasterAddressRepository
	UserEducationRepository
	UserLicenseRepository
	UserSkillRepository
	UserListProfileRepository
	EditUsernameRepository
	EditPasswordRepository
	EditEmailRepository
	EditPhoneRepository
	EditAddressRepository
	EditEducationRepository
	EditExperienceRepository
	EditSkillRepository
}

// constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewUserRepository(dbHandler),
		*NewUserEmailRepository(dbHandler),
		*NewUserPhoneRepository(dbHandler),
		*NewSignUpRepository(dbHandler),
		*NewSignUpEmployeeRepository(dbHandler),
		*NewUserExperienceRepository(dbHandler),
		*NewUserMediaRepository(dbHandler),
		*NewUserAddressRepository(dbHandler),
		*masterRepository.NewMasteraddressRepository(dbHandler),
		*NewUserEducationRepository(dbHandler),
		*NewUserLicenseRepository(dbHandler),
		*NewUserSkillRepository(dbHandler),
		*NewUserListProfileRepository(dbHandler),
		*NewEditUsernameRepository(dbHandler),
		*NewEditPasswordRepository(dbHandler),
		*NewEditEmailRepository(dbHandler),
		*NewEditPhoneRepository(dbHandler),
		*NewEditAddressRepository(dbHandler),
		*NewEditEducationRepository(dbHandler),
		*NewEditExperienceRepository(dbHandler),
		*NewEditSkillRepository(dbHandler),
	}
}
