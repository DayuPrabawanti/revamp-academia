package services

import repositories "codeid.revampacademy/repositories/paymentRepositories"

type ServiceManager struct {
	PaymentAccountService
	PaymentBankService
	PaymentFintechService
	PaymentTopupService
	PaymentTransactionService
}

// Constructor
func NewServicesManager(repoMgr *repositories.RepositoriesManager) *ServiceManager {
	return &ServiceManager{
		PaymentAccountService: *NewPaymentAccountService(repoMgr),
		PaymentBankService:    *NewPaymentBankService(repoMgr),
		PaymentFintechService: *NewPaymentFintechService(repoMgr),
		PaymentTopupService:   *NewPaymentTopupService(repoMgr),
		// PaymentTransactionService: *NewPaymentTransactionService(&repoMgr.PaymentTransactionRepository),
		PaymentTransactionService: *NewPaymentTransactionService(repoMgr),
	}
}
