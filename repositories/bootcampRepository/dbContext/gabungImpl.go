package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const Gabung = `-- name: Gabung :many

select bc.prap_user_entity_id,
		bc.prap_prog_entity_id,
		bc.prap_status,
		bc.prap_review,
		bc.prap_test_score,
		bc.prap_modified_date,
		
		ud.usdu_school,
		ud.usdu_field_study,
		ud.usdu_graduate_year,
		
		us.user_name,
		us.user_photo,
		ue.pmail_address,
		
		up.uspo_number 
from users.users us 
join users.users_email ue on us.user_entity_id = ue.pmail_entity_id 
join users.users_education ud on us.user_entity_id = ud.usdu_entity_id
join users.users_phones up on us.user_entity_id = up.uspo_entity_id
join curriculum.program_entity pe on us.user_entity_id = pe.prog_entity_id
join bootcamp.program_apply bc on pe.prog_entity_id = bc.prap_prog_entity_id
`

func (q *Queries) Gabung(ctx context.Context) ([]models.Gabung, error) {
	rows, err := q.db.QueryContext(ctx, Gabung)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Gabung
	for rows.Next() {
		var i models.Gabung
		if err := rows.Scan(
			&i.BootcampProgramApply.PrapUserEntityID,
			&i.CurriculumProgramEntity.ProgEntityID,
			&i.BootcampProgramApply.PrapStatus,
			&i.BootcampProgramApply.PrapReview,
			&i.BootcampProgramApply.PrapTestScore,
			&i.BootcampProgramApply.PrapModifiedDate,
			&i.UsersUser.UserName,
			&i.UsersUser.UserPhoto,
			&i.UsersUsersEmail.PmailAddress,
			&i.UsersUsersEducation.UsduFieldStudy,
			&i.UsersUsersEducation.UsduSchool,
			&i.UsersUsersEducation.UsduGraduateYear,
			&i.UsersPhoneNumberType.PontyCode,
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
