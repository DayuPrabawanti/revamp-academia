package salesContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const createSales_special_offer = `-- name: CreateSales_special_offer :one
INSERT INTO sales.special_offer(spof_id, spof_description, spof_discount, spof_type,spof_start_date,spof_end_date,spof_min_qty,spof_max_qty,spof_modified_date,spof_cate_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *
`

type CreateSales_special_offerParams struct {
	SpofID           int32          `db:"spof_id" json:"spofId"`
	SpofDescription  string         `db:"spof_description" json:"spofDescription"`
	SpofDiscount     int32          `db:"spof_discount" json:"spofDiscount"`
	SpofType         sql.NullString `db:"spof_type" json:"spofType"`
	SpofStartDate    time.Time      `db:"spof_start_date" json:"spofStartDate"`
	SpofEndDate      time.Time      `db:"spof_end_date" json:"spofEndDate"`
	SpofMinQty       int32          `db:"spof_min_qty" json:"spofMinQty"`
	SpofMaxQty       int32          `db:"spof_max_qty" json:"spofMaxQty"`
	SpofModifiedDate sql.NullString `db:"spof_modified_date" json:"spofModifiedDate"`
	SpofCateID       int32          `db:"spof_cate_id" json:"spofCateId"`
}

func (q *Queries) CreateSales_special_offer(ctx context.Context, arg CreateSales_special_offerParams) (*models.SalesSpecialOffer, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createSales_special_offer,
		arg.SpofID,
		arg.SpofDescription,
		arg.SpofDiscount,
		arg.SpofType,
		arg.SpofStartDate,
		arg.SpofEndDate,
		arg.SpofMinQty,
		arg.SpofMaxQty,
		arg.SpofModifiedDate,
		arg.SpofCateID,
	)
	i := models.SalesSpecialOffer{}
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

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.SalesSpecialOffer{
		SpofID:           i.SpofID,
		SpofDescription:  i.SpofDescription,
		SpofDiscount:     i.SpofDiscount,
		SpofType:         i.SpofType,
		SpofStartDate:    i.SpofStartDate,
		SpofEndDate:      i.SpofEndDate,
		SpofMinQty:       i.SpofMinQty,
		SpofMaxQty:       i.SpofMaxQty,
		SpofModifiedDate: i.SpofModifiedDate,
		SpofCateID:       i.SpofCateID,
	}, nil
}

const getcSpecial_offer = `-- name: GetcSpecial_offer :one
SELECT spof_id, spof_description, spof_discount, spof_type, spof_start_date, spof_end_date, spof_min_qty, spof_max_qty, spof_modified_date, spof_cate_id FROM sales.special_offer
WHERE spof_id = $1
`

func (q *Queries) GetSpecial_offer(ctx context.Context, spofID int32) (models.SalesSpecialOffer, error) {
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
