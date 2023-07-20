package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listTalents = `-- name: ListTalents :many
SELECT
bc.*,
hr.*,
us.*,
usk.*,
ms.*
FROM bootcamp.batch bc 
JOIN hr.employee hr
ON bc.batch_entity_id = hr.emp_entity_id
JOIN users.users us
ON hr.emp_entity_id = us.user_entity_id
JOIN users.users_skill usk
ON us.user_entity_id = usk.uski_entity_id
JOIN master.skill_type ms
ON usk.uski_skty_name = ms.skty_name
ORDER BY usk.uski_skty_name
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
		if err := rows.Scan(&i.BootcampBatch.BatchID, &i.BootcampBatch.BatchEntityID, &i.BootcampBatch.BatchName, &i.BootcampBatch.BatchDescription, &i.BootcampBatch.BatchStartDate, &i.BootcampBatch.BatchEndDate, &i.BootcampBatch.BatchReason, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchModifiedDate, &i.BootcampBatch.BatchStatus, &i.BootcampBatch.BatchPicID,
			&i.HrEmployee.EmpEntityID, &i.HrEmployee.EmpEmpNumber, &i.HrEmployee.EmpNationalID, &i.HrEmployee.EmpBirthDate, &i.HrEmployee.EmpMaritalStatus, &i.HrEmployee.EmpGender, &i.HrEmployee.EmpHireDate, &i.HrEmployee.EmpSalariedFlag, &i.HrEmployee.EmpVacationHours, &i.HrEmployee.EmpSickleaveHours, &i.HrEmployee.EmpCurrentFlag, &i.HrEmployee.EmpModifiedDate, &i.HrEmployee.EmpType, &i.HrEmployee.EmpJoroID, &i.HrEmployee.EmpEmpEntityID,
			&i.UsersUser.UserEntityID, &i.UsersUser.UserName, &i.UsersUser.UserPassword, &i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserBirthDate, &i.UsersUser.UserEmailPromotion, &i.UsersUser.UserDemographic, &i.UsersUser.UserModifiedDate, &i.UsersUser.UserPhoto, &i.UsersUser.UserCurrentRole,
			&i.UsersUsersSkill.UskiID, &i.UsersUsersSkill.UskiEntityID, &i.UsersUsersSkill.UskiModifiedDate, &i.UsersUsersSkill.UskiSktyName,
			&i.MasterSkillType.SktyName); err != nil {
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
