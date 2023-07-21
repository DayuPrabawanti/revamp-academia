package paymentControllers

import services "codeid.revampacademy/services/paymentServices"

type ControllersManager struct {
	PaymentAccountController
	PaymentBankController
	PaymentFintechController
	PaymentTopupController
	PaymentTransactionController
}

// Constructor
func NewControllersManager(serviceMgr *services.ServiceManager) *ControllersManager {
	return &ControllersManager{
		*NewPaymentAccountController(&serviceMgr.PaymentAccountService),
		*NewPaymentBankController(&serviceMgr.PaymentBankService),
		*NewPaymentFintechController(&serviceMgr.PaymentFintechService),
		*NewPaymentTopupController(&serviceMgr.PaymentTopupService),
		*NewPaymentTransactionController(&serviceMgr.PaymentTransactionService),
	}
}
