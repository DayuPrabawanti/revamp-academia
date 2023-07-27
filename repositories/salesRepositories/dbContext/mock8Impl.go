package dbcontext

import (
	"context"
	"database/sql"
)

type GetSummaryOrderMock8 struct {
	ProgTitle         sql.NullString `db:"prog_title" json:"progTitle"`
	ProgHeadline      sql.NullString `db:"prog_headline" json:"progHeadline"`
	ProgImage         sql.NullString `db:"prog_image" json:"progImage"`
	ProgPrice         sql.NullInt32  `db:"prog_price" json:"progPrice"`
	SoheOrderNumber   sql.NullString `db:"sohe_order_number" json:"soheOrderNumber"`
	SoheSubtotal      sql.NullString `db:"sohe_subtotal" json:"soheSubtotal"`
	TrpaNote          sql.NullString `db:"trpa_note" json:"trpaNote"`
	SoheAccountNumber sql.NullString `db:"sohe_account_number" json:"soheAccountNumber"`
	UserName          sql.NullString `db:"user_name" json:"userName"`
	TrpaCredit        sql.NullString `db:"trpa_credit" json:"trpaCredit"`
	TrpaCodeNumber    sql.NullString `db:"trpa_code_number" json:"trpaCodeNumber"`
}

const getCartMockup8 = `-- name: GetSummaryorderMock8 :one
select so.sohe_order_number,p.prog_title,p.prog_headline,p.prog_image,p.prog_price
,so.sohe_subtotal,tp.trpa_note,so.sohe_account_number,s.user_name,tp.trpa_credit
,tp.trpa_code_number from sales.sales_order_header so
join users.users s on s.user_entity_id = so.sohe_user_entity_id
join curriculum.program_entity p on p.prog_entity_id = so.sohe_user_entity_id
join payment.transaction_payment tp on tp.trpa_user_entity_id= so.sohe_user_entity_id
WHERE so.sohe_order_number = $1
`

func (q *Queries) GetIdSummaryOrderMock8(ctx context.Context, poNo string) (GetSummaryOrderMock8, error) {
	row := q.db.QueryRowContext(ctx, getCartMockup8, poNo)
	var i GetSummaryOrderMock8
	err := row.Scan(
		&i.SoheOrderNumber,
		&i.ProgTitle,
		&i.ProgHeadline,
		&i.ProgImage,
		&i.ProgPrice,
		&i.SoheSubtotal,
		&i.TrpaNote,
		&i.SoheAccountNumber,
		&i.UserName,
		&i.TrpaCredit,
		&i.TrpaCodeNumber,
	)
	return i, err
}
