package controllers

import "codeid.revampacademy/services"

type ControllersManager struct {
	PaymentBankController
	PaymentFintechController
	PaymentTransactionController
}

// consturctor
func NewControllersManager(serviceMgr *services.ServiceManager) *ControllersManager {
	return &ControllersManager{ // TODO: implement me!
		*NewPaymentBankController(&serviceMgr.PaymentBankService),
		*NewPaymentFintechController(&serviceMgr.PaymentFintechService),
		*NewPaymentTransactionController(&serviceMgr.PaymentTransactionService),
	}
}
