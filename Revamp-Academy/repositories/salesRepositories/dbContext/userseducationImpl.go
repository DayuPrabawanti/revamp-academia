package dbContext

import (
	"context"
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
)

const createEducations = `-- name: CreateEducation :one
INSERT INTO users.users_education
(usdu_id, usdu_entity_id, usdu_school, usdu_degree, usdu_field_study,
usdu_graduate_year, usdu_start_date, usdu_end_date, usdu_grade,
usdu_activities, usdu_description, usdu_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
RETURNING *
`

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
		UsduModifiedDate: i.UsduModifiedDate,
	}, nil
}

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
