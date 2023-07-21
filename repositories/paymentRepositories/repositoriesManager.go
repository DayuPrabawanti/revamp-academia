package repositories

import "database/sql"

type RepositoriesManager struct {
	PaymentAccountRepository
	PaymentFintechRepository
	PaymentTopupRepository
}

// Constructor
func NewRepositoriesManager(dbHandler *sql.DB) *RepositoriesManager {
	return &RepositoriesManager{
		*NewPaymentAccountRepository(dbHandler),
		*NewPaymentFintechRepository(dbHandler),
		*NewPaymentTopupRepository(dbHandler),
	}
}
