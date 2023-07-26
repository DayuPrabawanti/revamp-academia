package paymentRepositories

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
	return &RepositoriesManager{ // TODO: implement repository instances here and pass db handler to them
		*NewPaymentAccountRepository(dbHandler),
		*NewPaymentBankRepository(dbHandler),
		*NewPaymentFintechRepository(dbHandler),
		*NewPaymentTopupRepository(dbHandler),
		*NewPaymentTransactionRepository(dbHandler),
	}
}
