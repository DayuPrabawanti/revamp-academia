package dbcontext

import (
	"context"
	"database/sql"

	"codeid.revampacademy/models"
)

type CreateSales_special_offerParams struct {
	SpofID           int32          `db:"spof_id" json:"spofId"`
	SpofDescription  string         `db:"spof_description" json:"spofDescription"`
	SpofDiscount     int32          `db:"spof_discount" json:"spofDiscount"`
	SpofType         sql.NullInt32  `db:"spof_type" json:"spofType"`
	SpofStartDate    string         `db:"spof_start_date" json:"spofStartDate"`
	SpofEndDate      string         `db:"spof_end_date" json:"spofEndDate"`
	SpofMinQty       int32          `db:"spof_min_qty" json:"spofMinQty"`
	SpofMaxQty       int32          `db:"spof_max_qty" json:"spofMaxQty"`
	SpofModifiedDate sql.NullString `db:"spof_modified_date" json:"spofModifiedDate"`
	SpofCateID       int32          `db:"spof_cate_id" json:"spofCateId"`
}

type CreateCartItemsParams struct {
	CaitID           int32          `db:"cait_id" json:"caitId"`
	CaitQuantity     int32          `db:"cait_quantity" json:"caitQuantity"`
	CaitUnitPrice    sql.NullString `db:"cait_unit_price" json:"caitUnitPrice"`
	CaitModifiedDate sql.NullString `db:"cait_modified_date" json:"caitModifiedDate"`
	CaitUserEntityID int32          `db:"cait_user_entity_id" json:"caitUserEntityId"`
	CaitProgEntityID int32          `db:"cait_prog_entity_id" json:"caitProgEntityId"`
}

const listSpecial_offer = `-- name: ListSpecial_offer :many
SELECT spof_id, spof_description, spof_discount, spof_type, spof_start_date, spof_end_date, spof_min_qty, spof_max_qty, spof_modified_date, spof_cate_id FROM sales.special_offer
ORDER BY  spof_description
`

func (q *Queries) ListSpecial_offer(ctx context.Context) ([]models.SalesSpecialOffer, error) {
	rows, err := q.db.QueryContext(ctx, listSpecial_offer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.SalesSpecialOffer
	for rows.Next() {
		var i models.SalesSpecialOffer
		if err := rows.Scan(
			&i.SpofID,
			&i.SpofDescription,
			&i.SpofDiscount,
			&i.SpofType,
			&i.SpofStartDate,
			&i.SpofEndDate,
			&i.SpofMinQty,
			&i.SpofMaxQty,
			&i.SpofModifiedDate,
			&i.SpofCateID,
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
