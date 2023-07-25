package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

type CreateEducationParams struct {
	UsduID           int32          `db:"usdu_id" json:"usduId"`
	UsduEntityID     int32          `db:"usdu_entity_id" json:"usduEntityId"`
	UsduSchool       sql.NullString `db:"usdu_school" json:"usduSchool"`
	UsduDegree       sql.NullString `db:"usdu_degree" json:"usduDegree"`
	UsduFieldStudy   sql.NullString `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduGraduateYear sql.NullString `db:"usdu_graduate_year" json:"usduGraduateYear"`
	UsduStartDate    sql.NullTime   `db:"usdu_start_date" json:"usduStartDate"`
	UsduEndDate      sql.NullTime   `db:"usdu_end_date" json:"usduEndDate"`
	UsduGrade        sql.NullString `db:"usdu_grade" json:"usduGrade"`
	UsduActivities   sql.NullString `db:"usdu_activities" json:"usduActivities"`
	UsduDescription  sql.NullString `db:"usdu_description" json:"usduDescription"`
	UsduModifiedDate sql.NullTime   `db:"usdu_modified_date" json:"usduModifiedDate"`
}

const listEducation = `-- name: ListEducation :many
SELECT usdu_id, usdu_entity_id, usdu_school, usdu_degree, usdu_field_study, usdu_graduate_year, usdu_start_date, usdu_end_date, usdu_grade, usdu_activities, usdu_description, usdu_modified_date FROM users.users_education
ORDER BY usdu_entity_id
`

func (q *Queries) ListEducation(ctx context.Context) ([]models.UsersUsersEducation, error) {
	rows, err := q.db.QueryContext(ctx, listEducation)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUsersEducation
	for rows.Next() {
		var i models.UsersUsersEducation
		if err := rows.Scan(
			&i.UsduID,
			&i.UsduEntityID,
			&i.UsduSchool,
			&i.UsduDegree,
			&i.UsduFieldStudy,
			&i.UsduGraduateYear,
			&i.UsduStartDate,
			&i.UsduEndDate,
			&i.UsduGrade,
			&i.UsduActivities,
			&i.UsduDescription,
			&i.UsduModifiedDate,
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

// Get Users Education
const getEducation = `-- name: GetEducation :one

SELECT usdu_id, usdu_entity_id, usdu_school, usdu_degree, usdu_field_study, usdu_graduate_year, usdu_start_date, usdu_end_date, usdu_grade, usdu_activities, usdu_description, usdu_modified_date FROM users.users_education
WHERE usdu_id = $1
`

// Users Education
func (q *Queries) GetEducation(ctx context.Context, usduID int32) (models.UsersUsersEducation, error) {
	row := q.db.QueryRowContext(ctx, getEducation, usduID)
	var i models.UsersUsersEducation
	err := row.Scan(
		&i.UsduID,
		&i.UsduEntityID,
		&i.UsduSchool,
		&i.UsduDegree,
		&i.UsduFieldStudy,
		&i.UsduGraduateYear,
		&i.UsduStartDate,
		&i.UsduEndDate,
		&i.UsduGrade,
		&i.UsduActivities,
		&i.UsduDescription,
		&i.UsduModifiedDate,
	)
	return i, err
}

// Create User Education
const createEducation = `-- name: CreateEducation :one

INSERT INTO users.users_education
(usdu_id, usdu_entity_id, usdu_school, usdu_degree, usdu_field_study,
usdu_graduate_year, usdu_start_date, usdu_end_date, usdu_grade,
usdu_activities, usdu_description, usdu_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
RETURNING *
`

func (q *Queries) CreateEducation(ctx context.Context, arg CreateEducationParams) (*models.UsersUsersEducation, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEducation,
		arg.UsduID,
		arg.UsduEntityID,
		arg.UsduSchool,
		arg.UsduDegree,
		arg.UsduFieldStudy,
		arg.UsduGraduateYear,
		arg.UsduStartDate,
		arg.UsduEndDate,
		arg.UsduGrade,
		arg.UsduActivities,
		arg.UsduDescription,
		arg.UsduModifiedDate,
	)
	i := models.UsersUsersEducation{}
	err := row.Scan(
		&i.UsduID,
		&i.UsduEntityID,
		&i.UsduSchool,
		&i.UsduDegree,
		&i.UsduFieldStudy,
		&i.UsduGraduateYear,
		&i.UsduStartDate,
		&i.UsduEndDate,
		&i.UsduGrade,
		&i.UsduActivities,
		&i.UsduDescription,
		&i.UsduModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUsersEducation{
		UsduID:           i.UsduID,
		UsduEntityID:     i.UsduEntityID,
		UsduSchool:       i.UsduSchool,
		UsduDegree:       i.UsduDegree,
		UsduFieldStudy:   i.UsduFieldStudy,
		UsduGraduateYear: i.UsduGraduateYear,
		UsduStartDate:    i.UsduStartDate,
		UsduEndDate:      i.UsduEndDate,
		UsduGrade:        i.UsduGrade,
		UsduActivities:   i.UsduActivities,
		UsduDescription:  i.UsduDescription,
		UsduModifiedDate: sql.NullTime{Time: time.Now(), Valid: true},
	}, nil
}

// Update User Education
const updateEducation = `-- name: UpdateEducation :exec
UPDATE users.users_education
  set usdu_entity_id = $2,
  usdu_school = $3,
  usdu_degree = $4,
  usdu_field_study = $5,
  usdu_graduate_year = $6,
  usdu_start_date = $7,
  usdu_end_date = $8,
  usdu_grade = $9,
  usdu_activities = $10,
  usdu_description = $11,
  usdu_modified_date = $12
WHERE usdu_id = $1
`

func (q *Queries) UpdateEducation(ctx context.Context, arg CreateEducationParams) error {
	_, err := q.db.ExecContext(ctx, updateEducation,
		arg.UsduID,
		arg.UsduEntityID,
		arg.UsduSchool,
		arg.UsduDegree,
		arg.UsduFieldStudy,
		arg.UsduGraduateYear,
		arg.UsduStartDate,
		arg.UsduEndDate,
		arg.UsduGrade,
		arg.UsduActivities,
		arg.UsduDescription,
		arg.UsduModifiedDate,
	)
	return err
}

// Delete User Education
const deleteEducation = `-- name: DeleteEducation :exec
DELETE FROM users.users_education
WHERE usdu_id = $1
`

func (q *Queries) DeleteEducation(ctx context.Context, usduID int32) error {
	_, err := q.db.ExecContext(ctx, deleteEducation, usduID)
	return err
}
