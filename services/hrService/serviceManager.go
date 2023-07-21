package hrService

import (
	"codeid.revampacademy/repositories/hrRepository"
)

type ServiceManager struct {
	DepartmentService
	DepartmentHistoryService
	ClientContractService
	EmployeeService
	PayHistoryService
	TalentsMockupService
	TalentsDetailMockupService
}

func NewServiceManager(repoMgr *hrRepository.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		DepartmentService:          *NewDepartmentService(&repoMgr.DepartmentRepository),
		DepartmentHistoryService:   *NewDepartmentHistoryService(&repoMgr.DepartmentHistoryRepository),
		ClientContractService:      *NewClientContractService(&repoMgr.ClientContractRepository),
		EmployeeService:            *NewEmployeeService(&repoMgr.EmployeeRepository),
		PayHistoryService:          *NewPayHistoryService(&repoMgr.PayHistoryRepository),
		TalentsMockupService:       *NewTalentMockupService(&repoMgr.TalentsMockupRepository),
		TalentsDetailMockupService: *NewTalentDetailMockupService(&repoMgr.TalentsDetailMockupRepository),
	}
}
