package bootcampController

import (
	"codeid.revampacademy/services/bootcampService"
)

type ControllerManager struct {
	BatchController
}

// constructor
func NewControllerManager(serviceMgr *bootcampService.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewBatchController(&serviceMgr.BatchService),
	}
}
