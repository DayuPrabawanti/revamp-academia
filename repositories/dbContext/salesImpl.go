package dbcontext

import (
	"context"

	"codeid.revampacademy/models"
)

const getcSpecial_offer = `-- name: GetcSpecial_offer :one
SELECT spof_id, spof_description, spof_discount, spof_type, spof_start_date, spof_end_date, spof_min_qty, spof_max_qty, spof_modified_date, spof_cate_id FROM sales.special_offer
WHERE spof_id = $1
`

func (q *Queries) GetcSpecial_offer(ctx context.Context, spofID int32) (models.SalesSpecialOffer, error) {
	row := q.db.QueryRowContext(ctx, getcSpecial_offer, spofID)
	var i models.SalesSpecialOffer
	err := row.Scan(
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
	)
	return i, err
}

const getcSpecial_offer_programs = `-- name: GetcSpecial_offer_programs :one
SELECT soco_id, soco_spof_id, soco_prog_entity_id, soco_status, soco_modified_date FROM sales.special_offer_programs
WHERE soco_id = $1
`

func (q *Queries) GetcSpecial_offer_programs(ctx context.Context, socoID int32) (models.SalesSpecialOfferProgram, error) {
	row := q.db.QueryRowContext(ctx, getcSpecial_offer_programs, socoID)
	var i models.SalesSpecialOfferProgram
	err := row.Scan(
		&i.SocoID,
		&i.SocoSpofID,
		&i.SocoProgEntityID,
		&i.SocoStatus,
		&i.SocoModifiedDate,
	)
	return i, err
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
