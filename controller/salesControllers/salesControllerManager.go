package salescontrollers

import (
	salesservice "codeid.revampacademy/service/salesService"
)

type ControllerManager struct {
	ControllerMock
	ControlMock3
	ControlMock6
	ControlMock7
	ControlMockup8
	ControlMock4
}

func NewControllerManager(serviceMgr *salesservice.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewControllerMock(&serviceMgr.ServiceMock),
		*NewMockupApplyController(&serviceMgr.ServiceMock3),
		*NewControlShoppingCart1(&serviceMgr.ServiceMock6),
		*NewControlShoppingCart2(&serviceMgr.ServiceMock7),
		*NewControlShoppingCart3(&serviceMgr.ServiceMockup8),
		*NewMockupApplyController4(&serviceMgr.ServiceMock4),
	}
}
