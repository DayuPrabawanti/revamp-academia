package services

import repositories "codeid.revampacademy/repositories/paymentRepositories"

type ServiceManager struct {
	PaymentAccountService
	PaymentFintechService
	PaymentTopupService
}

// Constructor
func NewServicesManager(repoMgr *repositories.RepositoriesManager) *ServiceManager {
	return &ServiceManager{
		PaymentAccountService: *NewPaymentAccountService(&repoMgr.PaymentAccountRepository),
		PaymentFintechService: *NewPaymentFintechService(&repoMgr.PaymentFintechRepository),
		PaymentTopupService:   *NewPaymentTopupService(&repoMgr.PaymentTopupRepository),
	}
}
