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
	SourceName    sql.NullString
	SourceAccount sql.NullString
	SourceSaldo   float64
	TargetName    sql.NullString
	TargetAccount sql.NullString
	TargetSaldo   float64
}

const listTopupDetail = `-- name: ListTopupDetail :many

SELECT
			b.bank_code,
			b.bank_entity_id,
			fs.usac_saldo,
			f.fint_code,
			f.fint_entity_id,
			fs.usac_saldo
		FROM
			payment.users_account fs
		LEFT JOIN
			payment.bank b ON fs.usac_bank_entity_id = b.bank_entity_id
		LEFT JOIN
			payment.fintech f ON fs.usac_bank_entity_id = f.fint_entity_id

`

func (q *Queries) ListTopupDetail(ctx context.Context) ([]TopupDetail, error) {
	rows, err := q.db.QueryContext(ctx, listTopupDetail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TopupDetail
	for rows.Next() {
		var i TopupDetail
		if err := rows.Scan(
			&i.SourceName,
			&i.SourceAccount,
			&i.SourceSaldo,
			&i.TargetName,
			&i.TargetAccount,
			&i.TargetSaldo,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTopupDetailById = `-- name: GetTopupDetailById :many

SELECT
			b.bank_code,
			b.bank_entity_id,
			fs.usac_saldo,
			f.fint_code,
			f.fint_entity_id,
			fs.usac_saldo
		FROM
			payment.users_account fs
		LEFT JOIN
			payment.bank b ON fs.usac_bank_entity_id = b.bank_entity_id
		LEFT JOIN
			payment.fintech f ON fs.usac_bank_entity_id = f.fint_entity_id
		WHERE
			usac_user_entity_id = $1

`

func (q *Queries) GetTopupDetailById(ctx context.Context, id int32) ([]TopupDetail, error) {
	rows, err := q.db.QueryContext(ctx, getTopupDetailById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TopupDetail
	for rows.Next() {
		var i TopupDetail
		if err := rows.Scan(
			&i.SourceName,
			&i.SourceAccount,
			&i.SourceSaldo,
			&i.TargetName,
			&i.TargetAccount,
			&i.TargetSaldo,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
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
