package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

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
