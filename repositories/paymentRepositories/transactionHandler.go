package repositories

import (
	"context"
	"database/sql"
)

func BeginTransaction(repoMgr *RepositoriesManager) error {
	return repoMgr.PaymentTransactionRepository.BeginTransaction()
}

func RollbackTransaction(repoManager *RepositoriesManager) error {
	return repoManager.PaymentTransactionRepository.RollbackTransaction()
}

func CommitTransaction(repoManager *RepositoriesManager) error {
	return repoManager.PaymentTransactionRepository.CommitTransaction()
}

func (repo *PaymentTransactionRepository) BeginTransaction() error {
	ctx := context.Background()
	transaction, err := repo.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	repo.transaction = transaction
	return nil
}

func (repo *PaymentTransactionRepository) RollbackTransaction() error {
	if repo.transaction != nil {
		err := repo.transaction.Rollback()
		repo.transaction = nil
		return err
	}
	return nil
}

func (repo *PaymentTransactionRepository) CommitTransaction() error {
	if repo.transaction != nil {
		err := repo.transaction.Commit()
		repo.transaction = nil
		return err
	}
	return nil
}
