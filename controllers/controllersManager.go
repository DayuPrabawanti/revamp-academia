package controllers

import "codeid.revampacademy/services"

type ControllersManager struct {
	JobPostController
	MasterIndustryController
	JobPostingController
	ApplyProfController
}

// Constructor
func NewControllersManager(serviceMgr *services.ServiceManager) *ControllersManager {
	return &ControllersManager{
		*NewJobPostController(&serviceMgr.JobPostService),
		*NewMasterIndustryController(&serviceMgr.MasterIndustryService),
		*NewJobPostingController(&serviceMgr.JobPostingService),
		*NewApplyProfController(&serviceMgr.ApplyProfService),
	}
}
