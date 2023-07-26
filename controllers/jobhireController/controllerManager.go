package jobhireController

import "codeid.revampacademy/services/jobhireService"

type ControllerManager struct {
	CategoryController
	JobHireController
	MasterController
}

func NewControllerManager(serviceManager *jobhireService.ServiceManager) *ControllerManager {
	return &ControllerManager{
		CategoryController: *NewCategoryController(&serviceManager.CategoryService),
		JobHireController:  *NewJobControll(&serviceManager.JobService),
		MasterController:   *NewMasterController(&serviceManager.MasterService),
	}
}
