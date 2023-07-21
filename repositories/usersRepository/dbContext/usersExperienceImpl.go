package dbContext

import (
	"context"
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
)

type CreateExperienceParams struct {
	UsexID              int32          `db:"usex_id" json:"usexId"`
	UsexEntityID        int32          `db:"usex_entity_id" json:"usexEntityId"`
	UsexTitle           sql.NullString `db:"usex_title" json:"usexTitle"`
	UsexProfileHeadline sql.NullString `db:"usex_profile_headline" json:"usexProfileHeadline"`
	UsexEmploymentType  sql.NullString `db:"usex_employment_type" json:"usexEmploymentType"`
	UsexCompanyName     sql.NullString `db:"usex_company_name" json:"usexCompanyName"`
	UsexIsCurrent       sql.NullString `db:"usex_is_current" json:"usexIsCurrent"`
	UsexStartDate       sql.NullTime   `db:"usex_start_date" json:"usexStartDate"`
	UsexEndDate         sql.NullTime   `db:"usex_end_date" json:"usexEndDate"`
	UsexIndustry        sql.NullString `db:"usex_industry" json:"usexIndustry"`
	UsexDescription     sql.NullString `db:"usex_description" json:"usexDescription"`
	UsexExperienceType  sql.NullString `db:"usex_experience_type" json:"usexExperienceType"`
	UsexCityID          sql.NullInt32  `db:"usex_city_id" json:"usexCityId"`
}

const listExperience = `-- name: ListExperience :many
SELECT usex_id, usex_entity_id, usex_title, usex_profile_headline, usex_employment_type, usex_company_name, usex_is_current, usex_start_date, usex_end_date, usex_industry, usex_description, usex_experience_type, usex_city_id FROM users.users_experiences
ORDER BY usex_title
`

func (q *Queries) ListExperience(ctx context.Context) ([]models.UsersUsersExperience, error) {
	rows, err := q.db.QueryContext(ctx, listExperience)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUsersExperience
	for rows.Next() {
		var i models.UsersUsersExperience
		if err := rows.Scan(
			&i.UsexID,
			&i.UsexEntityID,
			&i.UsexTitle,
			&i.UsexProfileHeadline,
			&i.UsexEmploymentType,
			&i.UsexCompanyName,
			&i.UsexIsCurrent,
			&i.UsexStartDate,
			&i.UsexEndDate,
			&i.UsexIndustry,
			&i.UsexDescription,
			&i.UsexExperienceType,
			&i.UsexCityID,
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

const getExperience = `-- name: GetExperience :one

SELECT usex_id, usex_entity_id, usex_title, usex_profile_headline, usex_employment_type, usex_company_name, usex_is_current, usex_start_date, usex_end_date, usex_industry, usex_description, usex_experience_type, usex_city_id FROM users.users_experiences
WHERE usex_id = $1
`

// Users Experience
func (q *Queries) GetExperience(ctx context.Context, usexID int32) (models.UsersUsersExperience, error) {
	row := q.db.QueryRowContext(ctx, getExperience, usexID)
	var i models.UsersUsersExperience
	err := row.Scan(
		&i.UsexID,
		&i.UsexEntityID,
		&i.UsexTitle,
		&i.UsexProfileHeadline,
		&i.UsexEmploymentType,
		&i.UsexCompanyName,
		&i.UsexIsCurrent,
		&i.UsexStartDate,
		&i.UsexEndDate,
		&i.UsexIndustry,
		&i.UsexDescription,
		&i.UsexExperienceType,
		&i.UsexCityID,
	)
	return i, err
}

const createExperience = `-- name: CreateExperience :one

INSERT INTO users.users_experiences
(usex_id, usex_entity_id, usex_title, usex_profile_headline, usex_employment_type,
usex_company_name, usex_is_current, usex_start_date, usex_end_date, usex_industry,
usex_description, usex_experience_type, usex_city_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
RETURNING *
`

func (q *Queries) CreateExperience(ctx context.Context, arg CreateExperienceParams) (*models.UsersUsersExperience, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createExperience,
		arg.UsexID,
		arg.UsexEntityID,
		arg.UsexTitle,
		arg.UsexProfileHeadline,
		arg.UsexEmploymentType,
		arg.UsexCompanyName,
		arg.UsexIsCurrent,
		arg.UsexStartDate,
		arg.UsexEndDate,
		arg.UsexIndustry,
		arg.UsexDescription,
		arg.UsexExperienceType,
		arg.UsexCityID,
	)
	i := models.UsersUsersExperience{}
	err := row.Scan(
		&i.UsexID,
		&i.UsexEntityID,
		&i.UsexTitle,
		&i.UsexProfileHeadline,
		&i.UsexEmploymentType,
		&i.UsexCompanyName,
		&i.UsexIsCurrent,
		&i.UsexStartDate,
		&i.UsexEndDate,
		&i.UsexIndustry,
		&i.UsexDescription,
		&i.UsexExperienceType,
		&i.UsexCityID,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUsersExperience{
		UsexID: i.UsexID,
		UsexEntityID: i.UsexEntityID,
		UsexTitle:    i.UsexTitle,
		UsexProfileHeadline:   i.UsexProfileHeadline,
		UsexEmploymentType:     i.UsexEmploymentType,
		UsexCompanyName:        i.UsexCompanyName,
		UsexIsCurrent:          i.UsexIsCurrent,
		UsexStartDate:          i.UsexStartDate,
		UsexEndDate:            i.UsexEndDate,
		UsexIndustry:           i.UsexIndustry,
		UsexDescription:        i.UsexDescription,
		UsexExperienceType:      i.UsexExperienceType,
		UsexCityID:             i.UsexCityID,
	}, nil
}

const updateExperience = `-- name: UpdateExperience :exec
UPDATE users.users_experiences
  set 
  usex_entity_id = $2,
  usex_title      = $3,
  usex_profile_headline    = $4,
  usex_employment_type     = $5,
  usex_company_name        = $6,
  usex_is_current          = $7,
  usex_start_date         = $8,
  usex_end_date           = $9,
  usex_industry            = $10,
  usex_description       = $11,
  usex_experience_type     = $12,
  usex_city_id             = $13
WHERE usex_id = $1
`
func (q *Queries) UpdateExperience(ctx context.Context, arg CreateExperienceParams) error {
	_, err := q.db.ExecContext(ctx, updateExperience,
		arg.UsexID,
		arg.UsexEntityID,
		arg.UsexTitle,
		arg.UsexProfileHeadline,
		arg.UsexEmploymentType,
		arg.UsexCompanyName,
		arg.UsexIsCurrent,
		arg.UsexStartDate,
		arg.UsexEndDate,
		arg.UsexIndustry,
		arg.UsexDescription,
		arg.UsexExperienceType,
		arg.UsexCityID,
	)
	return err
}

const deleteExperience = `-- name: DeleteExperience :exec
DELETE FROM users.users_experiences
WHERE usex_id = $1
`

func (q *Queries) DeleteExperience(ctx context.Context, usexID int32) error {
	_, err := q.db.ExecContext(ctx, deleteExperience, usexID)
	return err
}