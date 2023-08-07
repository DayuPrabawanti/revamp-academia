package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const createClientContract = `-- name: CreateClientContract :one

WITH inserted_entity AS (
	SELECT * FROM hr.employee
	ORDER BY emp_entity_id DESC
	LIMIT 1
  )
INSERT INTO hr.employee_client_contract 
(ecco_entity_id, ecco_id, ecco_modified_date) 
SELECT emp_entity_id, $1, Now()
FROM inserted_entity
RETURNING ecco_id, ecco_entity_id, ecco_contract_no, ecco_contract_date, ecco_start_date, ecco_end_date, ecco_notes, ecco_modified_date, ecco_media_link, ecco_joty_id, ecco_account_manager, ecco_clit_id, ecco_status
`

type CreateClientContractParams struct {
	EccoID             int32     `db:"ecco_id" json:"eccoId"`
	EccoEntityID       int32     `db:"ecco_entity_id" json:"eccoEntityId"`
	EccoContractNo     string    `db:"ecco_contract_no" json:"eccoContractNo"`
	EccoContractDate   time.Time `db:"ecco_contract_date" json:"eccoContractDate"`
	EccoStartDate      time.Time `db:"ecco_start_date" json:"eccoStartDate"`
	EccoEndDate        time.Time `db:"ecco_end_date" json:"eccoEndDate"`
	EccoNotes          string    `db:"ecco_notes" json:"eccoNotes"`
	EccoModifiedDate   time.Time `db:"ecco_modified_date" json:"eccoModifiedDate"`
	EccoMediaLink      string    `db:"ecco_media_link" json:"eccoMediaLink"`
	EccoJotyID         int32     `db:"ecco_joty_id" json:"eccoJotyId"`
	EccoAccountManager int32     `db:"ecco_account_manager" json:"eccoAccountManager"`
	EccoClitID         int32     `db:"ecco_clit_id" json:"eccoClitId"`
	EccoStatus         string    `db:"ecco_status" json:"eccoStatus"`
}

func (q *Queries) CreateClientContract(ctx context.Context, arg CreateClientContractParams) (*models.HrEmployeeClientContract, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createClientContract,
		arg.EccoID,
	)
	i := models.HrEmployeeClientContract{}
	err := row.Scan(
		&i.EccoID,
		&i.EccoEntityID,
		&i.EccoContractNo,
		&i.EccoContractDate,
		&i.EccoStartDate,
		&i.EccoEndDate,
		&i.EccoNotes,
		&i.EccoModifiedDate,
		&i.EccoMediaLink,
		&i.EccoJotyID,
		&i.EccoAccountManager,
		&i.EccoClitID,
		&i.EccoStatus,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.HrEmployeeClientContract{
		EccoID:             i.EccoID,
		EccoEntityID:       i.EccoEntityID,
		EccoContractNo:     i.EccoContractNo,
		EccoContractDate:   i.EccoContractDate,
		EccoStartDate:      i.EccoStartDate,
		EccoEndDate:        i.EccoEndDate,
		EccoNotes:          i.EccoNotes,
		EccoModifiedDate:   i.EccoModifiedDate,
		EccoMediaLink:      i.EccoMediaLink,
		EccoJotyID:         i.EccoJotyID,
		EccoAccountManager: i.EccoAccountManager,
		EccoClitID:         i.EccoClitID,
		EccoStatus:         i.EccoStatus,
	}, nil
}

const listClientContract = `-- name: ListClientContract :many
SELECT ecco_id, ecco_entity_id, ecco_contract_no, ecco_start_date, ecco_end_date, ecco_notes
FROM hr.employee_client_contract
ORDER BY ecco_id
`

func (q *Queries) ListClientContract(ctx context.Context) ([]models.HrEmployeeClientContract, error) {
	rows, err := q.db.QueryContext(ctx, listClientContract)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.HrEmployeeClientContract
	for rows.Next() {
		var i models.HrEmployeeClientContract
		if err := rows.Scan(
			&i.EccoID,
			&i.EccoEntityID,
			&i.EccoContractNo,
			&i.EccoStartDate,
			&i.EccoEndDate,
			&i.EccoNotes,
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

const getClientContract = `-- name: GetClientContract :one

SELECT ecco_id, ecco_entity_id, ecco_contract_no, ecco_start_date, ecco_end_date, ecco_notes 
FROM hr.employee_client_contract
WHERE ecco_id = $1
`

// hr.employee_client_contract
func (q *Queries) GetClientContract(ctx context.Context, eccoID int32) (models.HrEmployeeClientContract, error) {
	row := q.db.QueryRowContext(ctx, getClientContract, eccoID)
	var i models.HrEmployeeClientContract
	err := row.Scan(
		&i.EccoID,
		&i.EccoEntityID,
		&i.EccoContractNo,
		&i.EccoStartDate,
		&i.EccoEndDate,
		&i.EccoNotes,
	)
	return i, err
}

const updateClientContract = `-- name: UpdateClientContract :exec
UPDATE hr.employee_client_contract
  set ecco_entity_id = $2,
  ecco_contract_no = $3,
  ecco_contract_date = $4,
  ecco_start_date = $5,
  ecco_end_date = $6,
  ecco_status = $7,
  ecco_notes = $8,
  ecco_modified_date = Now()
WHERE ecco_id = $1
`

func (q *Queries) UpdateClientContract(ctx context.Context, arg CreateClientContractParams) error {
	_, err := q.db.ExecContext(ctx, updateClientContract,
		arg.EccoID,
		arg.EccoEntityID,
		arg.EccoContractNo,
		arg.EccoContractDate,
		arg.EccoStartDate,
		arg.EccoEndDate,
		arg.EccoStatus,
		arg.EccoNotes,
	)
	return err
}

const deleteClientContract = `-- name: DeleteClientContract :exec
DELETE FROM hr.employee_client_contract
WHERE ecco_id = $1
`

func (q *Queries) DeleteClientContract(ctx context.Context, eccoID int32) error {
	_, err := q.db.ExecContext(ctx, deleteClientContract, eccoID)
	return err
}
