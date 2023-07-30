package repositories

import "database/sql"

type RepositoriesManager struct {
	PaymentAccountRepository
	PaymentBankRepository
	PaymentFintechRepository
	PaymentTopupRepository
	PaymentTransactionRepository
}

// constructor
func NewRepositoriesManager(dbHandler *sql.DB) *RepositoriesManager {
	return &RepositoriesManager{
		PaymentAccountRepository:     *NewPaymentAccountRepository(dbHandler),
		PaymentBankRepository:        *NewPaymentBankRepository(dbHandler),
		PaymentFintechRepository:     *NewPaymentFintechRepository(dbHandler),
		PaymentTopupRepository:       *NewPaymentTopupRepository(dbHandler),
		PaymentTransactionRepository: *NewPaymentTransactionRepository(dbHandler),
	}
}
