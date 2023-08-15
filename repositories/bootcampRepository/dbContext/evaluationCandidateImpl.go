package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listEvaluationCandidates = `-- name: Group :many
SELECT 
    b.batch_name,  
    pe.prog_title,
    us.user_photo,
    us.user_name,
    b.batch_status,
    b.batch_start_date,
    b.batch_end_date,
	use.usdu_school,
	use.usdu_field_study,
	use.usdu_grade,
    bte.btev_skor
FROM 
    bootcamp.batch b
JOIN 
    curriculum.program_entity pe ON b.batch_entity_id = pe.prog_entity_id
JOIN 
    users.users us ON us.user_entity_id = b.batch_entity_id
JOIN 
    users.education use ON use.usdu = b.batch_entity_id
JOIN 
    bootcamp.batch_trainee_evaluation bte ON b.batch_id = bte.btev_batch_id
ORDER BY
    b.batch_name;
`

func (q *Queries) ListEvaluationCandidate(ctx context.Context) ([]models.EvaluationCandidateMockup, error) {
	rows, err := q.db.QueryContext(ctx, listEvaluationCandidates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.EvaluationCandidateMockup
	for rows.Next() {
		var i models.EvaluationCandidateMockup
		if err := rows.Scan(
			&i.BootcampBatch.BatchName,
			&i.CurriculumProgramEntity.ProgTitle,
			&i.UsersUser.UserPhoto,
			&i.UsersUser.UserName,
			&i.BootcampBatch.BatchStatus,
			&i.BootcampBatchTraineeEvaluation.BtevSkor,
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
