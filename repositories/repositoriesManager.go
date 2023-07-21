package repositories

import "database/sql"

type RepositoriesManager struct {
	PaymentBankRepository
	PaymentFintechRepository
	PaymentTransactionRepository
}

// constructor
func NewRepositoriesManager(dbHandler *sql.DB) *RepositoriesManager {
	return &RepositoriesManager{ // TODO: implement repository instances here and pass db handler to them
		*NewPaymentBankRepository(dbHandler),
		*NewPaymentFintechRepository(dbHandler),
		*NewPaymentTransactionRepository(dbHandler),
	}
}
