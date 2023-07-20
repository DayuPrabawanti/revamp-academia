package curriculumcontrollers

import services "codeid.revampacademy/services/curriculumServices"

type ControllerManager struct {
	ProgEntityController
}

// constructor
func NewControllerManager(serviceMar *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewProgEntityController(&serviceMar.ProgEntityService),
	}
}
