package hrController

import (
	"codeid.revampacademy/services/hrService"
)

type ControllerManager struct {
	DepartmentController
	DepartmentHistoryController
	ClientContractController
	EmployeeController
	PayHistoryController
	TalentsMockupController
	TalentsDetailMockupController
}

// constructor
func NewControllerManager(serviceMgr *hrService.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewDepartmentController(&serviceMgr.DepartmentService),
		*NewDepartmentHistoryController(&serviceMgr.DepartmentHistoryService),
		*NewClientContractController(&serviceMgr.ClientContractService),
		*NewEmployeeController(&serviceMgr.EmployeeService),
		*NewPayHistoryController(&serviceMgr.PayHistoryService),
		*NewTalentMockupController(&serviceMgr.TalentsMockupService),
		*NewTalentDetailMockupController(&serviceMgr.TalentsDetailMockupService),
	}
}
