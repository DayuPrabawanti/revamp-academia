package dbContext

import (
	"context"
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
)

type CreateMergeMock8 struct {
	Createprogram_entityParam
	CreateUsersParam
	CreatePaymentFintechParams
	CreatePaymentUsers_accountParams
	CreatePaymentTransaction_paymentParams
	CreateSales_order_detailParams
}

const createprogram_entitys = `-- name: Createprogram_entitys :one

INSERT INTO curriculum.program_entity (prog_entity_id, 
prog_title, 
prog_headline, 
prog_type, 
prog_learning_type, 
prog_rating, 
prog_total_trainee, 
prog_modified_date, 
prog_image, 
prog_best_seller, 
prog_price, 
prog_language, 
prog_duration, 
prog_duration_type, 
prog_tag_skill, 
prog_city_id, 
prog_cate_id, 
prog_created_by, 
prog_status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)
RETURNING prog_entity_id
`

type Createprogram_entityParam struct {
	ProgHeadline sql.NullString `db:"prog_headline" json:"progHeadline"`
	ProgImage    sql.NullString `db:"prog_image" json:"progImage"`
}

func (q *Queries) Createprogram_entitys(ctx context.Context, arg Createprogram_entityParam) (*Createprogram_entityParam, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createprogram_entitys,
		arg.ProgHeadline,
		arg.ProgImage,
	)
	i := Createprogram_entityParam{}
	err := row.Scan(
		&i.ProgHeadline,
		&i.ProgImage,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &Createprogram_entityParam{
		ProgHeadline: i.ProgHeadline,
		ProgImage:    i.ProgImage,
	}, nil

}

const createPaymentFintech = `-- name: CreatePaymentFintech :one

INSERT INTO
    payment.fintech (
        fint_entity_id,
        fint_code,
        fint_name,
        fint_modified_date
    )
VALUES ($1, $2, $3, $4) RETURNING fint_entity_id
`

type CreatePaymentFintechParams struct {
	FintName sql.NullString `db:"fint_name" json:"fintName"`
}

func (q *Queries) CreatePaymentFintech(ctx context.Context, arg CreatePaymentFintechParams) (*CreatePaymentFintechParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentFintech,
		arg.FintName,
	)
	i := CreatePaymentFintechParams{}
	err := row.Scan(
		&i.FintName,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreatePaymentFintechParams{
		FintName: i.FintName,
	}, nil

}

const createUser = `-- name: CreateUser :one

INSERT INTO users.users 
(user_first_name, user_last_name, user_birth_date, user_photo)
VALUES($1,$2,$3,$4)
RETURNING *
`

type CreateUsersParam struct {
	UserPhoto string `db:"user_photo" json:"userPhoto"`
}

func (q *Queries) CreateUsersParam(ctx context.Context, arg CreateUsersParam) (*CreateUsersParam, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserPhoto,
	)
	i := CreateUsersParam{}
	err := row.Scan(
		&i.UserPhoto,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateUsersParam{
		UserPhoto: i.UserPhoto,
	}, nil
}

const createPaymentUsers_account = `-- name: CreatePaymentUsers_account :one

INSERT INTO
    payment.users_account (
        usac_bank_entity_id,
        usac_user_entity_id,
        usac_account_number,
        usac_saldo,
        usac_type,
        usac_start_date,
        usac_end_date,
        usac_modified_date,
        usac_status
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING usac_bank_entity_id
`

type CreatePaymentUsers_accountParams struct {
	UsacAccountNumber sql.NullString `db:"usac_account_number" json:"usacAccountNumber"`
}

func (q *Queries) CreatePaymentUsers_account(ctx context.Context, arg CreatePaymentUsers_accountParams) (*CreatePaymentUsers_accountParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentUsers_account,
		arg.UsacAccountNumber,
	)
	i := CreatePaymentUsers_accountParams{}
	err := row.Scan(
		&i.UsacAccountNumber,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreatePaymentUsers_accountParams{
		UsacAccountNumber: i.UsacAccountNumber,
	}, nil
}

const createPaymentTransaction_payment = `-- name: CreatePaymentTransaction_payment :one

INSERT INTO
    payment.transaction_payment (
        trpa_id,
        trpa_code_number,
        trpa_order_number,
        trpa_debit,
        trpa_credit,
        trpa_type,
        trpa_note,
        trpa_modified_date,
        trpa_source_id,
        trpa_target_id,
        trpa_user_entity_id
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11
    ) RETURNING trpa_id
`

type CreatePaymentTransaction_paymentParams struct {
	TrpaCodeNumber int32          `db:"trpa_code_number" json:"trpaCodeNumber"`
	TrpaCredit     sql.NullString `db:"trpa_credit" json:"trpaCredit"`
	TrpaNote       sql.NullString `db:"trpa_note" json:"trpaNote"`
}

func (q *Queries) CreatePaymentTransaction_payment(ctx context.Context, arg CreatePaymentTransaction_paymentParams) (*CreatePaymentTransaction_paymentParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentTransaction_payment,
		arg.TrpaCodeNumber,
		arg.TrpaCredit,
		arg.TrpaNote,
	)
	i := CreatePaymentTransaction_paymentParams{}
	err := row.Scan(
		&i.TrpaCodeNumber,
		&i.TrpaCredit,
		&i.TrpaNote,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreatePaymentTransaction_paymentParams{
		TrpaCodeNumber: i.TrpaCodeNumber,
		TrpaCredit:     i.TrpaCredit,
		TrpaNote:       i.TrpaNote,
	}, nil
}

const createSales_order_detail = `-- name: CreateSales_order_detail :one
INSERT INTO sales.sales_order_detail(sode_id, sode_qty, sode_unit_price, sode_unit_discount,sode_line_total,sode_modified_date,sode_sohe_id,sode_prog_entity_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING sode_id
`

type CreateSales_order_detailParams struct {
	SodeLineTotal sql.NullInt32 `db:"sode_line_total" json:"sodeLineTotal"`
}

func (q *Queries) CreateSales_order_detail(ctx context.Context, arg CreateSales_order_detailParams) (*CreateSales_order_detailParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createSales_order_detail,
		arg.SodeLineTotal,
	)
	i := CreateSales_order_detailParams{}
	err := row.Scan(
		&i.SodeLineTotal,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateSales_order_detailParams{
		SodeLineTotal: i.SodeLineTotal,
	}, nil
}

const listMock8Grup = `-- name: ListMock8Group :many
select pe.prog_headline,pe.prog_image,u.user_photo,sod.sode_line_total,f.fint_name,ua.usac_account_number
,tp.trpa_note,tp.trpa_credit,tp.trpa_code_number
from payment.transaction_payment tp
join users.users u
on tp.trpa_user_entity_id = u.user_entity_id
join payment.users_account ua
on ua.usac_user_entity_id = u.user_entity_id
join payment.fintech f
on f.fint_entity_id = ua.usac_bank_entity_id
join curriculum.program_reviews pr
on pr.prow_prog_entity_id = u.user_entity_id
join curriculum.program_entity pe
on pe.prog_entity_id = pr.prow_prog_entity_id
join sales.special_offer_programs sop
on sop.soco_prog_entity_id = pe.prog_entity_id
join sales.sales_order_detail sod
on sod.sode_id = sop.soco_id
`

func (q *Queries) ListMock8Group(ctx context.Context) ([]CreateMergeMock8, error) {
	rows, err := q.db.QueryContext(ctx, listMock8Grup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreateMergeMock8
	for rows.Next() {
		var i CreateMergeMock8
		if err := rows.Scan(
			&i.ProgHeadline,
			&i.ProgImage,
			&i.UserPhoto,
			&i.SodeLineTotal,
			&i.FintName,
			&i.UsacAccountNumber,
			&i.TrpaNote,
			&i.TrpaCredit,
			&i.TrpaCodeNumber,
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

const getMock8Grup = `-- name: GetMock8Group :many
select tp.trpa_code_number,pe.prog_headline,pe.prog_image,u.user_photo,sod.sode_line_total,f.fint_name,ua.usac_account_number
,tp.trpa_note,tp.trpa_credit
from payment.transaction_payment tp
join users.users u
on tp.trpa_user_entity_id = u.user_entity_id
join payment.users_account ua
on ua.usac_user_entity_id = u.user_entity_id
join payment.fintech f
on f.fint_entity_id = ua.usac_bank_entity_id
join curriculum.program_reviews pr
on pr.prow_prog_entity_id = u.user_entity_id
join curriculum.program_entity pe
on pe.prog_entity_id = pr.prow_prog_entity_id
join sales.special_offer_programs sop
on sop.soco_prog_entity_id = pe.prog_entity_id
join sales.sales_order_detail sod
on sod.sode_id = sop.soco_id
where tp.trpa_code_number = $1
`

func (q *Queries) GetMock8Group(ctx context.Context, poNo string) (models.MergeMock8, error) {
	row := q.db.QueryRowContext(ctx, getMock8Grup, poNo)
	var i models.MergeMock8
	err := row.Scan(
		&i.Transaction.TrpaCodeNumber,
		&i.Curriculums.ProgHeadline,
		&i.Curriculums.ProgImage,
		&i.Photo.UserPhoto,
		&i.Total.SodeLineTotal,
		&i.Fint.FintName,
		&i.Number.UsacAccountNumber,
		&i.Transaction.TrpaNote,
		&i.Transaction.TrpaCredit,
	)
	return i, err
}
