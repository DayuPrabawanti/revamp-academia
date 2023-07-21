package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

// 1a. fungsi utk ambil getlist
const listPaymentBank = `-- name: ListPaymentBank :many
SELECT bank_entity_id, bank_code, bank_name, bank_modified_date FROM payment.bank ORDER BY bank_name
`

func (q *Queries) ListPaymentBank(ctx context.Context) ([]models.PaymentBank, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentBank)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.PaymentBank
	for rows.Next() {
		var i models.PaymentBank
		if err := rows.Scan(
			&i.BankEntityID,
			&i.BankCode,
			&i.BankName,
			&i.BankModifiedDate,
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

// 1a. fungsi utk ambil get payment bank
const getPaymentBank = `-- name: GetPaymentBank :one
SELECT bank_entity_id, bank_code, bank_name, bank_modified_date FROM payment.bank
WHERE bank_code = $1`

func (q *Queries) GetPaymentBank(ctx context.Context, name string) (models.PaymentBank, error) {
	row := q.db.QueryRowContext(ctx, getPaymentBank, name)
	var i models.PaymentBank
	err := row.Scan(
		&i.BankEntityID,
		&i.BankCode,
		&i.BankName,
		&i.BankModifiedDate,
	)
	return i, err
}

// 1.b fungsi utk create paymentbank
const createPaymentBank = `-- name: CreatePaymentBank :one
INSERT INTO
    payment.bank(
        bank_entity_id,
        bank_code,
        bank_name,
        bank_modified_date
    )
VALUES ($1, $2, $3, $4) RETURNING *
`

type CreatePaymentBankParams struct {
	BankEntityID     int32        `db:"bank_entity_id" json:"bankEntityId"`
	BankCode         string       `db:"bank_code" json:"bankCode"`
	BankName         string       `db:"bank_name" json:"bankName"`
	BankModifiedDate sql.NullTime `db:"bank_modified_date" json:"bankModifiedDate"`
}

func (q *Queries) CreatePaymentBank(ctx context.Context, arg CreatePaymentBankParams) (*models.PaymentBank, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentBank,
		arg.BankEntityID,
		arg.BankCode,
		arg.BankName,
		arg.BankModifiedDate,
	)

	i := models.PaymentBank{}
	err := row.Scan(
		&i.BankEntityID,
		&i.BankCode,
		&i.BankName,
		&i.BankModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.PaymentBank{
		BankEntityID:     i.BankEntityID,
		BankCode:         i.BankCode,
		BankName:         i.BankName,
		BankModifiedDate: sql.NullTime{Time: time.Now(), Valid: true},
	}, nil
}

const updatePaymentBank = `-- name: UpdatePaymentBank :exec
UPDATE payment.bank
  set bank_code = $2,
  bank_name = $3
WHERE bank_entity_id = $1
`

func (q *Queries) UpdatePaymentBank(ctx context.Context, arg CreatePaymentBankParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentBank, arg.BankEntityID, arg.BankCode, arg.BankName)
	return err
}

const deletePaymentBank = `-- name: DeletePaymentBank :exec
DELETE FROM payment.bank
WHERE bank_entity_id = $1
`

func (q *Queries) DeletePaymentBank(ctx context.Context, bankEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentBank, bankEntityID)
	return err
}
