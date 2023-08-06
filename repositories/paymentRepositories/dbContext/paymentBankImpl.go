package dbContext

import (
	"context"
	"net/http"

	"codeid.revampacademy/models"
)

type Bank struct {
	BankEntityID int    `db:"bank_entity_id"`
	BankCode     string `db:"bank_code"`
}

// 1a. fungsi utk ambil getlist
const listPaymentBank = `-- name: ListPaymentBank :many
SELECT 
	bank_entity_id, 
	bank_code
FROM 
	payment.bank 
ORDER BY 
	bank_entity_id;
`

func (q *Queries) ListPaymentBank(ctx context.Context) ([]Bank, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentBank)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bank
	for rows.Next() {
		var i Bank
		if err := rows.Scan(
			&i.BankEntityID,
			&i.BankCode,
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
SELECT 
	bank_entity_id, 
	bank_code
FROM 
	payment.bank
WHERE 
	bank_code = $1;
`

func (q *Queries) GetPaymentBank(ctx context.Context, name string) (Bank, error) {
	row := q.db.QueryRowContext(ctx, getPaymentBank, name)
	var i Bank
	err := row.Scan(
		&i.BankEntityID,
		&i.BankCode,
	)
	return i, err
}

// 1.b fungsi utk create paymentbank
const createPaymentBank = `-- name: CreatePaymentBank :one
INSERT INTO
    payment.bank(
        bank_code
    )
VALUES ($1) RETURNING bank_entity_id, bank_code
`

type CreatePaymentBankParams struct {
	BankEntityID int32  `db:"bank_entity_id" json:"bankEntityId"`
	BankCode     string `db:"bank_code" json:"bankCode"`
}

func (q *Queries) CreatePaymentBank(ctx context.Context, arg CreatePaymentBankParams) (*Bank, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentBank,
		arg.BankCode,
	)

	i := Bank{}
	err := row.Scan(
		&i.BankEntityID,
		&i.BankCode,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &Bank{
		BankEntityID: i.BankEntityID,
		BankCode:     i.BankCode,
	}, nil
}

const updatePaymentBank = `-- name: UpdatePaymentBank :exec
UPDATE 
	payment.bank
SET
	bank_code = $1
WHERE 
	bank_entity_id = $2
`

func (q *Queries) UpdatePaymentBank(ctx context.Context, arg CreatePaymentBankParams, bankEntityID int32) error {
	_, err := q.db.ExecContext(ctx, updatePaymentBank, arg.BankCode, bankEntityID)
	return err
}

const deletePaymentBank = `-- name: DeletePaymentBank :exec
DELETE FROM 
	payment.bank
WHERE 
	bank_entity_id = $1
`

func (q *Queries) DeletePaymentBank(ctx context.Context, bankEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentBank, bankEntityID)
	return err
}
