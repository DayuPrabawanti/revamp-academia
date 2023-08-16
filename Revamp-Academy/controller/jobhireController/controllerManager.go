package jobhireController

import "codeid.revampacademy/service/jobhireService"

type ControllerManager struct {
	// CategoryController
	JobHireController
	// MasterController
}

func NewControllerManager(serviceManager *jobhireService.ServiceManager) *ControllerManager {
	return &ControllerManager{
		// CategoryController: *NewCategoryController(&serviceManager.CategoryService),
		*NewJobControll(&serviceManager.JobService),
		// MasterController:   *NewMasterController(&serviceManager.MasterService),
	}
}
