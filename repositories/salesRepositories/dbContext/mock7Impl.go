package dbcontext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

type PaymentParams struct {
	UsacAccountNumber sql.NullString `db:"usac_account_number" json:"usacAccountNumber"`
	UserName          sql.NullString `db:"user_name" json:"userName"`
	TrpaNote          sql.NullString `db:"trpa_note" json:"trpaNote"`
	TrpaCredit        sql.NullString `db:"trpa_credit" json:"trpaCredit"`
}

type UserEntityParamsMork7 struct {
	UserEntityID  int32          `db:"user_entity_id" json:"userEntityId"`
	UserName      sql.NullString `db:"user_name" json:"userName"`
	ProgTitle     sql.NullString `db:"prog_title" json:"progTitle"`
	ProgHeadline  sql.NullString `db:"prog_headline" json:"progHeadline"`
	ProgImage     sql.NullString `db:"prog_image" json:"progImage"`
	ProgPrice     sql.NullInt32  `db:"prog_price" json:"progPrice"`
	SodeLineTotal sql.NullString `db:"sode_line_total" json:"sodeLineTotal"`
}

const getUsersMock7 = `-- name: GetUsers :one

select u.user_entity_id, u.user_name, p.prog_title, p.prog_headline, p.prog_image, p.prog_price,
s.sode_line_total
from users.users u 
join curriculum.program_entity p on u.user_entity_id = p.prog_entity_id
join sales.sales_order_detail s on p.prog_entity_id = s.sode_prog_entity_id
WHERE user_entity_id = $1
`

func (q *Queries) GetUsersIdMock7(ctx context.Context, userEntityID int32) (UserEntityParamsMork7, error) {
	row := q.db.QueryRowContext(ctx, getUsersMock7, userEntityID)
	var i UserEntityParamsMork7
	err := row.Scan(
		&i.UserEntityID,
		&i.UserName,
		&i.ProgTitle,
		&i.ProgHeadline,
		&i.ProgImage,
		&i.ProgPrice,
		&i.SodeLineTotal,
	)
	return i, err
}

const createSales_order_detail = `-- name: CreateSales_order_detail :one
INSERT INTO sales.sales_order_detail(sode_id, sode_qty, sode_unit_price, sode_unit_discount,sode_line_total,sode_modified_date,sode_sohe_id,sode_prog_entity_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *
`

type CreateSales_order_detailParams struct {
	SodeID           int32          `db:"sode_id" json:"sodeId"`
	SodeQty          sql.NullInt32  `db:"sode_qty" json:"sodeQty"`
	SodeUnitPrice    sql.NullString `db:"sode_unit_price" json:"sodeUnitPrice"`
	SodeUnitDiscount sql.NullString `db:"sode_unit_discount" json:"sodeUnitDiscount"`
	SodeLineTotal    sql.NullString `db:"sode_line_total" json:"sodeLineTotal"`
	SodeModifiedDate sql.NullTime   `db:"sode_modified_date" json:"sodeModifiedDate"`
	SodeSoheID       sql.NullInt32  `db:"sode_sohe_id" json:"sodeSoheId"`
	SodeProgEntityID sql.NullInt32  `db:"sode_prog_entity_id" json:"sodeProgEntityId"`
}

func (q *Queries) CreateSales_order_detail(ctx context.Context, arg CreateSales_order_detailParams) (*models.SalesSalesOrderDetail, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createSales_order_detail,
		arg.SodeID,
		arg.SodeQty,
		arg.SodeUnitPrice,
		arg.SodeUnitDiscount,
		arg.SodeLineTotal,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.SodeSoheID,
		arg.SodeProgEntityID,
	)
	i := models.SalesSalesOrderDetail{}
	err := row.Scan(
		&i.SodeID,
		&i.SodeQty,
		&i.SodeUnitPrice,
		&i.SodeUnitDiscount,
		&i.SodeLineTotal,
		&i.SodeModifiedDate,
		&i.SodeSoheID,
		&i.SodeProgEntityID,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &models.SalesSalesOrderDetail{
		SodeID:           i.SodeID,
		SodeQty:          i.SodeQty,
		SodeUnitPrice:    i.SodeUnitPrice,
		SodeUnitDiscount: i.SodeUnitDiscount,
		SodeLineTotal:    i.SodeLineTotal,
		SodeModifiedDate: i.SodeModifiedDate,
		SodeSoheID:       i.SodeSoheID,
		SodeProgEntityID: i.SodeProgEntityID,
	}, nil
}

const deleteSales_order_detail = `-- name: DeleteSales_order_detail :exec
DELETE FROM sales.sales_order_detail
WHERE sode_id = $1
`

func (q *Queries) DeleteSales_order_detail(ctx context.Context, sodeID int16) error {
	_, err := q.db.ExecContext(ctx, deleteSales_order_detail, sodeID)
	return err
}

const getAccountNumberMock7 = `-- name: GetAccountNumberMock7 :exec
select ua.usac_account_number,us.user_name,tp.trpa_note,tp.trpa_credit
from payment.users_account ua join users.users us on ua.usac_user_entity_id = us.user_entity_id
join payment.transaction_payment tp on ua.usac_user_entity_id = tp.trpa_user_entity_id
WHERE ua.usac_account_number = $1
`

func (q *Queries) GetAccountNumberMock7(ctx context.Context, account string) (PaymentParams, error) {
	row := q.db.QueryRowContext(ctx, getAccountNumberMock7, account)
	var i PaymentParams
	err := row.Scan(
		&i.UsacAccountNumber,
		&i.UserName,
		&i.TrpaNote,
		&i.TrpaCredit,
	)
	return i, err
}
