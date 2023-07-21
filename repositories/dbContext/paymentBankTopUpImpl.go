package dbContext

// import (
// 	"context"
// 	"database/sql"
// 	"net/http"
// 	"time"

// 	"codeid.revampacademy/models"
// )

// // 1a. fungsi utk ambil getlist
// const listFintechTopup = `-- name: ListPaymentBank :many
// SELECT bank_entity_id, bank_code, bank_name, bank_modified_date FROM payment.bank ORDER BY bank_name
// `

// func (q *Queries) ListFintechTopup(ctx context.Context) ([]models.FintechTopup, error) {
// 	rows, err := q.db.QueryContext(ctx, listFintechTopup)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []models.FintechTopup
// 	for rows.Next() {
// 		var i models.FintechTopup
// 		if err := rows.Scan(
// 			&i.BankEntityID,
// 			&i.BankCode,
// 			&i.BankName,
// 			&i.BankModifiedDate,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// // 1a. fungsi utk ambil get payment bank
// const getFintechTopup = `-- name: GetPaymentBank :one
// SELECT bank_entity_id, bank_code, bank_name, bank_modified_date FROM payment.bank
// WHERE bank_code = $1`

// func (q *Queries) GetFintechTopup(ctx context.Context, name string) (models.FintechTopup, error) {
// 	row := q.db.QueryRowContext(ctx, getFintechTopup, name)
// 	var i models.FintechTopup
// 	err := row.Scan(
// 		&i.BankEntityID,
// 		&i.BankCode,
// 		&i.BankName,
// 		&i.BankModifiedDate,
// 	)
// 	return i, err
// }

// // 1.b fungsi utk create paymentbank
// const createFintechTopup = `-- name: CreatefintechTopup :one
// INSERT INTO
//     payment.bank(
//         bank_entity_id,
//         bank_code,
//         bank_name,
//         bank_modified_date
//     )
// VALUES ($1, $2, $3, $4) RETURNING *
// `

// type CreateFintechTopupParams struct {
// 	BankEntityID     int32        `db:"bank_entity_id" json:"bankEntityId"`
// 	BankCode         string       `db:"bank_code" json:"bankCode"`
// 	BankName         string       `db:"bank_name" json:"bankName"`
// 	BankModifiedDate sql.NullTime `db:"bank_modified_date" json:"bankModifiedDate"`
// }

// func (q *Queries) CreateFintechTopup(ctx context.Context, arg CreateFintechTopupParams) (*models.FintechTopup, *models.ResponseError) {
// 	row := q.db.QueryRowContext(ctx, createFintechTopup,
// 		arg.BankEntityID,
// 		arg.BankCode,
// 		arg.BankName,
// 		arg.BankModifiedDate,
// 	)

// 	i := models.FintechTopup{}
// 	err := row.Scan(
// 		&i.BankEntityID,
// 		&i.BankCode,
// 		&i.BankName,
// 		&i.BankModifiedDate,
// 	)

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return &models.FintechTopup{
// 		BankEntityID:     i.BankEntityID,
// 		BankCode:         i.BankCode,
// 		BankName:         i.BankName,
// 		BankModifiedDate: sql.NullTime{Time: time.Now(), Valid: true},
// 	}, nil
// }

// const updateFintechTopup = `-- name: UpdateFintechTopup :exec
// UPDATE payment.bank
//   set bank_code = $2,
//   bank_name = $3
// WHERE bank_entity_id = $1
// `

// func (q *Queries) UpdateFintechTopup(ctx context.Context, arg CreateFintechTopupParams) error {
// 	_, err := q.db.ExecContext(ctx, updateFintechTopup, arg.BankEntityID, arg.BankCode, arg.BankName)
// 	return err
// }

// const deleteFintechTopup = `-- name: DeletefintechTopup :exec
// DELETE FROM payment.bank
// WHERE bank_entity_id = $1
// `

// func (q *Queries) DeleteFintechTopup(ctx context.Context, bankEntityID int32) error {
// 	_, err := q.db.ExecContext(ctx, deleteFintechTopup, bankEntityID)
// 	return err
// }
