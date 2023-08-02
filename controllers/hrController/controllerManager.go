package hrController

import "codeid.revampacademy/services/hrService"

type ControllerManager struct {
	ClientContractController
	DepartmentController
	DepartmentHistoryController
	EmployeeController
	PayHistoryController
	TalentsDetailMockupController
	TalentsMockupController
	EmployeeMockupController
}

// constructor
func NewControllerManager(serviceMgr *hrService.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewClientContractController(&serviceMgr.ClientContractService),
		*NewDepartmentController(&serviceMgr.DepartmentService),
		*NewDepartmentHistoryController(&serviceMgr.DepartmentHistoryService),
		*NewEmployeeController(&serviceMgr.EmployeeService),
		*NewPayHistoryController(&serviceMgr.PayHistoryService),
		*NewTalentDetailMockupController(&serviceMgr.TalentsDetailMockupService),
		*NewTalentMockupController(&serviceMgr.TalentsMockupService),
		*NewEmployeeMockupController(&serviceMgr.EmployeeMockupService),
	}
}
