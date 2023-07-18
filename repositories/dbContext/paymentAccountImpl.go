package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const listPaymentUsers_account = `-- name: ListPaymentUsers_account :many

SELECT usac_bank_entity_id, usac_user_entity_id, usac_account_number, usac_saldo, usac_type, usac_start_date, usac_end_date, usac_modified_date, usac_status FROM payment.users_account ORDER BY usac_account_number
`

func (q *Queries) ListPaymentUsers_account(ctx context.Context) ([]models.PaymentUsersAccount, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentUsers_account)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.PaymentUsersAccount
	for rows.Next() {
		var i models.PaymentUsersAccount
		if err := rows.Scan(
			&i.UsacBankEntityID,
			&i.UsacUserEntityID,
			&i.UsacAccountNumber,
			&i.UsacSaldo,
			&i.UsacType,
			&i.UsacStartDate,
			&i.UsacEndDate,
			&i.UsacModifiedDate,
			&i.UsacStatus,
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

const getPaymentUsers_account = `-- name: GetPaymentUsers_account :one


SELECT usac_bank_entity_id, usac_user_entity_id, usac_account_number, usac_saldo, usac_type, usac_start_date, usac_end_date, usac_modified_date, usac_status FROM payment.users_account WHERE usac_account_number = $1
`

// payment.users_account
func (q *Queries) GetPaymentUsers_account(ctx context.Context, usacAccountNumber string) (models.PaymentUsersAccount, error) {
	row := q.db.QueryRowContext(ctx, getPaymentUsers_account, usacAccountNumber)
	var i models.PaymentUsersAccount
	err := row.Scan(
		&i.UsacBankEntityID,
		&i.UsacUserEntityID,
		&i.UsacAccountNumber,
		&i.UsacSaldo,
		&i.UsacType,
		&i.UsacStartDate,
		&i.UsacEndDate,
		&i.UsacModifiedDate,
		&i.UsacStatus,
	)
	return i, err
}

const createPaymentUsers_account = `-- name: CreatePaymentUsers_account :one

INSERT INTO
    payment.users_account (
        usac_bank_entity_id,
        usac_user_entity_id,
        usac_account_number,
        usac_saldo,
        usac_type,
        usac_start_date,
        usac_end_date,
        usac_modified_date,
        usac_status
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *
`

type CreatePaymentUsers_accountParams struct {
	UsacBankEntityID  int32          `db:"usac_bank_entity_id" json:"usacBankEntityId"`
	UsacUserEntityID  int32          `db:"usac_user_entity_id" json:"usacUserEntityId"`
	UsacAccountNumber sql.NullString `db:"usac_account_number" json:"usacAccountNumber"`
	UsacSaldo         sql.NullString `db:"usac_saldo" json:"usacSaldo"`
	UsacType          sql.NullString `db:"usac_type" json:"usacType"`
	UsacStartDate     sql.NullTime   `db:"usac_start_date" json:"usacStartDate"`
	UsacEndDate       sql.NullTime   `db:"usac_end_date" json:"usacEndDate"`
	UsacModifiedDate  sql.NullTime   `db:"usac_modified_date" json:"usacModifiedDate"`
	UsacStatus        sql.NullString `db:"usac_status" json:"usacStatus"`
}

func (q *Queries) CreatePaymentUsers_account(ctx context.Context, arg CreatePaymentUsers_accountParams) (*models.PaymentUsersAccount, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentUsers_account,
		arg.UsacBankEntityID,
		arg.UsacUserEntityID,
		arg.UsacAccountNumber,
		arg.UsacSaldo,
		arg.UsacType,
		arg.UsacStartDate,
		arg.UsacEndDate,
		arg.UsacModifiedDate,
		arg.UsacStatus,
	)

	i := models.PaymentUsersAccount{}
	err := row.Scan(
		&i.UsacBankEntityID,
		&i.UsacUserEntityID,
		&i.UsacAccountNumber,
		&i.UsacSaldo,
		&i.UsacType,
		&i.UsacStartDate,
		&i.UsacEndDate,
		&i.UsacModifiedDate,
		&i.UsacStatus,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &models.PaymentUsersAccount{
		UsacBankEntityID:  i.UsacBankEntityID,
		UsacUserEntityID:  i.UsacUserEntityID,
		UsacAccountNumber: i.UsacAccountNumber,
		UsacSaldo:         i.UsacSaldo,
		UsacType:          i.UsacType,
		UsacStartDate:     sql.NullTime{Time: time.Now(), Valid: true},
		UsacEndDate:       sql.NullTime{Time: time.Now(), Valid: true},
		UsacModifiedDate:  sql.NullTime{Time: time.Now(), Valid: true},
		UsacStatus:        i.UsacStatus,
	}, nil
}

const updatePaymentUsers_account = `-- name: UpdatePaymentUsers_account :exec

UPDATE payment.users_account
set
    usac_user_entity_id = $2,
    usac_account_number = $3,
    usac_saldo = $4,
    usac_type = $5,
    usac_start_date = $6,
    usac_end_date = $7,
    usac_modified_date = $8,
    usac_status = $9
WHERE usac_bank_entity_id = $1, 
`

func (q *Queries) UpdatePaymentUsers_account(ctx context.Context, arg CreatePaymentUsers_accountParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentUsers_account,
		arg.UsacBankEntityID,
		arg.UsacUserEntityID,
		arg.UsacAccountNumber,
		arg.UsacSaldo,
		arg.UsacType,
		arg.UsacStartDate,
		arg.UsacEndDate,
		arg.UsacModifiedDate,
		arg.UsacStatus,
	)
	return err
}

const deletePaymentUsers_account = `-- name: DeletePaymentUsers_account :exec

DELETE FROM payment.users_account WHERE usac_bank_entity_id = $1
`

func (q *Queries) DeletePaymentUsers_account(ctx context.Context, usacBankEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentUsers_account, usacBankEntityID)
	return err
}
