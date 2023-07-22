package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const getcart_items = `-- name: Getcart_items :one
SELECT cait_id, cait_quantity, cait_unit_price, cait_modified_date, cait_user_entity_id, cait_prog_entity_id FROM sales.cart_items
WHERE cait_id = $1
`

func (q *Queries) Getcart_items(ctx context.Context, caitID int32) (models.SalesCartItem, error) {
	row := q.db.QueryRowContext(ctx, getcart_items, caitID)
	var i models.SalesCartItem
	err := row.Scan(
		&i.CaitID,
		&i.CaitQuantity,
		&i.CaitUnitPrice,
		&i.CaitModifiedDate,
		&i.CaitUserEntityID,
		&i.CaitProgEntityID,
	)
	return i, err

}

const listCart_item = `-- name: ListCart_item :many
SELECT cait_id, cait_quantity, cait_unit_price, cait_modified_date, cait_user_entity_id, cait_prog_entity_id FROM sales.cart_items
ORDER BY cait_quantity
`

func (q *Queries) ListCart_item(ctx context.Context) ([]models.SalesCartItem, error) {
	rows, err := q.db.QueryContext(ctx, listCart_item)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.SalesCartItem
	for rows.Next() {
		var i models.SalesCartItem
		if err := rows.Scan(
			&i.CaitID,
			&i.CaitQuantity,
			&i.CaitUnitPrice,
			&i.CaitModifiedDate,
			&i.CaitUserEntityID,
			&i.CaitProgEntityID,
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
