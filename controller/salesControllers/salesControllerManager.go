package salescontrollers

import (
	salesservice "codeid.revampacademy/service/salesService"
)

type ControllerManager struct {
	ControllerMock
	ControlMock3
}

func NewControllerManager(serviceMgr *salesservice.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewControllerMock(&serviceMgr.ServiceMock),
		*NewMockupApplyController(&serviceMgr.ServiceMock3),
	}
}
