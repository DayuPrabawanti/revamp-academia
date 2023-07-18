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
