package repositories

import (
	"database/sql"
)

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
		*NewPaymentAccountRepository(dbHandler),
		*NewPaymentBankRepository(dbHandler),
		*NewPaymentFintechRepository(dbHandler),
		*NewPaymentTopupRepository(dbHandler),
		*NewPaymentTransactionRepository(dbHandler),
	}
}
