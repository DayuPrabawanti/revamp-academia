package salesController

import (
	"codeid.revampacademy/services/salesServices"
)

type ControllerManager struct {
	CartItemsController
	ControllerMock
	FintechController
	SpecialOfferController
	EducationController
}

// constructor
func NewControllerManager(serviceMgr *salesServices.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewCartItemsController(&serviceMgr.CartItemsService),
		*NewControllerMock(&serviceMgr.ServiceMock),
		*NewFintechController(&serviceMgr.FintechService),
		*NewSpecialController(&serviceMgr.SpecialOfferService),
		*NewEducationController(&serviceMgr.EducationService),
	}
}
