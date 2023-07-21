package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const listPaymentFintech = `-- name: ListPaymentFintech :many

SELECT fint_entity_id, fint_code, fint_name, fint_modified_date FROM payment.fintech ORDER BY fint_name
`

func (q *Queries) ListPaymentFintech(ctx context.Context) ([]models.PaymentFintech, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentFintech)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.PaymentFintech
	for rows.Next() {
		var i models.PaymentFintech
		if err := rows.Scan(
			&i.FintEntityID,
			&i.FintCode,
			&i.FintName,
			&i.FintModifiedDate,
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

const getPaymentFintech = `-- name: GetPaymentFintech :one


SELECT fint_entity_id, fint_code, fint_name, fint_modified_date FROM payment.fintech WHERE fint_code = $1
`

// payment.fintech
func (q *Queries) GetPaymentFintech(ctx context.Context, fintCode string) (models.PaymentFintech, error) {
	row := q.db.QueryRowContext(ctx, getPaymentFintech, fintCode)
	var i models.PaymentFintech
	err := row.Scan(
		&i.FintEntityID,
		&i.FintCode,
		&i.FintName,
		&i.FintModifiedDate,
	)
	return i, err
}

const createPaymentFintech = `-- name: CreatePaymentFintech :one

INSERT INTO
    payment.fintech (
        fint_entity_id,
        fint_code,
        fint_name,
        fint_modified_date
    )
VALUES ($1, $2, $3, $4) RETURNING *
`

type CreatePaymentFintechParams struct {
	FintEntityID     int32          `db:"fint_entity_id" json:"fintEntityId"`
	FintCode         sql.NullString `db:"fint_code" json:"fintCode"`
	FintName         sql.NullString `db:"fint_name" json:"fintName"`
	FintModifiedDate sql.NullTime   `db:"fint_modified_date" json:"fintModifiedDate"`
}

func (q *Queries) CreatePaymentFintech(ctx context.Context, arg CreatePaymentFintechParams) (*models.PaymentFintech, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentFintech,
		arg.FintEntityID,
		arg.FintCode,
		arg.FintName,
		arg.FintModifiedDate,
	)

	i := models.PaymentFintech{}
	err := row.Scan(
		&i.FintEntityID,
		&i.FintCode,
		&i.FintName,
		&i.FintModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.PaymentFintech{
		FintEntityID:     i.FintEntityID,
		FintCode:         i.FintCode,
		FintName:         i.FintName,
		FintModifiedDate: sql.NullTime{Time: time.Now(), Valid: true},
	}, nil
}

const updatePaymentFintech = `-- name: UpdatePaymentFintech :exec

UPDATE payment.fintech
set
    fint_code = $2,
    fint_name = $3
WHERE fint_entity_id = $1
`

func (q *Queries) UpdatePaymentFintech(ctx context.Context, arg CreatePaymentFintechParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentFintech, arg.FintEntityID, arg.FintCode, arg.FintName)
	return err
}

const deletePaymentFintech = `-- name: DeletePaymentFintech :exec

DELETE FROM payment.fintech WHERE fint_entity_id = $1
`

func (q *Queries) DeletePaymentFintech(ctx context.Context, fintEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentFintech, fintEntityID)
	return err
}
