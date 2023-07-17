package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listBatchTraineeEvaluations = `-- name: ListBatchTraineeEvaluations :many
SELECT btev_id, btev_type, btev_header, btev_section, btev_skill, btev_week, btev_skor, btev_note, btev_modified_date, btev_batch_id, btev_trainee_entity_id FROM bootcamp.batch_trainee_evaluation
ORDER BY btev_id
`

func (q *Queries) ListBatchTraineeEvaluations(ctx context.Context) ([]models.BootcampBatchTraineeEvaluation, error) {
	rows, err := q.db.QueryContext(ctx, listBatchTraineeEvaluations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BootcampBatchTraineeEvaluation
	for rows.Next() {
		var i models.BootcampBatchTraineeEvaluation
		if err := rows.Scan(
			&i.BtevID,
			&i.BtevType,
			&i.BtevHeader,
			&i.BtevSection,
			&i.BtevSkill,
			&i.BtevWeek,
			&i.BtevSkor,
			&i.BtevNote,
			&i.BtevModifiedDate,
			&i.BtevBatchID,
			&i.BtevTraineeEntityID,
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

// const getBatchTraineeEvaluation = `-- name: GetBatchTraineeEvaluation :one
// SELECT btev_id, btev_type, btev_header, btev_section, btev_skill, btev_week, btev_skor, btev_note, btev_modified_date, btev_batch_id, btev_trainee_entity_id FROM bootcamp.batch_trainee_evaluation
// WHERE btev_id = $1
// `

// func (q *Queries) GetBatchTraineeEvaluation(ctx context.Context, btevID int32) (models.BootcampBatchTraineeEvaluation, error) {
// 	row := q.db.QueryRowContext(ctx, getBatchTraineeEvaluation, btevID)
// 	var i models.BootcampBatchTraineeEvaluation
// 	err := row.Scan(
// 		&i.BtevID,
// 		&i.BtevType,
// 		&i.BtevHeader,
// 		&i.BtevSection,
// 		&i.BtevSkill,
// 		&i.BtevWeek,
// 		&i.BtevSkor,
// 		&i.BtevNote,
// 		&i.BtevModifiedDate,
// 		&i.BtevBatchID,
// 		&i.BtevTraineeEntityID,
// 	)
// 	return i, err
// }

// const createBatchTraineeEvaluation = `-- name: CreateBatchTraineeEvaluation :one
// INSERT INTO bootcamp.batch_trainee_evaluation
// (btev_id, btev_type, btev_header, btev_section, btev_skill, btev_week, btev_skor, btev_note, btev_modified_date, btev_batch_id, btev_trainee_entity_id)
// VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
// RETURNING btev_id
// `

// type CreateBatchTraineeEvaluationParams struct {
// 	BtevID              int32     `db:"btev_id" json:"btevId"`
// 	BtevType            string    `db:"btev_type" json:"btevType"`
// 	BtevHeader          string    `db:"btev_header" json:"btevHeader"`
// 	BtevSection         string    `db:"btev_section" json:"btevSection"`
// 	BtevSkill           string    `db:"btev_skill" json:"btevSkill"`
// 	BtevWeek            int32     `db:"btev_week" json:"btevWeek"`
// 	BtevSkor            int32     `db:"btev_skor" json:"btevSkor"`
// 	BtevNote            string    `db:"btev_note" json:"btevNote"`
// 	BtevModifiedDate    time.Time `db:"btev_modified_date" json:"btevModifiedDate"`
// 	BtevBatchID         int32     `db:"btev_batch_id" json:"btevBatchId"`
// 	BtevTraineeEntityID int32     `db:"btev_trainee_entity_id" json:"btevTraineeEntityId"`
// }

// func (q *Queries) CreateBatchTraineeEvaluation(ctx context.Context, arg CreateBatchTraineeEvaluationParams) (*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {
// 	row := q.db.QueryRowContext(ctx, createBatchTraineeEvaluation,
// 		arg.BtevID,
// 		arg.BtevType,
// 		arg.BtevHeader,
// 		arg.BtevSection,
// 		arg.BtevSkill,
// 		arg.BtevWeek,
// 		arg.BtevSkor,
// 		arg.BtevNote,
// 		arg.BtevModifiedDate,
// 		arg.BtevBatchID,
// 		arg.BtevTraineeEntityID,
// 	)
// 	i := models.BootcampBatchTraineeEvaluation{}
// 	err := row.Scan(
// 		&i.BtevID,
// 		&i.BtevType,
// 		&i.BtevHeader,
// 		&i.BtevSection,
// 		&i.BtevSkill,
// 		&i.BtevWeek,
// 		&i.BtevSkor,
// 		&i.BtevNote,
// 		&i.BtevModifiedDate,
// 		&i.BtevBatchID,
// 		&i.BtevTraineeEntityID,
// 	)
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return &models.BootcampBatchTraineeEvaluation{
// 		BtevID:              i.BtevID,
// 		BtevType:            i.BtevType,
// 		BtevHeader:          i.BtevHeader,
// 		BtevSection:         i.BtevSection,
// 		BtevSkill:           i.BtevSkill,
// 		BtevWeek:            i.BtevWeek,
// 		BtevSkor:            i.BtevSkor,
// 		BtevNote:            i.BtevNote,
// 		BtevModifiedDate:    i.BtevModifiedDate,
// 		BtevBatchID:         i.BtevBatchID,
// 		BtevTraineeEntityID: i.BtevTraineeEntityID,
// 	}, nil
// }

// const updateBatchTraineeEvaluation = `-- name: UpdateBatchTraineeEvaluation :exec
// UPDATE bootcamp.batch_trainee_evaluation
// SET btev_type = $2,
//     btev_header = $3
// WHERE btev_id = $1
// `

// type UpdateBatchTraineeEvaluationParams struct {
// 	BtevID     int32  `db:"btev_id" json:"btevId"`
// 	BtevType   string `db:"btev_type" json:"btevType"`
// 	BtevHeader string `db:"btev_header" json:"btevHeader"`
// }

// func (q *Queries) UpdateBatchTraineeEvaluation(ctx context.Context, arg CreateBatchTraineeEvaluationParams) error {
// 	_, err := q.db.ExecContext(ctx, updateBatchTraineeEvaluation, arg.BtevID, arg.BtevType, arg.BtevHeader)
// 	return err
// }

// const deleteBatchTraineeEvaluation = `-- name: DeleteBatchTraineeEvaluation :exec
// DELETE FROM bootcamp.batch_trainee_evaluation
// WHERE btev_id = $1
// `

// func (q *Queries) DeleteBatchTraineeEvaluation(ctx context.Context, btevID int32) error {
// 	_, err := q.db.ExecContext(ctx, deleteBatchTraineeEvaluation, btevID)
// 	return err
// }
