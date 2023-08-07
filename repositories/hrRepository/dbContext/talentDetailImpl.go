package dbContext

import (
	"context"
	"database/sql"
	"time"

	"codeid.revampacademy/models"
)

const listTalentsDetail = `-- name: ListTalentsDetail :many

SELECT 
us.user_first_name, us.user_last_name, us.user_photo,
bc.batch_name, bc.batch_type, bc.batch_start_date, bc.batch_end_date, bc.batch_status,
bt.batr_review,
bte.btev_skor,
ue.pmail_address,
up.uspo_number,
pe.prog_title,
jp.jopo_title,
jc.clit_name,
empcc.ecco_start_date, empcc.ecco_end_date

FROM users.users us
JOIN bootcamp.batch bc
ON bc.batch_entity_id = us.user_entity_id
JOIN bootcamp.batch_trainee bt
ON bc.batch_id = bt.batr_batch_id
JOIN bootcamp.batch_trainee_evaluation bte
ON bc.batch_id = bte.btev_batch_id
JOIN curriculum.program_entity pe
ON us.user_entity_id = pe.prog_entity_id
JOIN users.users_email ue
ON us.user_entity_id = ue.pmail_entity_id
JOIN users.users_phones up
ON us.user_entity_id = up.uspo_entity_id
JOIN jobhire.job_post jp
ON us.user_entity_id = jp.jopo_entity_id
JOIN jobhire.client jc
ON jc.clit_id = jp.jopo_clit_id
JOIN hr.employee emp
ON us.user_entity_id = emp.emp_entity_id
JOIN hr.employee_client_contract empcc
ON emp.emp_entity_id = empcc.ecco_entity_id

ORDER BY us.user_entity_id
`

func (q *Queries) ListTalentsDetail(ctx context.Context) ([]models.TalentsDetailMockup, error) {
	rows, err := q.db.QueryContext(ctx, listTalentsDetail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.TalentsDetailMockup
	for rows.Next() {
		var i models.TalentsDetailMockup
		if err := rows.Scan(&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,
			&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStartDate, &i.BootcampBatch.BatchEndDate, &i.BootcampBatch.BatchStatus,
			&i.BootcampBatchTrainee.BatrReview,
			&i.BootcampBatchTraineeEvaluation.BtevSkor,
			&i.UsersUsersEmail.PmailAddress,
			&i.UsersUsersPhone.UspoNumber,
			&i.CurriculumProgramEntity.ProgTitle,
			&i.JobhireJobPost.JopoTitle,
			&i.JobhireClient.ClitName,
			&i.HrEmployeeClientContract.EccoStartDate, &i.HrEmployeeClientContract.EccoEndDate,
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

const getTalentDetail = `-- name: GetTalentDetail :one

SELECT 
us.user_first_name, us.user_last_name, us.user_photo,
bc.batch_name, bc.batch_type, bc.batch_start_date, bc.batch_end_date, bc.batch_status,
bt.batr_review,
bte.btev_skor,
ue.pmail_address,
up.uspo_number,
pe.prog_title,
jp.jopo_title,
jc.clit_name,
empcc.ecco_start_date, empcc.ecco_end_date

FROM users.users us
JOIN bootcamp.batch bc
ON bc.batch_entity_id = us.user_entity_id
JOIN bootcamp.batch_trainee bt
ON bc.batch_id = bt.batr_batch_id
JOIN bootcamp.batch_trainee_evaluation bte
ON bc.batch_id = bte.btev_batch_id
JOIN curriculum.program_entity pe
ON us.user_entity_id = pe.prog_entity_id
JOIN users.users_email ue
ON us.user_entity_id = ue.pmail_entity_id
JOIN users.users_phones up
ON us.user_entity_id = up.uspo_entity_id
JOIN jobhire.job_post jp
ON us.user_entity_id = jp.jopo_entity_id
JOIN jobhire.client jc
ON jc.clit_id = jp.jopo_clit_id
JOIN hr.employee emp
ON us.user_entity_id = emp.emp_entity_id
JOIN hr.employee_client_contract empcc
ON emp.emp_entity_id = empcc.ecco_entity_id

WHERE us.user_entity_id = $1
`

func (q *Queries) GetTalentDetail(ctx context.Context, batchId int32) (models.TalentsDetailMockup, error) {
	row := q.db.QueryRowContext(ctx, getTalentDetail, batchId)
	var i models.TalentsDetailMockup
	err := row.Scan(&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,
		&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStartDate, &i.BootcampBatch.BatchEndDate, &i.BootcampBatch.BatchStatus,
		&i.BootcampBatchTrainee.BatrReview,
		&i.BootcampBatchTraineeEvaluation.BtevSkor,
		&i.UsersUsersEmail.PmailAddress,
		&i.UsersUsersPhone.UspoNumber,
		&i.CurriculumProgramEntity.ProgTitle,
		&i.JobhireJobPost.JopoTitle,
		&i.JobhireClient.ClitName,
		&i.HrEmployeeClientContract.EccoStartDate, &i.HrEmployeeClientContract.EccoEndDate,
	)
	return i, err
}

const searchTalentDetail = `-- name: SearchTalentDetail :many
SELECT empcc.ecco_id,
jc.clit_name,
empcc.ecco_contract_no, empcc.ecco_start_date, 
empcc.ecco_end_date, empcc.ecco_status, empcc.ecco_notes
FROM hr.employee_client_contract empcc
JOIN jobhire.client jc
ON empcc.ecco_clit_id = jc.clit_id
WHERE jc.clit_name ilike $1 || '%'
`

func (q *Queries) SearchTalentDetail(ctx context.Context, clitName string) ([]models.TalentDetailSearchUpdate, error) {
	rows, err := q.db.QueryContext(ctx, searchTalentDetail, clitName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var talents []models.TalentDetailSearchUpdate
	for rows.Next() {
		var i models.TalentDetailSearchUpdate
		if err := rows.Scan(
			&i.HrEmployeeClientContract.EccoID,
			&i.JobhireClient.ClitName,
			&i.HrEmployeeClientContract.EccoContractNo,
			&i.HrEmployeeClientContract.EccoStartDate,
			&i.HrEmployeeClientContract.EccoEndDate,
			&i.HrEmployeeClientContract.EccoStatus,
			&i.HrEmployeeClientContract.EccoNotes,
		); err != nil {
			return nil, err
		}
		talents = append(talents, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return talents, nil
}

const updateSwitch = `-- name: UpdateSwitch :exec
UPDATE bootcamp.batch
SET batch_start_date = $2,
	batch_reason = $3,
	batch_modified_date = Now(),
    batch_status = $4
WHERE batch_id = $1
`

type UpdateSwitchParams struct {
	BatchID           int32          `db:"batch_id" json:"batchId"`
	BatchStartDate    time.Time      `db:"batch_start_date" json:"batchStartDate"`
	BatchReason       string         `db:"batch_reason" json:"batchReason"`
	BatchModifiedDate time.Time      `db:"batch_modified_date" json:"batchModifiedDate"`
	BatchStatus       sql.NullString `db:"batch_status" json:"batchStatus"`
}

func (q *Queries) UpdateSwitch(ctx context.Context, arg UpdateSwitchParams) error {
	// Set the default value "Idle" for BatchStatus if it is not provided or empty
	if !arg.BatchStatus.Valid || arg.BatchStatus.String == "" {
		arg.BatchStatus.String = "Idle"
		arg.BatchStatus.Valid = true
	}

	_, err := q.db.ExecContext(ctx, updateSwitch,
		arg.BatchID,
		arg.BatchStartDate,
		arg.BatchReason,
		arg.BatchStatus,
	)
	return err
}
