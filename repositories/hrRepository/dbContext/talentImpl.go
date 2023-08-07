package dbContext

import (
	"context"
	"time"

	"codeid.revampacademy/models"
)

const listTalents = `-- name: ListTalents :many
SELECT
us.user_first_name, us.user_last_name, us.user_photo,
bc.batch_name, bc.batch_type, bc.batch_status,
bte.btev_skor,
pe.prog_title
FROM users.users us
JOIN bootcamp.batch bc
ON bc.batch_entity_id = us.user_entity_id
JOIN bootcamp.batch_trainee_evaluation bte
ON bc.batch_id = bte.btev_batch_id
JOIN curriculum.program_entity pe
ON us.user_entity_id = pe.prog_entity_id
ORDER BY us.user_entity_id
`

func (q *Queries) ListTalents(ctx context.Context) ([]models.TalentsMockup, error) {
	rows, err := q.db.QueryContext(ctx, listTalents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.TalentsMockup
	for rows.Next() {
		var i models.TalentsMockup
		if err := rows.Scan(
			&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,
			&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStatus,
			&i.BootcampBatchTraineeEvaluation.BtevSkor,
			&i.CurriculumProgramEntity.ProgTitle,
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

const searchTalent = `-- name: SearchTalent :many
SELECT
us.user_first_name, us.user_last_name, us.user_photo,
bc.batch_name, bc.batch_type, bc.batch_status,
bte.btev_skor,
pe.prog_title
FROM users.users us
JOIN bootcamp.batch bc
ON bc.batch_entity_id = us.user_entity_id
JOIN bootcamp.batch_trainee_evaluation bte
ON bc.batch_id = bte.btev_batch_id
JOIN curriculum.program_entity pe
ON us.user_entity_id = pe.prog_entity_id
WHERE us.user_name like '%' || $1 || '%' OR pe.prog_title like '%' || $2 || '%' OR bc.batch_status like '%' || $3 || '%'

`

func (q *Queries) SearchTalent(ctx context.Context, userName, skillName, batchStatus string) ([]models.TalentsMockup, error) {
	rows, err := q.db.QueryContext(ctx, searchTalent, userName, skillName, batchStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var talents []models.TalentsMockup
	for rows.Next() {
		var i models.TalentsMockup
		if err := rows.Scan(
			&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,
			&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStatus,
			&i.BootcampBatchTraineeEvaluation.BtevSkor,
			&i.CurriculumProgramEntity.ProgTitle,
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

const pagingTalent = `-- name: PagingTalent :many

SELECT
us.user_first_name, us.user_last_name, us.user_photo,
bc.batch_name, bc.batch_type, bc.batch_status,
bte.btev_skor,
pe.prog_title
FROM users.users us
JOIN bootcamp.batch bc
ON bc.batch_entity_id = us.user_entity_id
JOIN bootcamp.batch_trainee_evaluation bte
ON bc.batch_id = bte.btev_batch_id
JOIN curriculum.program_entity pe
ON us.user_entity_id = pe.prog_entity_id
LIMIT $1 OFFSET $2
`

func (q *Queries) PagingTalent(ctx context.Context, limit, offset int) ([]models.TalentsMockup, error) {
	rows, err := q.db.QueryContext(ctx, pagingTalent, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var talents []models.TalentsMockup
	for rows.Next() {
		var i models.TalentsMockup
		if err := rows.Scan(
			&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,
			&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStatus,
			&i.BootcampBatchTraineeEvaluation.BtevSkor,
			&i.CurriculumProgramEntity.ProgTitle,
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

const getBatch = `-- name: GetBatch :one
SELECT batch_id, batch_start_date, batch_reason, batch_modified_date, batch_status FROM bootcamp.batch
WHERE batch_id = $1
`

func (q *Queries) GetBatch(ctx context.Context, batchID int32) (models.BootcampBatch, error) {
	row := q.db.QueryRowContext(ctx, getBatch, batchID)
	var i models.BootcampBatch
	err := row.Scan(
		&i.BatchID,
		&i.BatchStartDate,
		&i.BatchReason,
		&i.BatchModifiedDate,
		&i.BatchStatus,
	)
	return i, err
}

const updateBatch = `-- name: UpdateBatch :exec
UPDATE bootcamp.batch
SET batch_start_date = $2,
	batch_reason = $3,
	batch_modified_date = Now(),
    batch_status = $4
WHERE batch_id = $1
`

type UpdateBatchParams struct {
	BatchID           int32     `db:"batch_id" json:"batchId"`
	BatchStartDate    time.Time `db:"batch_start_date" json:"batchStartDate"`
	BatchReason       string    `db:"batch_reason" json:"batchReason"`
	BatchModifiedDate time.Time `db:"batch_modified_date" json:"batchModifiedDate"`
	BatchStatus       string    `db:"batch_status" json:"batchStatus"`
}

func (q *Queries) UpdateBatch(ctx context.Context, arg UpdateBatchParams) error {
	_, err := q.db.ExecContext(ctx, updateBatch,
		arg.BatchID,
		arg.BatchStartDate,
		arg.BatchReason,
		arg.BatchStatus,
	)
	return err
}
