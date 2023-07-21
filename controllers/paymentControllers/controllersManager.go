package controllers

import services "codeid.revampacademy/services/paymentServices"

type ControllersManager struct {
	PaymentAccountController
	PaymentFintechController
	PaymentTopupController
}

// Constructor
func NewControllersManager(serviceMgr *services.ServiceManager) *ControllersManager {
	return &ControllersManager{
		*NewPaymentAccountController(&serviceMgr.PaymentAccountService),
		*NewPaymentFintechController(&serviceMgr.PaymentFintechService),
		*NewPaymentTopupController(&serviceMgr.PaymentTopupService),
	}
}
