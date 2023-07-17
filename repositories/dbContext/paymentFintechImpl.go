package dbContext

import (
	"context"

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

SELECT fint_entity_id, fint_code, fint_name, fint_modified_date FROM payment.fintech WHERE fint_entity_id = $1
`

// payment.fintech
func (q *Queries) GetPaymentFintech(ctx context.Context, fintEntityID int32) (models.PaymentFintech, error) {
	row := q.db.QueryRowContext(ctx, getPaymentFintech, fintEntityID)
	var i models.PaymentFintech
	err := row.Scan(
		&i.FintEntityID,
		&i.FintCode,
		&i.FintName,
		&i.FintModifiedDate,
	)
	return i, err
}
