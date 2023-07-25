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
	UserExperienceRepository
	UserMediaRepository
	UserAddressRepository
	masterRepository.MasterAddressRepository
	UserEducationRepository
	UserLicenseRepository
	UserSkillRepository
}

// constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewUserRepository(dbHandler),
		*NewUserEmailRepository(dbHandler),
		*NewUserPhoneRepository(dbHandler),
		*NewSignUpRepository(dbHandler),
		*NewUserExperienceRepository(dbHandler),
		*NewUserMediaRepository(dbHandler),
		*NewUserAddressRepository(dbHandler),
		*masterRepository.NewMasteraddressRepository(dbHandler),
		*NewUserEducationRepository(dbHandler),
		*NewUserLicenseRepository(dbHandler),
		*NewUserSkillRepository(dbHandler),
	}
}
