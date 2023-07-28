package hrService

import "codeid.revampacademy/repositories/hrRepository"

type ServiceManager struct {
	ClientContractService
	DepartmentHistoryService
	DepartmentService
	EmployeeService
	PayHistoryService
	TalentsDetailMockupService
	TalentsMockupService
	EmployeeMockupService
}

// constructor
func NewServiceManager(repoMgr *hrRepository.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		ClientContractService:      *NewClientContractService(&repoMgr.ClientContractRepository),
		DepartmentHistoryService:   *NewDepartmentHistoryService(&repoMgr.DepartmentHistoryRepository),
		DepartmentService:          *NewDepartmentService(&repoMgr.DepartmentRepository),
		EmployeeService:            *NewEmployeeService(&repoMgr.EmployeeRepository),
		PayHistoryService:          *NewPayHistoryService(&repoMgr.PayHistoryRepository),
		TalentsDetailMockupService: *NewTalentDetailMockupService(&repoMgr.TalentsDetailMockupRepository),
		TalentsMockupService:       *NewTalentMockupService(&repoMgr.TalentsMockupRepository),
		EmployeeMockupService:      *NewEmployeeMockupService(&repoMgr.EmployeeMockupRepository),
	}
}
