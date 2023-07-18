package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revamptwo/models/hrsMdl"
)

const listClientContract = `-- name: ListClientContract :many
SELECT ecco_id, ecco_entity_id, ecco_contract_no, ecco_contract_date, ecco_start_date, ecco_end_date, ecco_notes, ecco_modified_date, ecco_media_link, ecco_joty_id, ecco_account_manager, ecco_clit_id, ecco_status FROM hr.employee_client_contract
ORDER BY ecco_id
`

func (q *Queries) GetListEmployeeClientContract(ctx context.Context) ([]hrsMdl.HrEmployeeClientContract, error) {
	rows, err := q.db.QueryContext(ctx, listClientContract)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []hrsMdl.HrEmployeeClientContract
	for rows.Next() {
		var i hrsMdl.HrEmployeeClientContract
		if err := rows.Scan(
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

SELECT ecco_id, ecco_entity_id, ecco_contract_no, ecco_contract_date, ecco_start_date, ecco_end_date, ecco_notes, ecco_modified_date, ecco_media_link, ecco_joty_id, ecco_account_manager, ecco_clit_id, ecco_status FROM hr.employee_client_contract
WHERE ecco_id = $1
`

// hr.employee_client_contract
func (q *Queries) GetEmployeeClientContract(ctx context.Context, eccoID int32) (hrsMdl.HrEmployeeClientContract, error) {
	row := q.db.QueryRowContext(ctx, getClientContract, eccoID)
	var i hrsMdl.HrEmployeeClientContract
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
	return i, err
}

const createClientContract = `-- name: CreateClientContract :one

INSERT INTO hr.employee_client_contract (ecco_id, ecco_entity_id, ecco_contract_no, ecco_contract_date, ecco_start_date, ecco_end_date, ecco_notes, ecco_modified_date, ecco_media_link, ecco_joty_id, ecco_account_manager, ecco_clit_id, ecco_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING *
`

type CreateEmployeeClientContractParams struct {
	EccoID             int32          `db:"ecco_id" json:"eccoId"`
	EccoEntityID       int32          `db:"ecco_entity_id" json:"eccoEntityId"`
	EccoContractNo     sql.NullString `db:"ecco_contract_no" json:"eccoContractNo"`
	EccoContractDate   sql.NullTime   `db:"ecco_contract_date" json:"eccoContractDate"`
	EccoStartDate      sql.NullTime   `db:"ecco_start_date" json:"eccoStartDate"`
	EccoEndDate        sql.NullTime   `db:"ecco_end_date" json:"eccoEndDate"`
	EccoNotes          sql.NullString `db:"ecco_notes" json:"eccoNotes"`
	EccoModifiedDate   sql.NullTime   `db:"ecco_modified_date" json:"eccoModifiedDate"`
	EccoMediaLink      sql.NullString `db:"ecco_media_link" json:"eccoMediaLink"`
	EccoJotyID         sql.NullInt32  `db:"ecco_joty_id" json:"eccoJotyId"`
	EccoAccountManager sql.NullInt32  `db:"ecco_account_manager" json:"eccoAccountManager"`
	EccoClitID         sql.NullInt32  `db:"ecco_clit_id" json:"eccoClitId"`
	EccoStatus         sql.NullString `db:"ecco_status" json:"eccoStatus"`
}

func (q *Queries) CreateEmployeeClientContract(ctx context.Context, arg CreateEmployeeClientContractParams) (*hrsMdl.HrEmployeeClientContract, *hrsMdl.ResponseError) {
	row := q.db.QueryRowContext(ctx, createClientContract,
		arg.EccoID,
		arg.EccoEntityID,
		arg.EccoContractNo,
		arg.EccoContractDate,
		arg.EccoStartDate,
		arg.EccoEndDate,
		arg.EccoNotes,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.EccoMediaLink,
		arg.EccoJotyID,
		arg.EccoAccountManager,
		arg.EccoClitID,
		arg.EccoStatus,
	)
	i := hrsMdl.HrEmployeeClientContract{}
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
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &hrsMdl.HrEmployeeClientContract{
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

const updateClientContract = `-- name: UpdateClientContract :exec
UPDATE hr.employee_client_contract
  set ecco_start_date = $2,
  ecco_end_date = $3,
  ecco_notes = $4,
  ecco_modified_date = $5,
  ecco_media_link = $6,
  ecco_joty_id= $7,
  ecco_account_manager= $8,
  ecco_clit_id = $9,
  ecco_status = $10
WHERE ecco_id = $1
`

func (q *Queries) UpdateEmployeeClientContract(ctx context.Context, arg CreateEmployeeClientContractParams) error {
	_, err := q.db.ExecContext(ctx, updateClientContract,
		arg.EccoID,
		arg.EccoStartDate,
		arg.EccoEndDate,
		arg.EccoNotes,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.EccoMediaLink,
		arg.EccoJotyID,
		arg.EccoAccountManager,
		arg.EccoClitID,
		arg.EccoStatus,
	)
	return err
}

const deleteClientContract = `-- name: DeleteClientContract :exec
DELETE FROM hr.employee_client_contract
WHERE ecco_id = $1
`

func (q *Queries) DeleteEmployeeClientContract(ctx context.Context, eccoID int32) error {
	_, err := q.db.ExecContext(ctx, deleteClientContract, eccoID)
	return err
}
