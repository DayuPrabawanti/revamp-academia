package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listTalentsDetail = `-- name: ListTalentsDetail :many

SELECT us.user_first_name, us.user_last_name, us.user_photo,
bc.batch_name, bc.batch_type, bc.batch_status, bt.batr_review,
hr.emp_marital_status, hr.emp_gender,
ue.pmail_address,
up.uspo_number,
usk.uski_skty_name,
ms.skty_name,
ct.cate_name,
pe.prog_title

FROM curriculum.program_entity pe
JOIN master.category ct
ON pe.prog_cate_id = ct.cate_id
JOIN bootcamp.batch bc 
ON bc.batch_id = pe.prog_entity_id
JOIN bootcamp.batch_trainee bt
ON bt.batr_batch_id = bc.batch_id
JOIN hr.employee hr
ON bc.batch_entity_id = hr.emp_entity_id
JOIN users.users us
ON hr.emp_entity_id = us.user_entity_id
JOIN users.users_email ue
ON ue.pmail_entity_id = us.user_entity_id
JOIN users.users_phones up
ON up.uspo_entity_id = us.user_entity_id
JOIN users.users_skill usk
ON us.user_entity_id = usk.uski_entity_id
JOIN master.skill_type ms
ON usk.uski_skty_name = ms.skty_name
ORDER BY us.user_entity_id
`

func (q *Queries) ListTalentsDetail(ctx context.Context) ([]models.TalentsDetailMockup, error) {
	rows, err := q.db.QueryContext(ctx, listTalentsDetail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.TalentsDetailMockup
	for rows.Next() {
		var i models.TalentsDetailMockup
		if err := rows.Scan(&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStatus, &i.BootcampBatchTrainee.BatrReview,
			&i.HrEmployee.EmpMaritalStatus, &i.HrEmployee.EmpGender,
			&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,
			&i.UsersUsersEmail.PmailAddress,
			&i.UsersUsersPhone.UspoNumber,
			&i.UsersUsersSkill.UskiSktyName,
			&i.MasterSkillType.SktyName,
			&i.MasterCategory.CateName,
			&i.CurriculumProgramEntity.ProgTitle); err != nil {
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

const getTalentDetail = `-- name: GetTalentDetail :one

SELECT us.user_first_name, us.user_last_name, us.user_photo,
bc.batch_name, bc.batch_type, bc.batch_status, bt.batr_review,
hr.emp_marital_status, hr.emp_gender,
ue.pmail_address,
up.uspo_number,
usk.uski_skty_name,
ms.skty_name,
ct.cate_name,
pe.prog_title

FROM curriculum.program_entity pe
JOIN master.category ct
ON pe.prog_cate_id = ct.cate_id
JOIN bootcamp.batch bc 
ON bc.batch_id = pe.prog_entity_id
JOIN bootcamp.batch_trainee bt
ON bt.batr_batch_id = bc.batch_id
JOIN hr.employee hr
ON bc.batch_entity_id = hr.emp_entity_id
JOIN users.users us
ON hr.emp_entity_id = us.user_entity_id
JOIN users.users_email ue
ON ue.pmail_entity_id = us.user_entity_id
JOIN users.users_phones up
ON up.uspo_entity_id = us.user_entity_id
JOIN users.users_skill usk
ON us.user_entity_id = usk.uski_entity_id
JOIN master.skill_type ms
ON usk.uski_skty_name = ms.skty_name
WHERE bc.batch_id = $1
`

// hr.department
func (q *Queries) GetTalentDetail(ctx context.Context, batchId int32) (models.TalentsDetailMockup, error) {
	row := q.db.QueryRowContext(ctx, getTalentDetail, batchId)
	var i models.TalentsDetailMockup
	err := row.Scan(&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStatus, &i.BootcampBatchTrainee.BatrReview,
		&i.HrEmployee.EmpMaritalStatus, &i.HrEmployee.EmpGender,
		&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,
		&i.UsersUsersEmail.PmailAddress,
		&i.UsersUsersPhone.UspoNumber,
		&i.UsersUsersSkill.UskiSktyName,
		&i.MasterSkillType.SktyName,
		&i.MasterCategory.CateName,
		&i.CurriculumProgramEntity.ProgTitle)
	return i, err
}
