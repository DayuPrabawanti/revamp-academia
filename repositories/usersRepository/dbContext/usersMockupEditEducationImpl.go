package dbContext

import (
	"context"
	"time"
)

type AddEducationParams struct {
	UsduSchool       string `db:"usdu_school" json:"usduSchool"`
	UsduDegree       string `db:"usdu_degree" json:"usduDegree"`
	UsduFieldStudy   string `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduGrade        string `db:"usdu_grade" json:"usduGrade"`
	UsduStartDate    time.Time `db:"usdu_start_date" json:"usduStartDate"`
	UsduEndDate      time.Time `db:"usdu_end_date" json:"usduEndDate"`
	UsduGraduateYear string `db:"usdu_graduate_year" json:"usduGraduateYear"`
	UsduActivities   string `db:"usdu_activities" json:"usduActivities"`
	UsduDescription  string `db:"usdu_description" json:"usduDescription"`
}

const addEducation = `-- name: addEducation :one

INSERT INTO users.users_education(
	usdu_entity_id, usdu_school, usdu_degree, usdu_field_study, usdu_grade, 
	usdu_start_date, usdu_end_date, usdu_graduate_year, usdu_activities, usdu_description,
	usdu_modified_date
 )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, Now())
`

func (q *Queries) AddEducation(ctx context.Context, arg AddEducationParams, id int32) error {
	_, err := q.db.ExecContext(ctx, addEducation, id,
	arg.UsduSchool,
	arg.UsduDegree,
	arg.UsduFieldStudy,
	arg.UsduGrade,
	arg.UsduStartDate,
	arg.UsduEndDate,
	arg.UsduGraduateYear,
	arg.UsduActivities,
	arg.UsduDescription,
	)
	return err
}