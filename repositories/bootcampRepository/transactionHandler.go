package bootcampRepository

import (
	"context"
	"database/sql"
)

func BeginTransaction(repoMgr *RepositoryManager) error {
	return repoMgr.BootcampBatchEvaluationRepository.BeginTransaction()
}

func RollbackTransaction(repoManager *RepositoryManager) error {
	return repoManager.BootcampBatchEvaluationRepository.RollbackTransaction()
}

func CommitTransaction(repoManager *RepositoryManager) error {
	return repoManager.BootcampBatchEvaluationRepository.CommitTransaction()
}

func (repo *BootcampBatchEvaluationRepository) BeginTransaction() error {
	ctx := context.Background()
	transaction, err := repo.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	repo.transaction = transaction
	return nil
}

func (repo *BootcampBatchEvaluationRepository) RollbackTransaction() error {
	if repo.transaction != nil {
		err := repo.transaction.Rollback()
		repo.transaction = nil
		return err
	}
	return nil
}

func (repo *BootcampBatchEvaluationRepository) CommitTransaction() error {
	if repo.transaction != nil {
		err := repo.transaction.Commit()
		repo.transaction = nil
		return err
	}
	return nil
}
