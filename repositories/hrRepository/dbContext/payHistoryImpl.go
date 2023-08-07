package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const createPayHistory = `-- name: CreatePayHistory :one

WITH inserted_entity AS (
	SELECT * FROM hr.employee
	ORDER BY emp_entity_id DESC
	LIMIT 1
  )
  INSERT INTO hr.employee_pay_history (ephi_entity_id, ephi_rate_change_date, ephi_rate_salary, ephi_pay_frequence, ephi_modified_date) 
  SELECT emp_entity_id, Now(), $1, $2, Now()
  FROM inserted_entity
  RETURNING ephi_entity_id, ephi_rate_change_date, ephi_rate_salary, ephi_pay_frequence, ephi_modified_date
`

type CreatePayHistoryParams struct {
	EphiEntityID       int32     `db:"ephi_entity_id" json:"ephiEntityId"`
	EphiRateChangeDate time.Time `db:"ephi_rate_change_date" json:"ephiRateChangeDate"`
	EphiRateSalary     int32     `db:"ephi_rate_salary" json:"ephiRateSalary"`
	EphiPayFrequence   int16     `db:"ephi_pay_frequence" json:"ephiPayFrequence"`
	EphiModifiedDate   time.Time `db:"ephi_modified_date" json:"ephiModifiedDate"`
}

func (q *Queries) CreatePayHistory(ctx context.Context, arg CreatePayHistoryParams) (*models.HrEmployeePayHistory, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPayHistory,
		arg.EphiRateSalary,
		arg.EphiPayFrequence,
	)
	i := models.HrEmployeePayHistory{}
	err := row.Scan(
		&i.EphiEntityID,
		&i.EphiRateChangeDate,
		&i.EphiRateSalary,
		&i.EphiPayFrequence,
		&i.EphiModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.HrEmployeePayHistory{
		EphiEntityID:       i.EphiEntityID,
		EphiRateChangeDate: i.EphiRateChangeDate,
		EphiRateSalary:     i.EphiRateSalary,
		EphiPayFrequence:   i.EphiPayFrequence,
		EphiModifiedDate:   i.EphiModifiedDate,
	}, nil
}

const listPayHistory = `-- name: ListPayHistory :many
SELECT ephi_entity_id, ephi_rate_change_date, ephi_rate_salary, ephi_pay_frequence, ephi_modified_date FROM hr.employee_pay_history
ORDER BY ephi_entity_id
`

func (q *Queries) ListPayHistory(ctx context.Context) ([]models.HrEmployeePayHistory, error) {
	rows, err := q.db.QueryContext(ctx, listPayHistory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.HrEmployeePayHistory
	for rows.Next() {
		var i models.HrEmployeePayHistory
		if err := rows.Scan(
			&i.EphiEntityID,
			&i.EphiRateChangeDate,
			&i.EphiRateSalary,
			&i.EphiPayFrequence,
			&i.EphiModifiedDate,
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

const getPayHistory = `-- name: GetPayHistory :one

SELECT ephi_entity_id, ephi_rate_change_date, ephi_rate_salary, ephi_pay_frequence, ephi_modified_date FROM hr.employee_pay_history
WHERE ephi_entity_id = $1
`

// hr.employee_pay_history
func (q *Queries) GetPayHistory(ctx context.Context, ephiEntityID int32) (models.HrEmployeePayHistory, error) {
	row := q.db.QueryRowContext(ctx, getPayHistory, ephiEntityID)
	var i models.HrEmployeePayHistory
	err := row.Scan(
		&i.EphiEntityID,
		&i.EphiRateChangeDate,
		&i.EphiRateSalary,
		&i.EphiPayFrequence,
		&i.EphiModifiedDate,
	)
	return i, err
}

const updatePayHistory = `-- name: UpdatePayHistory :exec
UPDATE hr.employee_pay_history
  set ephi_rate_change_date = $2,
  ephi_rate_salary = $3,
  ephi_pay_frequence = $4,
  ephi_modified_date = Now()
WHERE ephi_entity_id = $1
`

type UpdatePayHistoryParams struct {
	EphiEntityID       int32     `db:"ephi_entity_id" json:"ephiEntityId"`
	EphiRateChangeDate time.Time `db:"ephi_rate_change_date" json:"ephiRateChangeDate"`
	EphiRateSalary     int32     `db:"ephi_rate_salary" json:"ephiRateSalary"`
	EphiPayFrequence   int16     `db:"ephi_pay_frequence" json:"ephiPayFrequence"`
	EphiModifiedDate   time.Time `db:"ephi_modified_date" json:"ephiModifiedDate"`
}

func (q *Queries) UpdatePayHistory(ctx context.Context, arg UpdatePayHistoryParams) error {
	_, err := q.db.ExecContext(ctx, updatePayHistory, arg.EphiEntityID, arg.EphiRateChangeDate, arg.EphiRateSalary, arg.EphiPayFrequence)
	return err
}

const deletePayHistory = `-- name: DeletePayHistory :exec
DELETE FROM hr.employee_pay_history
WHERE ephi_entity_id = $1
`

func (q *Queries) DeletePayHistory(ctx context.Context, ephiEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePayHistory, ephiEntityID)
	return err
}
