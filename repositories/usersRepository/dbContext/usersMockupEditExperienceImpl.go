package dbContext

import (
	"context"
	"time"
)

type AddExperienceParams struct {
	UsexTitle           string    `db:"usex_title" json:"usexTitle"`
	UsexProfileHeadline string    `db:"usex_profile_headline" json:"usexProfileHeadline"`
	UsexCompanyName     string    `db:"usex_company_name" json:"usexCompanyName"`
	UsexCityID          int32     `db:"usex_city_id" json:"usexCityId"`
	UsexStartDate       time.Time `db:"usex_start_date" json:"usexStartDate"`
	UsexEndDate         time.Time `db:"usex_end_date" json:"usexEndDate"`
	UsexIndustry        string `db:"usex_industry" json:"usexIndustry"`
	UsexEmploymentType  string    `db:"usex_employment_type" json:"usexEmploymentType"`
	UsexDescription     string    `db:"usex_description" json:"usexDescription"`
	UsexExperienceType  string    `db:"usex_experience_type" json:"usexExperienceType"`
}

const addExperience = `-- name: addExperience :one

INSERT INTO users.users_experiences(
	usex_entity_id, usex_title, usex_profile_headline, usex_company_name,
	usex_city_id, usex_start_date, usex_end_date, usex_industry, usex_employment_type,
	usex_description, usex_experience_type
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`

func (q *Queries) AddExperience(ctx context.Context, arg AddExperienceParams, id int32) error {
	_, err := q.db.ExecContext(ctx, addExperience, id,
	arg.UsexTitle,
	arg.UsexProfileHeadline,
	arg.UsexCompanyName,
	arg.UsexCityID,
	arg.UsexStartDate,
	arg.UsexEndDate,
	arg.UsexIndustry,
	arg.UsexEmploymentType,
	arg.UsexDescription,
	arg.UsexExperienceType,
	)
	return err
}
