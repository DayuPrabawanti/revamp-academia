package dbContext

import (
	"context"
	"database/sql"
	"time"

	"codeid.revampacademy/models"
)

type UsersUsersEducation struct {
	UsduID           int32        `db:"usdu_id" json:"usduId"`
	UsduEntityID     int32        `db:"usdu_entity_id" json:"usduEntityId"`
	UsduSchool       string       `db:"usdu_school" json:"usduSchool"`
	UsduDegree       string       `db:"usdu_degree" json:"usduDegree"`
	UsduFieldStudy   string       `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduGraduateYear string       `db:"usdu_graduate_year" json:"usduGraduateYear"`
	UsduStartDate    sql.NullTime `db:"usdu_start_date" json:"usduStartDate"`
	UsduEndDate      sql.NullTime `db:"usdu_end_date" json:"usduEndDate"`
	UsduGrade        string       `db:"usdu_grade" json:"usduGrade"`
	UsduActivities   string       `db:"usdu_activities" json:"usduActivities"`
	UsduDescription  string       `db:"usdu_description" json:"usduDescription"`
	UsduModifiedDate sql.NullTime `db:"usdu_modified_date" json:"usduModifiedDate"`
}

const ListUsersEducation = `-- name:ListUsersEducation :many
SELECT UsduID, UsduEntityID, UsduSchool, UsduDegree,UsduFieldStudy,	UsduGraduateYear,UsduStartDate,UsduEndDate,UsduGrade,UsduActivities,UsduDescription,UsduModifiedDate
FROM users.users_education
ORDER BY UsduID
`

func (q *Queries) ListUsersEducation(ctx context.Context) ([]models.UsersUsersEducation, error) {
	rows, err := q.db.QueryContext(ctx, ListUsersEducation)
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

const UpdateEducation = `-- name: UpdateEducation :exec
UPDATE users.users_education
set UsduEntityID = $2,
UsduSchool   = $3,
UsduDegree = $4,
UsduFieldStudy = $5,
UsduGraduateYear = $6,
UsduStartDate = $7,
UsduEndDate = $8,
UsduGrade = $9,
UsduActivities = $10,
UsduDescription = $11,
UsduModifiedDate = $12,
WHERE UsduID = $1`

func (q *Queries) UpdateEducation(ctx context.Context, arg UsersUsersEducation) error {
	_, err := q.db.ExecContext(ctx, UpdateEducation,
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
		sql.NullTime{Time: time.Now(), Valid: true},
	)
	return err
}

const GetEducation = `-- name: GetEducation :one
SELECT UsduID, UsduEntityID, UsduSchool, UsduDegree,UsduFieldStudy,	UsduGraduateYear,UsduStartDate,UsduEndDate,UsduGrade,UsduActivities,UsduDescription,UsduModifiedDate
, picture FROM UsduID
WHERE UsduEntityID = $1
`

func (q *Queries) GetEducation(ctx context.Context, usduId int32) (models.UsersUsersEducation, error) {
	row := q.db.QueryRowContext(ctx, GetEducation, usduId)
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
