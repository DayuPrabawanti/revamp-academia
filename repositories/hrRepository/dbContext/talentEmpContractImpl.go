package dbContext

import (
	"context"
	"database/sql"
	"time"

	"codeid.revampacademy/models"
)

const getTalentClientContract = `-- name: GetTalentClientContract :one

SELECT 
ecco_id, ecco_entity_id, ecco_contract_no,  ecco_start_date, ecco_end_date, ecco_notes FROM hr.employee_client_contract
WHERE ecco_entity_id = $1
`

// hr.employee_client_contract
func (q *Queries) GetTalentClientContract(ctx context.Context, eccoEntityId int32) (models.TalentClientContractGetUpdate, error) {
	row := q.db.QueryRowContext(ctx, getTalentClientContract, eccoEntityId)
	var i models.TalentClientContractGetUpdate
	err := row.Scan(
		&i.HrEmployeeClientContract.EccoID,
		&i.HrEmployeeClientContract.EccoEntityID,
		&i.HrEmployeeClientContract.EccoContractNo,
		&i.HrEmployeeClientContract.EccoStartDate,
		&i.HrEmployeeClientContract.EccoEndDate,
		&i.HrEmployeeClientContract.EccoNotes,
	)
	return i, err
}

const updateTalentClientContract = `-- name: UpdateTalentClientContract :exec
UPDATE hr.employee_client_contract
  set 
  ecco_contract_no = $2,
  ecco_start_date = $3,
  ecco_end_date = $4,
  ecco_notes = $5,
  ecco_modified_date = $6,
WHERE ecco_entity_id = $1
`

type UpdateTalentClientContractParams struct {
	EccoEntityID     int32          `db:"ecco_entity_id" json:"eccoEntityId"`
	EccoContractNo   sql.NullString `db:"ecco_contract_no" json:"eccoContractNo"`
	EccoStartDate    sql.NullTime   `db:"ecco_start_date" json:"eccoStartDate"`
	EccoEndDate      sql.NullTime   `db:"ecco_end_date" json:"eccoEndDate"`
	EccoNotes        sql.NullString `db:"ecco_notes" json:"eccoNotes"`
	EccoModifiedDate sql.NullTime   `db:"ecco_modified_date" json:"eccoModifiedDate"`
}

func (q *Queries) UpdateTalentClientContract(ctx context.Context, arg UpdateTalentClientContractParams) error {
	_, err := q.db.ExecContext(ctx, updateTalentClientContract,
		arg.EccoEntityID,
		arg.EccoContractNo,
		arg.EccoStartDate,
		arg.EccoEndDate,
		arg.EccoNotes,
		sql.NullTime{Time: time.Now(), Valid: true},
	)
	return err
}
