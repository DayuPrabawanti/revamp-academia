package curriculumcontrollers

import services "codeid.revampacademy/services/curriculumServices"

type ControllerManager struct {
	ProgEntityController
	ProgEntityDescController
	ProgReviewsController
	SectionDetailMaterialController
	SectionDetailController
	CurriculumController
}

// constructor
func NewControllerManager(serviceMar *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewProgEntityController(&serviceMar.ProgEntityService),
		*NewProgEntityDescController(&serviceMar.ProgEntityDescService),
		*NewProgReviewsController(&serviceMar.ProgReviewService),
		*NewSectionDetailMaterialController(&serviceMar.SectionDetailMaterialService),
		*NewSectionDetailController(&serviceMar.SectionDetailService),
		*NewCurriculumController(&serviceMar.CurriculumService),
	}
}
