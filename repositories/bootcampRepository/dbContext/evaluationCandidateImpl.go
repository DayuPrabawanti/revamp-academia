package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

type BootcampEvaluationCandidate struct {
	UserEntityID        int32     `db:"user_entity_id" json:"userEntityId"`
	UserFullname        string    `json:"fullname"`
	UserPhoto           string    `db:"user_photo" json:"userPhoto"`
	UsduSchool          string    `db:"usdu_school" json:"usduSchool"`
	UsduFieldStudy      string    `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduGrade           string    `db:"usdu_grade" json:"usduGrade"`
	ProgTitle           string    `db:"prog_title" json:"progTitle"`
	BatchID             int32     `db:"batch_id" json:"batchId"`
	BatchName           string    `db:"batch_name" json:"batchName"`
	BatchStartDate      time.Time `db:"batch_start_date" json:"batchStartDate"`
	BatchEndDate        time.Time `db:"batch_end_date" json:"batchEndDate"`
	BatrStatus          string    `db:"batr_status" json:"batrStatus"`
	BtevBatchID         *int32    `db:"btev_batch_id" json:"btevBatchId"`
	BtevSkor            int32     `db:"total_skor" json:"btevSkor"`
	BtevType            string    `db:"btev_type" json:"btevType"`
	BtevHeader          string    `db:"btev_header" json:"btevHeader"`
	BtevSkill           string    `db:"btev_skill" json:"btevSkill"`
	BtevTraineeEntityID int32     `db:"btev_trainee_entity_id" json:"btevTraineeEntityId"`
}

const getEvaluationCandidates = `-- name: GetEvaluationCandidate :many
SELECT 
    us.user_entity_id, 
    CONCAT (us.user_first_name, ' ', us.user_last_name) AS fullname, 
    us.user_photo,  
    ud.usdu_school, 
    ud.usdu_field_study, 
    ud.usdu_grade,
    c.prog_title, 
    b.batch_id, 
    b.batch_name, 
	batch_start_date,
	batch_end_date,
    bt.batr_status, 
	bte.btev_batch_id,
    SUM(bte.btev_skor) AS total_skor,
	bte.btev_type,
	bte.btev_header,
	bte.btev_skill
FROM 
    bootcamp.batch_trainee_evaluation as bte
JOIN 
    bootcamp.batch as b
ON 
    btev_trainee_entity_id = batch_entity_id
JOIN 
    bootcamp.batch_trainee as bt
ON 
    batch_entity_id = batr_trainee_entity_id
JOIN 
    users.users as us
ON 
    batr_trainee_entity_id = user_entity_id
JOIN 
    users.users_education as ud
ON 
    user_entity_id = usdu_entity_id
JOIN 
    curriculum.program_entity as c
ON 
    usdu_entity_id = prog_entity_id
WHERE 
    user_entity_id = $1
GROUP BY
    us.user_entity_id, 
    fullname, 
    us.user_photo, 
    ud.usdu_school, 
    ud.usdu_field_study, 
    ud.usdu_grade,
    c.prog_title, 
    b.batch_id, 
    b.batch_name, 
	batch_start_date,
	batch_end_date,
    bt.batr_status,
	bte.btev_batch_id,
	bte.btev_type,
	bte.btev_header,
	bte.btev_skill
ORDER BY
	btev_batch_id;
`

func (q *Queries) GetEvaluationCandidate(ctx context.Context, userEntityID int32) ([]BootcampEvaluationCandidate, error) {
	rows, err := q.db.QueryContext(ctx, getEvaluationCandidates, userEntityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var evaluationCandidate []BootcampEvaluationCandidate

	for rows.Next() {
		var i BootcampEvaluationCandidate
		err := rows.Scan(
			&i.UserEntityID,
			&i.UserFullname,
			&i.UserPhoto,
			&i.UsduSchool,
			&i.UsduFieldStudy,
			&i.UsduGrade,
			&i.ProgTitle,
			&i.BatchID,
			&i.BatchName,
			&i.BatchStartDate,
			&i.BatchEndDate,
			&i.BatrStatus,
			&i.BtevBatchID,
			&i.BtevSkor,
			&i.BtevType,
			&i.BtevHeader,
			&i.BtevSkill,
		)

		if err != nil {
			return nil, err
		}
		evaluationCandidate = append(evaluationCandidate, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return evaluationCandidate, nil
}

const createEvaluationCandidates = `-- name: CreateEvaluationCandidate:one
INSERT INTO 
	bootcamp.batch_trainee_evaluation (
		btev_trainee_entity_id,
		btev_type,
		btev_header,
		btev_skill,
		btev_skor
	)
VALUES ($1, $2, $3, $4, $5)
RETURNING 
	(SELECT CONCAT (user_first_name, ' ', user_last_name) FROM users.users WHERE user_entity_id = $1) AS fullname;
`

type CreateEvaluationCandidateParams struct {
	BtevTraineeEntityID int32  `db:"btev_trainee_entity_id" json:"btevTraineeEntityId"`
	BtevType            string `db:"btev_type" json:"btevType"`
	BtevHeader          string `db:"btev_header" json:"btevHeader"`
	BtevSkill           string `db:"btev_skill" json:"btevSkill"`
	BtevSkor            int32  `db:"btev_skor" json:"btevSkor"`
}

func (q *Queries) CreateEvaluationCandidate(ctx context.Context, arg CreateEvaluationCandidateParams) (*BootcampEvaluationCandidate, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEvaluationCandidates,
		arg.BtevTraineeEntityID,
		arg.BtevType,
		arg.BtevHeader,
		arg.BtevSkill,
		arg.BtevSkor,
	)

	i := BootcampEvaluationCandidate{}
	err := row.Scan(
		&i.UserFullname,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &BootcampEvaluationCandidate{
		UserFullname: i.UserFullname,
		BtevType:     i.BtevType,
		BtevHeader:   i.BtevHeader,
		BtevSkill:    i.BtevSkill,
		BtevSkor:     i.BtevSkor,
	}, nil
}

const updateEvaluationCandidates = `-- name: UpdateEvaluationCandidate :exec
UPDATE bootcamp.batch_trainee_evaluation
SET btev_skor = $1
WHERE btev_trainee_entity_id = $2
RETURNING
    (SELECT CONCAT (user_first_name, ' ', user_last_name) FROM users.users WHERE user_entity_id = $2) AS fullname,
    btev_type,
    btev_header,
    btev_skill;
`

type UpdateEvaluationCandidateParams struct {
	BtevSkor            int32 `db:"total_skor" json:"btevSkor"`
	BtevTraineeEntityID int32 `db:"btev_trainee_entity_id" json:"btevTraineeEntityId"`
}

func (q *Queries) UpdateEvaluationCandidate(ctx context.Context, arg UpdateEvaluationCandidateParams) (*BootcampEvaluationCandidate, error) {
	row := q.db.QueryRowContext(ctx, updateEvaluationCandidates,
		arg.BtevSkor,
		arg.BtevTraineeEntityID)
	var i BootcampEvaluationCandidate
	err := row.Scan(
		&i.UserFullname,
		&i.BtevType,
		&i.BtevHeader,
		&i.BtevSkill,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
