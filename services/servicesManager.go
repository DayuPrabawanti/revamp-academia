package services

import "codeid.revampacademy/repositories"

type ServiceManager struct {
	PaymentBankService
	PaymentFintechService
	PaymentTransactionService
}

// constructor
func NewServicesManager(repoMgr *repositories.RepositoriesManager) *ServiceManager {
	return &ServiceManager{
		PaymentBankService:        *NewPaymentBankService(&repoMgr.PaymentBankRepository),
		PaymentFintechService:     *NewPaymentFintechService(&repoMgr.PaymentFintechRepository),
		PaymentTransactionService: *NewPaymentTransactionService(&repoMgr.PaymentTransactionRepository),
	}
}
