package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listTalents = `-- name: ListTalents :many
SELECT bc.batch_name, bc.batch_type, bc.batch_status,

us.user_first_name, us.user_last_name, us.user_photo,

usk.uski_skty_name,
ms.skty_name,
ct.cate_name,
pe.prog_title

FROM curriculum.program_entity pe
JOIN master.category ct
ON pe.prog_cate_id = ct.cate_id
JOIN bootcamp.batch bc 
ON bc.batch_id = pe.prog_entity_id
JOIN hr.employee hr
ON bc.batch_entity_id = hr.emp_entity_id
JOIN users.users us
ON hr.emp_entity_id = us.user_entity_id

JOIN users.users_skill usk
ON us.user_entity_id = usk.uski_entity_id
JOIN master.skill_type ms
ON usk.uski_skty_name = ms.skty_name
ORDER BY us.user_entity_id
`

func (q *Queries) ListTalents(ctx context.Context) ([]models.TalentsMockup, error) {
	rows, err := q.db.QueryContext(ctx, listTalents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.TalentsMockup
	for rows.Next() {
		var i models.TalentsMockup
		if err := rows.Scan(&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStatus,

			&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,

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

const searchTalent = `-- name: SearchTalent :many
SELECT bc.batch_name, bc.batch_type, bc.batch_status,

us.user_first_name, us.user_last_name, us.user_photo,

usk.uski_skty_name,
ms.skty_name,
ct.cate_name,
pe.prog_title

FROM curriculum.program_entity pe
JOIN master.category ct
ON pe.prog_cate_id = ct.cate_id
JOIN bootcamp.batch bc 
ON bc.batch_id = pe.prog_entity_id
JOIN hr.employee hr
ON bc.batch_entity_id = hr.emp_entity_id
JOIN users.users us
ON hr.emp_entity_id = us.user_entity_id

JOIN users.users_skill usk
ON us.user_entity_id = usk.uski_entity_id
JOIN master.skill_type ms
ON usk.uski_skty_name = ms.skty_name
WHERE us.user_name like '%' || $1 || '%' AND usk.usk_skty_name like '%' || $2 || '%' AND bc.batch_name like '%' || $3 || '%' AND batch_status = $4
`

func (q *Queries) SearchTalent(ctx context.Context, userName string, userSkill string, batchName string, status string) ([]models.TalentsMockup, error) {
	rows, err := q.db.QueryContext(ctx, searchTalent, userName, userSkill, batchName, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var talents []models.TalentsMockup
	for rows.Next() {
		var i models.TalentsMockup
		if err := rows.Scan(
			&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStatus,
			&i.HrEmployee.EmpMaritalStatus, &i.HrEmployee.EmpGender,
			&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,
			&i.UsersUsersSkill.UskiSktyName,
			&i.MasterSkillType.SktyName,
			&i.MasterCategory.CateName,
			&i.CurriculumProgramEntity.ProgTitle,
		); err != nil {
			return nil, err
		}
		talents = append(talents, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return talents, nil
}

const pagingTalent = `-- name: PagingTalent :many
SELECT bc.batch_name, bc.batch_type, bc.batch_status,
hr.emp_marital_status, hr.emp_gender,
us.user_first_name, us.user_last_name, us.user_photo,
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
LIMIT $1 OFFSET $2
`

func (q *Queries) PagingTalent(ctx context.Context, limit, offset int) ([]models.TalentsMockup, error) {
	rows, err := q.db.QueryContext(ctx, pagingTalent, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var talents []models.TalentsMockup
	for rows.Next() {
		var i models.TalentsMockup
		if err := rows.Scan(
			&i.BootcampBatch.BatchName, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchStatus,
			&i.HrEmployee.EmpMaritalStatus, &i.HrEmployee.EmpGender,
			&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserPhoto,
			&i.UsersUsersSkill.UskiSktyName,
			&i.MasterSkillType.SktyName,
			&i.MasterCategory.CateName,
			&i.CurriculumProgramEntity.ProgTitle,
		); err != nil {
			return nil, err
		}
		talents = append(talents, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return talents, nil
}
