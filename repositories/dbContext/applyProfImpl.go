package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const ListApplyProfImpl = `-- name: ListApplyProfImpl :many
select  
		uu.user_name,
		uu.user_birth_date,
		ue.usdu_degree,
		ue.usdu_school,
		ue.usdu_field_study,
		up.uspo_number,
		um.usme_filetype
from users.users uu
join users.users_education ue
on ue.usdu_entity_id = uu.user_entity_id
join users.users_phone up
on up.uspo_entity_id = uu.user_entity_id
join users.users_media um
on um.usme_entity_id = uu.user_entity_id
where uu.user_name = $1
`

func (q *Queries) ListApplyProfImpl(ctx context.Context, nama string) ([]models.ApplyProf, error) {
	rows, err := q.db.QueryContext(ctx, ListApplyProfImpl, nama)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.ApplyProf
	for rows.Next() {
		var i models.ApplyProf
		if err := rows.Scan(
			&i.UserUser.UserName,
			&i.UserUser.UserBirthDate,
			&i.UserEducation.UsduDegree,
			&i.UserEducation.UsduSchool,
			&i.UserEducation.UsduFieldStudy,
			&i.UserPhone.UspoNumber,
			&i.UserMedia.UsmeFiletype,
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
