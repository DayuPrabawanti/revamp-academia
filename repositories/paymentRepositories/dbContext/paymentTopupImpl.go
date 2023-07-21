package dbContext

import (
	"context"
	"database/sql"
	"errors"
)

type PaymentService struct {
	db *sql.DB
}

type TopupDetail struct {
	SourceName    string
	SourceAccount string
	SourceSaldo   float64
	TargetName    string
	TargetAccount string
	TargetSaldo   float64
}

const getTopupDetail = `-- name: GetTopupDetail :one

SELECT
			b.bank_code,
			b.bank_entity_id,
			bs.usac_saldo,
			f.fint_code,
			f.fint_entity_id,
			fs.usac_saldo
		FROM
			payment.bank b
		JOIN
			payment.users_account bs ON b.bank_entity_id = bs.usac_bank_entity_id
		JOIN
			payment.fintech f ON f.fint_entity_id = targetFintechEntityID
		JOIN
			payment.users_account fs ON f.fint_entity_id = fs.usac_user_entity_id
		WHERE
			b.bank_entity_id = $1 OR f.fint_entity_id = $2

`

func (q *Queries) GetTopupDetail(ctx context.Context, sourceBankEntityID int32, targetFintechEntityID int32) (*TopupDetail, error) {
	var detail TopupDetail

	err := q.db.QueryRowContext(ctx, getTopupDetail, sourceBankEntityID, targetFintechEntityID).Scan(
		&detail.SourceName,
		&detail.SourceAccount,
		&detail.SourceSaldo,
		&detail.TargetName,
		&detail.TargetAccount,
		&detail.TargetSaldo,
	)

	if err != nil {
		return nil, err
	}

	return &detail, nil
}

func (ps *PaymentService) Topup(ctx context.Context, sourceBankEntityID int32, targetFintechEntityID int32, amount float64) error {
	tx, err := ps.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// check source balance
	var sourceBalance float64
	err = tx.QueryRowContext(ctx, `SELECT usac_saldo FROM payment.users_account WHERE usac_bank_entity_id = ?`, sourceBankEntityID).Scan(&sourceBalance)
	if err != nil {
		return err
	}

	if sourceBalance < amount {
		return errors.New("insufficient funds")
	}

	// deduct amount from source account
	_, err = tx.ExecContext(ctx, `UPDATE payment.users_account SET usac_saldo = usac_saldo - ? WHERE usac_bank_entity_id = ?`, amount, sourceBankEntityID)
	if err != nil {
		return err
	}

	// add amount to target account
	_, err = tx.ExecContext(ctx, `UPDATE payment.users_account SET usac_saldo = usac_saldo + ? WHERE usac_bank_entity_id = ?`, amount, targetFintechEntityID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
