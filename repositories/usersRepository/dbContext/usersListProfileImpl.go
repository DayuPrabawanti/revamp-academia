package dbContext

import (
	"context"
	"database/sql"
)

// Menambahkan model untuk menampilkan profile
type UserProfile struct {
	User  UsernameParams
	Email ProfileEmailParams
}

// Membuat tampilan email untuk mockup profile user

type UsernameParams struct {
	UserFirstName      sql.NullString `db:"user_first_name" json:"userFirstName"`
	UserLastName       sql.NullString `db:"user_last_name" json:"userLastName"`
	UserFullname		sql.NullString	`json:"fullname"`
	UserPhoto          sql.NullString        `db:"user_photo" json:"userPhoto"`
}

type ProfileEmailParams struct {
	PmailAddress      sql.NullString `db:"pmail_address" json:"pmailAddress"`
}

type TampilProfile struct {
	UserFirstName      sql.NullString `db:"user_first_name" json:"userFirstName"`
	UserLastName       sql.NullString `db:"user_last_name" json:"userLastName"`
	UserFullname		sql.NullString	`json:"fullname"`
	UserPhoto          sql.NullString        `db:"user_photo" json:"userPhoto"`
	PmailAddress      sql.NullString `db:"pmail_address" json:"pmailAddress"`
	UspoNumber       string         `db:"uspo_number" json:"uspoNumber"`
	UspoPontyCode   string `db:"uspo_ponty_code" json:"uspoPontyCode"`
	AddrLine1           sql.NullString `db:"addr_line1" json:"addrLine1"`
	AddrLine2           sql.NullString `db:"addr_line2" json:"addrLine2"`
	AddrPostalCode      string `db:"addr_postal_code" json:"addrPostalCode"`
	CityName         sql.NullString `db:"city_name" json:"cityName"`
	UsduSchool       sql.NullString `db:"usdu_school" json:"usduSchool"`
	UsduDegree       sql.NullString `db:"usdu_degree" json:"usduDegree"`
	UsduFieldStudy   sql.NullString `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduGrade        sql.NullString `db:"usdu_grade" json:"usduGrade"`
	UsduStartDate    sql.NullTime   `db:"usdu_start_date" json:"usduStartDate"`
	UsduEndDate      sql.NullTime   `db:"usdu_end_date" json:"usduEndDate"`
	UsduActivities   sql.NullString `db:"usdu_activities" json:"usduActivities"`
	UsduDescription  sql.NullString `db:"usdu_description" json:"usduDescription"`
	UsexTitle           sql.NullString `db:"usex_title" json:"usexTitle"`
	UsexProfileHeadline sql.NullString `db:"usex_profile_headline" json:"usexProfileHeadline"`
	UsexCompanyName     sql.NullString `db:"usex_company_name" json:"usexCompanyName"`
	UsexStartDate       sql.NullTime   `db:"usex_start_date" json:"usexStartDate"`
	UsexEndDate         sql.NullTime   `db:"usex_end_date" json:"usexEndDate"`
	UsexCityID          sql.NullInt32  `db:"usex_city_id" json:"usexCityId"`
	UsexDescription     sql.NullString `db:"usex_description" json:"usexDescription"`
	UskiSktyName     sql.NullString `db:"uski_skty_name" json:"uskiSktyName"`
}




const getProfileUsers = `-- name: GetProfileUsers :one

SELECT user_first_name, user_last_name, CONCAT (user_first_name, user_last_name) AS fullname ,user_photo FROM users.users
WHERE user_entity_id = $1
`
// users
func (q *Queries) GetProfileUser(ctx context.Context, userEntityID int32) (UsernameParams, error) {
	row := q.db.QueryRowContext(ctx, getProfileUsers, userEntityID)
	var i UsernameParams
	err := row.Scan(
		&i.UserFirstName,
		&i.UserLastName,
		&i.UserFullname,
		&i.UserPhoto,
	)
	return i, err
}



const GetProfileEmail = `-- name: GetProfileEmail :many
SELECT pmail_address FROM users.users_email
WHERE pmail_entity_id = $1
`

func (q *Queries) GetProfileEmail(ctx context.Context, pmailEnttityId int32) ([]ProfileEmailParams, error) {
	rows, err := q.db.QueryContext(ctx, GetProfileEmail, pmailEnttityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProfileEmailParams
	for rows.Next() {
		var i ProfileEmailParams
		if err := rows.Scan(
			&i.PmailAddress,
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


const GetProfile = `-- name: GetProfile :many
SELECT us.user_first_name, user_last_name, CONCAT (user_first_name, user_last_name) AS fullname ,user_photo, 
ue.pmail_address, 
up.uspo_number, up.uspo_ponty_code,
ma.addr_line1, ma.addr_line2, ma.addr_postal_code,
mc.city_name,
uedu.usdu_school, uedu.usdu_degree, uedu.usdu_field_study, uedu.usdu_grade, uedu.usdu_start_date,
uedu.usdu_end_date, uedu.usdu_activities, uedu.usdu_description,
uex.usex_title, uex.usex_profile_headline, uex.usex_company_name, uex.usex_start_date, uex.usex_end_date, 
uex.usex_city_id, uex.usex_description,
uski.uski_skty_name
FROM users.users us
JOIN users.users_email ue 
ON us.user_entity_id = ue.pmail_entity_id
JOIN users.users_phones up
ON us.user_entity_id = up.uspo_entity_id
JOIN users.users_address uad
ON us.user_entity_id = uad.etad_entity_id
JOIN master.address ma
ON uad.etad_addr_id = ma.addr_id
JOIN master.city mc
ON ma.addr_city_id = mc.city_id
JOIN users.users_education uedu
ON us.user_entity_id = uedu.usdu_entity_id
JOIN users.users_experiences uex
ON us.user_entity_id = uex.usex_entity_id
JOIN users.users_skill uski
ON us.user_entity_id = uski.uski_entity_id
WHERE user_entity_id = $1
`
// GROUP BY us.user_first_name, user_last_name, user_photo, e.pmail_address


func (q *Queries) GetProfile(ctx context.Context, EnttityId int32) ([]TampilProfile, error) {
	rows, err := q.db.QueryContext(ctx, GetProfile, EnttityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []TampilProfile

	for rows.Next() {
		var i TampilProfile

		err := rows.Scan(
			&i.UserFirstName,
			&i.UserLastName,
			&i.UserFullname,
			&i.UserPhoto,
			&i.PmailAddress,
			&i.UspoNumber,
			&i.UspoPontyCode,
			&i.AddrLine1,
			&i.AddrLine2,
			&i.AddrPostalCode,
			&i.CityName,
			&i.UsduSchool,
			&i.UsduDegree,
			&i.UsduFieldStudy,
			&i.UsduGrade,
			&i.UsduStartDate,
			&i.UsduEndDate,
			&i.UsduActivities,
			&i.UsduDescription,
			&i.UsexTitle,
			&i.UsexProfileHeadline,
			&i.UsexCompanyName,
			&i.UsexStartDate,
			&i.UsexEndDate,
			&i.UsexCityID,
			&i.UsexDescription,
			&i.UskiSktyName,
		)

		if err != nil {
			return nil, err
		}
		profiles = append(profiles, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return profiles, nil
}