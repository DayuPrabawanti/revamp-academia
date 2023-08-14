package jobhireRepositories

import (
	"context"
	"database/sql"
)

func BeginTransaction(repoMgr *RepositoryManager) error {
	ctx := context.Background()
	transaction, err := repoMgr.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	repoMgr.transaction = transaction

	return nil
}

func RollbackTransaction(repoManager *RepositoryManager) error {
	transaction := repoManager.transaction

	repoManager.transaction = nil

	return transaction.Rollback()
}

func CommitTransaction(repoManager *RepositoryManager) error {
	transaction := repoManager.transaction

	repoManager.transaction = nil

	return transaction.Commit()
}
