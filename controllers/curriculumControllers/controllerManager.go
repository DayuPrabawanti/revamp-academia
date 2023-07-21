package curriculumcontrollers

import services "codeid.revampacademy/services/curriculumServices"

type ControllerManager struct {
	ProgEntityController
	ProgEntityDescController
	ProgReviewsController
	SectionDetailMaterialController
}

// constructor
func NewControllerManager(serviceMar *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewProgEntityController(&serviceMar.ProgEntityService),
		*NewProgEntityDescController(&serviceMar.ProgEntityDescService),
		*NewProgReviewsController(&serviceMar.ProgReviewService),
		*NewSectionDetailMaterialController(&serviceMar.SectionDetailMaterialService),
	}
}
