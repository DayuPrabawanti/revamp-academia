package usersRepository

import "database/sql"

type RepositoryManager struct {
	UserRepository
	UserEmailRepository
	UserPhoneRepository
	SignUpRepository
	UserExperienceRepository
}

// constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewUserRepository(dbHandler),
		*NewUserEmailRepository(dbHandler),
		*NewUserPhoneRepository(dbHandler),
		*NewSignUpRepository(dbHandler),
		*NewUserExperienceRepository(dbHandler),
	}
}