package paymentServices

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
		PaymentAccountService:     *NewPaymentAccountService(&repoMgr.PaymentAccountRepository),
		PaymentBankService:        *NewPaymentBankService(&repoMgr.PaymentBankRepository),
		PaymentFintechService:     *NewPaymentFintechService(&repoMgr.PaymentFintechRepository),
		PaymentTopupService:       *NewPaymentTopupService(&repoMgr.PaymentTopupRepository),
		PaymentTransactionService: *NewPaymentTransactionService(&repoMgr.PaymentTransactionRepository),
	}
}
