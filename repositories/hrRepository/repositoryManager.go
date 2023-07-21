package hrRepository

import "database/sql"

type RepositoryManager struct {
	DepartmentRepository
	DepartmentHistoryRepository
	ClientContractRepository
	EmployeeRepository
	PayHistoryRepository
	TalentsMockupRepository
	TalentsDetailMockupRepository
}

// Constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewDepartmentRepository(dbHandler),
		*NewDepartmentHistoryRepository(dbHandler),
		*NewClientContractRepository(dbHandler),
		*NewEmployeeRepository(dbHandler),
		*NewPayHistoryRepository(dbHandler),
		*NewTalentMockupRepository(dbHandler),
		*NewTalentDetailMockupRepository(dbHandler),
	}
}
