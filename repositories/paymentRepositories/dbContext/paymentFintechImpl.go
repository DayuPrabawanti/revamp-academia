package dbContext

import (
	"context"
	"net/http"

	"codeid.revampacademy/models"
)

type Fintech struct {
	FintEntityID int    `db:"fint_entity_id"`
	FintCode     string `db:"fint_code"`
}

const listPaymentFintech = `-- name: ListPaymentFintech :many
SELECT 
	fint_entity_id, 
	fint_code 
FROM 
	payment.fintech
ORDER BY 
	fint_entity_id;
`

func (q *Queries) ListPaymentFintech(ctx context.Context) ([]Fintech, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentFintech)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Fintech
	for rows.Next() {
		var i Fintech
		if err := rows.Scan(
			&i.FintEntityID,
			&i.FintCode,
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
SELECT 
	fint_entity_id, 
	fint_code 
FROM 
	payment.fintech
WHERE
	fint_code = $1;
`

// payment.fintech
func (q *Queries) GetPaymentFintech(ctx context.Context, fintCode string) (Fintech, error) {
	row := q.db.QueryRowContext(ctx, getPaymentFintech, fintCode)
	var i Fintech
	err := row.Scan(
		&i.FintEntityID,
		&i.FintCode,
	)
	return i, err
}

const createPaymentFintech = `-- name: CreatePaymentFintech :one
INSERT INTO
    payment.fintech(
        fint_code
    )
VALUES ($1)
RETURNING 
	fint_entity_id,
	fint_code
`

type CreatePaymentFintechParams struct {
	FintEntityID int32  `db:"fint_entity_id" json:"fintEntityId"`
	FintCode     string `db:"fint_code" json:"fintCode"`
}

func (q *Queries) CreatePaymentFintech(ctx context.Context, arg CreatePaymentFintechParams) (*Fintech, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentFintech,
		arg.FintCode,
	)

	i := Fintech{}
	err := row.Scan(
		&i.FintEntityID,
		&i.FintCode,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &Fintech{
		FintEntityID: i.FintEntityID,
		FintCode:     i.FintCode,
	}, nil
}

const updatePaymentFintech = `-- name: UpdatePaymentFintech :exec
UPDATE 
	payment.fintech
SET
    fint_code = $2
WHERE
	fint_entity_id = $1
`

func (q *Queries) UpdatePaymentFintech(ctx context.Context, arg CreatePaymentFintechParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentFintech, arg.FintEntityID, arg.FintCode)
	return err
}

const deletePaymentFintech = `-- name: DeletePaymentFintech :exec
DELETE FROM 
	payment.fintech
WHERE 
	fint_entity_id = $1
`

func (q *Queries) DeletePaymentFintech(ctx context.Context, fintEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentFintech, fintEntityID)
	return err
}
