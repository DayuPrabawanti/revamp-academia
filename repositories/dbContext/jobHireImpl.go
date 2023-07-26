package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

// ------------------------------   JOB CATEGORY -----------------------------

const GetJobCategoryImpl = `-- name: GetJobCategoryImpl :one
SELECT Joca_ID, Joca_Name, Joca_Modified_Date FROM jobhire.job_category
WHERE Joca_ID = $1
`

func (q *Queries) GetJobCategoryImpl(ctx context.Context, jobCategoryID int16) (models.JobhireJobCategory, error) {
	row := q.db.QueryRowContext(ctx, GetJobCategoryImpl, jobCategoryID)
	var i models.JobhireJobCategory
	err := row.Scan(
		&i.JocaID,
		&i.JocaName,
		&i.JocaModifiedDate,
	)
	return i, err
}

const ListJobCategoryImpl = `-- name: ListJobCategoryImpl :many
SELECT Joca_ID, Joca_Name, Joca_Modified_Date FROM jobhire.job_category
ORDER BY Joca_ID
`

func (q *Queries) ListJobCategoryImpl(ctx context.Context) ([]models.JobhireJobCategory, error) {
	rows, err := q.db.QueryContext(ctx, ListJobCategoryImpl)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.JobhireJobCategory
	for rows.Next() {
		var i models.JobhireJobCategory
		if err := rows.Scan(
			&i.JocaID,
			&i.JocaName,
			&i.JocaModifiedDate,
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

const CreateJobCategoryImpl = `-- name: CreateJobCategoryImpl :one
INSERT INTO 
	jobhire.job_category (
		joca_id,
		joca_name,
		joca_modified_date
	)
VALUES ($1,$2,$3) RETURNING *
`

func (q *Queries) CreateJobCategoryImpl(ctx context.Context, arg CreateJobCategoryParams) (*models.JobhireJobCategory, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, CreateJobCategoryImpl,
		arg.JocaID,
		arg.JocaName,
		arg.JocaModifiedDate,
	)
	i := models.JobhireJobCategory{}
	err := row.Scan(
		&i.JocaID,
		&i.JocaName,
		&i.JocaModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.JobhireJobCategory{
		JocaID:           i.JocaID,
		JocaName:         i.JocaName,
		JocaModifiedDate: i.JocaModifiedDate,
	}, nil
}

type CreateJobCategoryParams struct {
	JocaID           int32     `db:"joca_id" json:"jocaId"`
	JocaName         string    `db:"joca_name" json:"jocaName"`
	JocaModifiedDate time.Time `db:"joca_modified_date" json:"jocaModifiedDate"`
}

const UpdateJobCategoryImpl = `-- name: UpdateJobCategoryImpl :exec
UPDATE 
	jobhire.job_category
  		set Joca_name = $2,
  		Joca_Modified_Date = $3
WHERE Joca_id = $1
`

func (q *Queries) UpdateJobCategoryImpl(ctx context.Context, arg CreateJobCategoryParams) error {
	_, err := q.db.ExecContext(ctx, UpdateJobCategoryImpl, arg.JocaID, arg.JocaName, arg.JocaModifiedDate)
	return err
}

const DeleteJobCategoryImpl = `-- name: DeleteJobCategoryImpl :exec
DELETE FROM jobhire.job_category
  WHERE joca_id = $1
`

func (q *Queries) DeleteJobCategoryImpl(ctx context.Context, jocaID int32) error {
	_, err := q.db.ExecContext(ctx, DeleteJobCategoryImpl, jocaID)
	return err
}

// ------------------------------   JOB CLIENT -----------------------------

const GetJobClientImpl = `-- name: GetJobClientImpl :one
SELECT 
		clit_id, clit_name, clit_about, clit_modified_date, clit_addr_id, clit_emra_id 
		FROM jobhire.client
WHERE clit_id = $1
`

func (q *Queries) GetJobClientImpl(ctx context.Context, clitID int64) (models.JobhireClient, error) {
	row := q.db.QueryRowContext(ctx, GetJobClientImpl, clitID)
	var i models.JobhireClient
	err := row.Scan(
		&i.ClitID,
		&i.ClitName,
		&i.ClitAbout,
		&i.ClitModifiedDate,
		&i.ClitAddrID,
		&i.ClitEmraID,
	)
	return i, err
}

const ListJobClientImpl = `-- name: ListJobClientImpl :many
SELECT 
		clit_id, clit_name, clit_about, clit_modified_date, clit_addr_id, clit_emra_id
		FROM jobhire.client
ORDER BY clit_id
`

func (q *Queries) ListJobClientImpl(ctx context.Context) ([]models.JobhireClient, error) {
	rows, err := q.db.QueryContext(ctx, ListJobClientImpl)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items3 []models.JobhireClient
	for rows.Next() {
		var i models.JobhireClient
		if err := rows.Scan(
			&i.ClitID,
			&i.ClitName,
			&i.ClitAbout,
			&i.ClitModifiedDate,
			&i.ClitAddrID,
			&i.ClitEmraID,
		); err != nil {
			return nil, err
		}
		items3 = append(items3, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items3, nil
}

const CreateJobClientImpl = `-- name: CreateJobClientImpl :one
INSERT INTO 
	jobhire.client 
	(clit_id, clit_name, clit_about, clit_modified_date, clit_addr_id, clit_emra_id)
 	VALUES 
 	($1,$2,$3,$4,$5,$6)
	RETURNING *
`

type CreateJobClientParams struct {
	ClitID           int32     `db:"clit_id" json:"clitId"`
	ClitName         string    `db:"clit_name" json:"clitName"`
	ClitAbout        string    `db:"clit_about" json:"clitAbout"`
	ClitModifiedDate time.Time `db:"clit_modified_date" json:"clitModifiedDate"`
	ClitAddrID       int32     `db:"clit_addr_id" json:"clitAddrId"`
	ClitEmraID       int32     `db:"clit_emra_id" json:"clitEmraId"`
}

func (q *Queries) CreateJobClientImpl(ctx context.Context, arg CreateJobClientParams) (*models.JobhireClient, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, CreateJobClientImpl,
		arg.ClitID,
		arg.ClitName,
		arg.ClitAbout,
		arg.ClitModifiedDate,
		arg.ClitAddrID,
		arg.ClitEmraID,
	)
	i := models.JobhireClient{}
	err := row.Scan(
		&i.ClitID,
		&i.ClitName,
		&i.ClitAbout,
		&i.ClitModifiedDate,
		&i.ClitAddrID,
		&i.ClitEmraID,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.JobhireClient{
		ClitID:           i.ClitID,
		ClitName:         i.ClitName,
		ClitAbout:        i.ClitAbout,
		ClitModifiedDate: i.ClitModifiedDate,
		ClitAddrID:       i.ClitAddrID,
		ClitEmraID:       i.ClitEmraID,
	}, nil
}

const UpdateJobClientImpl = `-- name: UpdateJobClientImpl :exec
UPDATE 
	jobhire.client
  		set clit_name = $2,
  		clit_about = $3
WHERE clit_id = $1
`

func (q *Queries) UpdateJobClientImpl(ctx context.Context, arg CreateJobClientParams) error {
	_, err := q.db.ExecContext(ctx, UpdateJobClientImpl, arg.ClitID, arg.ClitName, arg.ClitAbout)
	return err
}

const DeleteJobClientImpl = `-- name: DeleteJobClientImpl :exec
DELETE FROM jobhire.client
WHERE clit_id = $1
`

func (q *Queries) DeleteJobClientImpl(ctx context.Context, clitID int32) error {
	_, err := q.db.ExecContext(ctx, DeleteJobClientImpl, clitID)
	return err
}

// ------------------------------   JOB POST  -----------------------------

const GetJobPostImpl = `-- name: GetJobPostImpl :one
SELECT 
	jopo_entity_id, jopo_number, jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, 
	jopo_max_salary, jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date, 
	jopo_modified_date, jopo_emp_entity_id, jopo_clit_id, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, 
	jopo_work_code, jopo_edu_code, jopo_indu_code, jopo_status FROM jobHire.job_post
WHERE jopo_entity_id = $1
`

func (q *Queries) GetJobPostImpl(ctx context.Context, jopoEntityID int32) (models.JobhireJobPost, error) {
	row := q.db.QueryRowContext(ctx, GetJobPostImpl, jopoEntityID)
	var i models.JobhireJobPost
	err := row.Scan(
		&i.JopoEntityID,
		&i.JopoNumber,
		&i.JopoTitle,
		&i.JopoStartDate,
		&i.JopoEndDate,
		&i.JopoMinSalary,
		&i.JopoMaxSalary,
		&i.JopoMinExperience,
		&i.JopoMaxExperience,
		&i.JopoPrimarySkill,
		&i.JopoSecondarySkill,
		&i.JopoPublishDate,
		&i.JopoModifiedDate,
		&i.JopoEmpEntityID,
		&i.JopoClitID,
		&i.JopoJoroID,
		&i.JopoJotyID,
		&i.JopoJocaID,
		&i.JopoAddrID,
		&i.JopoWorkCode,
		&i.JopoEduCode,
		&i.JopoInduCode,
		&i.JopoStatus,
	)
	return i, err
}

const ListJobPostImpl = `-- name: ListJobPostImpl :many
SELECT 
	jopo_entity_id, jopo_number, jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, 
	jopo_max_salary, jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date,
	jopo_modified_date, jopo_emp_entity_id, jopo_clit_id, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, jopo_work_code,
	jopo_edu_code, jopo_indu_code, jopo_status FROM jobHire.job_post
ORDER BY jopo_title
`

func (q *Queries) ListJobPostImpl(ctx context.Context) ([]models.JobhireJobPost, error) {
	rows, err := q.db.QueryContext(ctx, ListJobPostImpl)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.JobhireJobPost
	for rows.Next() {
		var i models.JobhireJobPost
		if err := rows.Scan(
			&i.JopoEntityID,
			&i.JopoNumber,
			&i.JopoTitle,
			&i.JopoStartDate,
			&i.JopoEndDate,
			&i.JopoMinSalary,
			&i.JopoMaxSalary,
			&i.JopoMinExperience,
			&i.JopoMaxExperience,
			&i.JopoPrimarySkill,
			&i.JopoSecondarySkill,
			&i.JopoPublishDate,
			&i.JopoModifiedDate,
			&i.JopoEmpEntityID,
			&i.JopoClitID,
			&i.JopoJoroID,
			&i.JopoJotyID,
			&i.JopoJocaID,
			&i.JopoAddrID,
			&i.JopoWorkCode,
			&i.JopoEduCode,
			&i.JopoInduCode,
			&i.JopoStatus,
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

const CreateJobPostImpl = `-- name: CreateJobPostImpl :one
INSERT INTO jobHire.job_post
	(jopo_entity_id, jopo_number, jopo_title, jopo_start_date, jopo_end_date,
	jopo_min_salary, jopo_max_salary, jopo_min_experience, jopo_max_experience, jopo_primary_skill,
	jopo_secondary_skill, jopo_publish_date, jopo_modified_date, jopo_emp_entity_id, jopo_clit_id, 
	jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id,jopo_work_code, jopo_edu_code, jopo_indu_code,
	jopo_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,$21,$22,$23)
RETURNING *
`

type CreateJobPostParams struct {
	JopoEntityID       int32     `db:"jopo_entity_id" json:"jopoEntityId"`
	JopoNumber         string    `db:"jopo_number" json:"jopoNumber"`
	JopoTitle          string    `db:"jopo_title" json:"jopoTitle"`
	JopoStartDate      time.Time `db:"jopo_start_date" json:"jopoStartDate"`
	JopoEndDate        int32     `db:"jopo_end_date" json:"jopoEndDate"`
	JopoMinSalary      int32     `db:"jopo_min_salary" json:"jopoMinSalary"`
	JopoMaxSalary      int32     `db:"jopo_max_salary" json:"jopoMaxSalary"`
	JopoMinExperience  int32     `db:"jopo_min_experience" json:"jopoMinExperience"`
	JopoMaxExperience  int32     `db:"jopo_max_experience" json:"jopoMaxExperience"`
	JopoPrimarySkill   string    `db:"jopo_primary_skill" json:"jopoPrimarySkill"`
	JopoSecondarySkill string    `db:"jopo_secondary_skill" json:"jopoSecondarySkill"`
	JopoPublishDate    time.Time `db:"jopo_publish_date" json:"jopoPublishDate"`
	JopoModifiedDate   time.Time `db:"jopo_modified_date" json:"jopoModifiedDate"`
	JopoEmpEntityID    int32     `db:"jopo_emp_entity_id" json:"jopoEmpEntityId"`
	JopoClitID         int32     `db:"jopo_clit_id" json:"jopoClitId"`
	JopoJoroID         int32     `db:"jopo_joro_id" json:"jopoJoroId"`
	JopoJotyID         int32     `db:"jopo_joty_id" json:"jopoJotyId"`
	JopoJocaID         int32     `db:"jopo_joca_id" json:"jopoJocaId"`
	JopoAddrID         int32     `db:"jopo_addr_id" json:"jopoAddrId"`
	JopoWorkCode       string
	JopoEduCode        string
	JopoInduCode       string
	JopoStatus         string `db:"jopo_status" json:"jopoStatus"`
}

func (q *Queries) CreateJobPostImpl(ctx context.Context, arg CreateJobPostParams) (*models.JobhireJobPost, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, CreateJobPostImpl,
		arg.JopoEntityID,
		arg.JopoNumber,
		arg.JopoTitle,
		arg.JopoStartDate,
		arg.JopoEndDate,
		arg.JopoMinSalary,
		arg.JopoMaxSalary,
		arg.JopoMinExperience,
		arg.JopoMaxExperience,
		arg.JopoPrimarySkill,
		arg.JopoSecondarySkill,
		arg.JopoPublishDate,
		arg.JopoModifiedDate,
		arg.JopoEmpEntityID,
		arg.JopoClitID,
		arg.JopoJoroID,
		arg.JopoJotyID,
		arg.JopoJocaID,
		arg.JopoAddrID,
		arg.JopoWorkCode,
		arg.JopoEduCode,
		arg.JopoInduCode,
		arg.JopoStatus,
	)

	i := models.JobhireJobPost{}
	err := row.Scan(
		&i.JopoEntityID,
		&i.JopoNumber,
		&i.JopoTitle,
		&i.JopoStartDate,
		&i.JopoEndDate,
		&i.JopoMinSalary,
		&i.JopoMaxSalary,
		&i.JopoMinExperience,
		&i.JopoMaxExperience,
		&i.JopoPrimarySkill,
		&i.JopoSecondarySkill,
		&i.JopoPublishDate,
		&i.JopoModifiedDate,
		&i.JopoEmpEntityID,
		&i.JopoClitID,
		&i.JopoJoroID,
		&i.JopoJotyID,
		&i.JopoJocaID,
		&i.JopoAddrID,
		&i.JopoWorkCode,
		&i.JopoEduCode,
		&i.JopoInduCode,
		&i.JopoStatus,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.JobhireJobPost{
		JopoEntityID:       i.JopoEntityID,
		JopoNumber:         i.JopoNumber,
		JopoTitle:          i.JopoTitle,
		JopoStartDate:      i.JopoStartDate,
		JopoEndDate:        i.JopoEndDate,
		JopoMinSalary:      i.JopoMinSalary,
		JopoMaxSalary:      i.JopoMaxSalary,
		JopoMinExperience:  i.JopoMinExperience,
		JopoMaxExperience:  i.JopoMaxExperience,
		JopoPrimarySkill:   i.JopoPrimarySkill,
		JopoSecondarySkill: i.JopoSecondarySkill,
		JopoPublishDate:    i.JopoPublishDate,
		JopoModifiedDate:   i.JopoModifiedDate,
		JopoEmpEntityID:    i.JopoEmpEntityID,
		JopoClitID:         i.JopoClitID,
		JopoJoroID:         i.JopoJoroID,
		JopoJotyID:         i.JopoJotyID,
		JopoJocaID:         i.JopoJocaID,
		JopoAddrID:         i.JopoAddrID,
		JopoWorkCode:       i.JopoWorkCode,
		JopoEduCode:        i.JopoEduCode,
		JopoInduCode:       i.JopoInduCode,
		JopoStatus:         i.JopoStatus,
	}, nil
}

const UpdateJobPostImpl = `-- name: UpdateJobPostImpl :exec
UPDATE 
	jobHire.job_post
  		set jopo_title = $2,
		jopo_min_salary = $3,
		jopo_max_salary = $4,
		jopo_start_date = $5,
		jopo_end_date = $6,
		jopo_min_experience = $7
WHERE jopo_entity_id = $1
`

func (q *Queries) UpdateJobPostImpl(ctx context.Context, arg CreateJobPostParams) error {
	_, err := q.db.ExecContext(ctx, UpdateJobPostImpl, arg.JopoEntityID, arg.JopoTitle, arg.JopoMinSalary, arg.JopoMaxSalary, arg.JopoStartDate, arg.JopoEndDate, arg.JopoMinExperience)
	return err
}

const DeleteJobPostImpl = `-- name: DeleteJobPostImpl :exec
DELETE FROM jobHire.job_post
WHERE jopo_entity_id = $1
`

func (q *Queries) DeleteJobPostImpl(ctx context.Context, jopoEntityID int32) error {
	_, err := q.db.ExecContext(ctx, DeleteJobPostImpl, jopoEntityID)
	return err
}
