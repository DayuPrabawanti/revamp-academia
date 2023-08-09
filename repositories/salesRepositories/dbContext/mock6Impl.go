package dbcontext

import (
	"context"

	"codeid.revampacademy/models"
)

const getUsersMock6 = `-- name: GetUsers :one

select u.user_entity_id, u.user_name, p.prog_title, p.prog_headline, p.prog_image, p.prog_price,
s.sode_unit_discount, s.sode_line_total, f.fint_code, f.fint_name
from users.users u 
join curriculum.program_entity p on u.user_entity_id = p.prog_entity_id
join sales.sales_order_detail s on p.prog_entity_id = s.sode_prog_entity_id
join payment.fintech f on u.user_entity_id = f.fint_entity_id
WHERE user_entity_id = $1
`

func (q *Queries) GetUsersIdMock6(ctx context.Context, userEntityID int32) (models.MergeShopMock7, error) {
	row := q.db.QueryRowContext(ctx, getUsersMock6, userEntityID)
	var i models.MergeShopMock7
	err := row.Scan(
		&i.UserEntityID,
		&i.UserName,
		&i.ProgTitle,
		&i.ProgHeadline,
		&i.ProgImage,
		&i.ProgPrice,
		&i.SodeUnitDiscount,
		&i.SodeLineTotal,
		&i.FintCode,
		&i.FintName,
	)
	return i, err
}

const listSales_order_detail = `-- name: ListSales_order_detail :many
SELECT sode_id, sode_qty, sode_unit_price, sode_unit_discount, sode_line_total, sode_modified_date, sode_sohe_id, sode_prog_entity_id FROM sales.sales_order_detail
ORDER BY sode_qty
`

func (q *Queries) ListSalesOrderMock6(ctx context.Context) ([]models.SalesSalesOrderDetail, error) {
	rows, err := q.db.QueryContext(ctx, listSales_order_detail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.SalesSalesOrderDetail
	for rows.Next() {
		var i models.SalesSalesOrderDetail
		if err := rows.Scan(
			&i.SodeID,
			&i.SodeQty,
			&i.SodeUnitPrice,
			&i.SodeUnitDiscount,
			&i.SodeLineTotal,
			&i.SodeModifiedDate,
			&i.SodeSoheID,
			&i.SodeProgEntityID,
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

const getPaymentOrder = `-- name: Getpaymentorder :one

select ua.usac_account_number,ua.usac_type,f.fint_code,f.fint_name from payment.users_account ua join payment.fintech f 
on ua.usac_user_entity_id = f.fint_entity_id 
where ua.usac_account_number=$1
`

func (q *Queries) GetAccountNumbers(ctx context.Context, account string) (models.MergePayment, error) {
	row := q.db.QueryRowContext(ctx, getPaymentOrder, account)
	var i models.MergePayment
	err := row.Scan(
		&i.UsacAccountNumber,
		&i.UsacType,
		&i.FintCode,
		&i.FintName,
	)
	return i, err
}
