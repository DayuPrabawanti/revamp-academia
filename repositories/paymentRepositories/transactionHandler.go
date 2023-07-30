package repositories

import (
	"context"
	"database/sql"
)

func BeginTransaction(repoMgr *RepositoriesManager) error {
	return repoMgr.PaymentTopupRepository.BeginTransaction()
}

func RollbackTransaction(repoManager *RepositoriesManager) error {
	return repoManager.PaymentTopupRepository.RollbackTransaction()
}

func CommitTransaction(repoManager *RepositoriesManager) error {
	return repoManager.PaymentTopupRepository.CommitTransaction()
}

func (repo *PaymentTopupRepository) BeginTransaction() error {
	ctx := context.Background()
	transaction, err := repo.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	repo.transaction = transaction
	return nil
}

func (repo *PaymentTopupRepository) RollbackTransaction() error {
	if repo.transaction != nil {
		err := repo.transaction.Rollback()
		repo.transaction = nil
		return err
	}
	return nil
}

func (repo *PaymentTopupRepository) CommitTransaction() error {
	if repo.transaction != nil {
		err := repo.transaction.Commit()
		repo.transaction = nil
		return err
	}
	return nil
}
