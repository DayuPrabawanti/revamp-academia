package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const listPaymentTransaction_payment = `-- name: ListPaymentTransaction_payment :many
SELECT trpa_id, trpa_code_number, trpa_order_number, trpa_debit, trpa_credit, trpa_type, trpa_note, 
trpa_modified_date, trpa_source_id, trpa_target_id, trpa_user_entity_id FROM payment.transaction_payment ORDER BY trpa_code_number
`

func (q *Queries) ListPaymentTransaction_payment(ctx context.Context) ([]models.PaymentTransactionPayment, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentTransaction_payment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.PaymentTransactionPayment
	for rows.Next() {
		var i models.PaymentTransactionPayment
		if err := rows.Scan(
			&i.TrpaID,
			&i.TrpaCodeNumber,
			&i.TrpaOrderNumber,
			&i.TrpaDebit,
			&i.TrpaCredit,
			&i.TrpaType,
			&i.TrpaNote,
			&i.TrpaModifiedDate,
			&i.TrpaSourceID,
			&i.TrpaTargetID,
			&i.TrpaUserEntityID,
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

const getPaymentTransaction_payment = `-- name: GetPaymentTransaction_payment :one
SELECT trpa_id, trpa_code_number, trpa_order_number, trpa_debit, trpa_credit, trpa_type, trpa_note, 
trpa_modified_date, trpa_source_id, trpa_target_id, trpa_user_entity_id FROM payment.transaction_payment WHERE trpa_id = $1
`

// payment.transaction_payment
func (q *Queries) GetPaymentTransaction_payment(ctx context.Context, trpaID int32) (models.PaymentTransactionPayment, error) {
	row := q.db.QueryRowContext(ctx, getPaymentTransaction_payment, trpaID)
	var i models.PaymentTransactionPayment
	err := row.Scan(
		&i.TrpaID,
		&i.TrpaCodeNumber,
		&i.TrpaOrderNumber,
		&i.TrpaDebit,
		&i.TrpaCredit,
		&i.TrpaType,
		&i.TrpaNote,
		&i.TrpaModifiedDate,
		&i.TrpaSourceID,
		&i.TrpaTargetID,
		&i.TrpaUserEntityID,
	)
	return i, err
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
    ) RETURNING *
`

type CreatePaymentTransaction_paymentParams struct {
	TrpaID           int32          `db:"trpa_id" json:"trpaId"`
	TrpaCodeNumber   sql.NullString `db:"trpa_code_number" json:"trpaCodeNumber"`
	TrpaOrderNumber  sql.NullString `db:"trpa_order_number" json:"trpaOrderNumber"`
	TrpaDebit        sql.NullString `db:"trpa_debit" json:"trpaDebit"`
	TrpaCredit       sql.NullString `db:"trpa_credit" json:"trpaCredit"`
	TrpaType         sql.NullString `db:"trpa_type" json:"trpaType"`
	TrpaNote         sql.NullString `db:"trpa_note" json:"trpaNote"`
	TrpaModifiedDate sql.NullTime   `db:"trpa_modified_date" json:"trpaModifiedDate"`
	TrpaSourceID     string         `db:"trpa_source_id" json:"trpaSourceId"`
	TrpaTargetID     string         `db:"trpa_target_id" json:"trpaTargetId"`
	TrpaUserEntityID sql.NullInt32  `db:"trpa_user_entity_id" json:"trpaUserEntityId"`
}

func (q *Queries) CreatePaymentTransaction_payment(ctx context.Context, arg CreatePaymentTransaction_paymentParams) (*models.PaymentTransactionPayment, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPaymentTransaction_payment,
		arg.TrpaID,
		arg.TrpaCodeNumber,
		arg.TrpaOrderNumber,
		arg.TrpaDebit,
		arg.TrpaCredit,
		arg.TrpaType,
		arg.TrpaNote,
		arg.TrpaModifiedDate,
		arg.TrpaSourceID,
		arg.TrpaTargetID,
		arg.TrpaUserEntityID,
	)
	i := models.PaymentTransactionPayment{}
	err := row.Scan(
		&i.TrpaID,
		&i.TrpaCodeNumber,
		&i.TrpaOrderNumber,
		&i.TrpaDebit,
		&i.TrpaCredit,
		&i.TrpaType,
		&i.TrpaNote,
		&i.TrpaModifiedDate,
		&i.TrpaSourceID,
		&i.TrpaTargetID,
		&i.TrpaUserEntityID,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.PaymentTransactionPayment{
		TrpaID:           i.TrpaID,
		TrpaCodeNumber:   i.TrpaCodeNumber,
		TrpaOrderNumber:  i.TrpaOrderNumber,
		TrpaDebit:        i.TrpaDebit,
		TrpaCredit:       i.TrpaCredit,
		TrpaType:         i.TrpaType,
		TrpaNote:         i.TrpaNote,
		TrpaModifiedDate: sql.NullTime{Time: time.Now(), Valid: true},
		TrpaSourceID:     i.TrpaSourceID,
		TrpaTargetID:     i.TrpaTargetID,
		TrpaUserEntityID: i.TrpaUserEntityID,
	}, nil
}

const updatePaymentTransaction_payment = `-- name: UpdatePaymentTransaction_payment :exec

UPDATE
    payment.transaction_payment
set
    trpa_code_number = $2,
    trpa_order_number = $3,
    trpa_debit = $4,
    trpa_credit = $5,
    trpa_type = $6,
    trpa_note = $7,
    trpa_modified_date = $8,
    trpa_source_id = $9,
    trpa_target_id = $10,
    trpa_user_entity_id = $11
WHERE trpa_id = $1
`

func (q *Queries) UpdatePaymentTransaction_payment(ctx context.Context, arg CreatePaymentTransaction_paymentParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentTransaction_payment,
		arg.TrpaID,
		arg.TrpaCodeNumber,
		arg.TrpaOrderNumber,
		arg.TrpaDebit,
		arg.TrpaCredit,
		arg.TrpaType,
		arg.TrpaNote,
		sql.NullTime{Time: arg.TrpaModifiedDate.Time, Valid: true}, //
		arg.TrpaSourceID,
		arg.TrpaTargetID,
		arg.TrpaUserEntityID,
	)
	return err
}

const deletePaymentTransaction_payment = `-- name: DeletePaymentTransaction_payment :exec
DELETE FROM payment.transaction_payment WHERE trpa_id = $1
`

func (q *Queries) DeletePaymentTransaction_payment(ctx context.Context, trpaID int32) error {
	_, err := q.db.ExecContext(ctx, deletePaymentTransaction_payment, trpaID)
	return err
}
