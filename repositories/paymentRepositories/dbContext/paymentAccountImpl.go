package dbContext

import (
	"context"
	"net/http"

	"codeid.revampacademy/models"
)

type UserAccount struct {
	AccountNumber string  `json:"account_number"`
	Description   string  `json:"description"`
	Saldo         float64 `json:"saldo"`
	Type          string  `json:"type"`
}

const listPaymentUsers_account = `-- name: ListPaymentUsers_account :many
SELECT 
    ua.usac_account_number, 
    COALESCE(b.bank_code, f.fint_code) AS description,
    ua.usac_saldo,
    ua.usac_type
FROM 
    payment.users_account ua
LEFT JOIN 
    payment.bank b ON ua.usac_bank_entity_id = b.bank_entity_id
LEFT JOIN 
    payment.fintech f ON ua.usac_bank_entity_id = f.fint_entity_id;
`

func (q *Queries) ListPaymentUsers_account(ctx context.Context) ([]UserAccount, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentUsers_account)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAccount
	for rows.Next() {
		var i UserAccount
		if err := rows.Scan(
			&i.AccountNumber,
			&i.Description,
			&i.Saldo,
			&i.Type,
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
SELECT 
    ua.usac_account_number, 
    COALESCE(b.bank_code, f.fint_code) AS description,
    ua.usac_saldo,
    ua.usac_type
FROM 
    payment.users_account ua
LEFT JOIN 
    payment.bank b ON ua.usac_bank_entity_id = b.bank_entity_id
LEFT JOIN 
    payment.fintech f ON ua.usac_bank_entity_id = f.fint_entity_id
WHERE 
	usac_account_number = $1;
`

// payment.users_account
func (q *Queries) GetPaymentUsers_account(ctx context.Context, usacAccountNumber string) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, getPaymentUsers_account, usacAccountNumber)
	var i UserAccount
	err := row.Scan(
		&i.AccountNumber,
		&i.Description,
		&i.Saldo,
		&i.Type,
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
		usac_status
	)
VALUES 
	($1, $2, $3, $4, $5, 'active')
RETURNING *
`

type CreatePaymentUsers_accountParams struct {
	UsacBankEntityID  int32   `db:"usac_bank_entity_id" json:"usacBankEntityID"`
	UsacUserEntityID  int32   `db:"usac_user_entity_id" json:"usacUserEntityID"`
	UsacAccountNumber string  `db:"usac_account_number" json:"usacAccountNumber"`
	UsacSaldo         float64 `db:"usac_saldo" json:"usacSaldo"`
	UsacType          string  `db:"usac_type" json:"usacType"`
	UsacStatus        string  `db:"usac_status" json:"usacStatus"`
}

func (q *Queries) CreatePaymentUsers_account(ctx context.Context, arg CreatePaymentUsers_accountParams) (*models.PaymentUsersAccount, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentUsers_account,
		arg.UsacBankEntityID,
		arg.UsacUserEntityID,
		arg.UsacAccountNumber,
		arg.UsacSaldo,
		arg.UsacType,
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
		UsacStartDate:     i.UsacStartDate,
		UsacEndDate:       i.UsacEndDate,
		UsacModifiedDate:  i.UsacModifiedDate,
		UsacStatus:        i.UsacStatus,
	}, nil
}

const updatePaymentUsers_account = `-- name: UpdatePaymentUsers_account :exec
UPDATE 
	payment.users_account
SET
    usac_saldo = $2,
    usac_type = $3
WHERE usac_account_number = $1
RETURNING *;
`

func (q *Queries) UpdatePaymentUsers_account(ctx context.Context, arg CreatePaymentUsers_accountParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentUsers_account,
		arg.UsacAccountNumber,
		arg.UsacSaldo,
		arg.UsacType,
	)
	return err
}

const deletePaymentUsers_account = `-- name: DeletePaymentUsers_account :exec
DELETE FROM 
	payment.users_account 
WHERE 
	usac_account_number = $1
`

func (q *Queries) DeletePaymentUsers_account(ctx context.Context, accountNumber string) error {
	_, err := q.db.ExecContext(ctx, deletePaymentUsers_account, accountNumber)
	return err
}
