package dbContext

import "database/sql"

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