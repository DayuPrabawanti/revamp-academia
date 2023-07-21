package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listTalentsDetail = `-- name: ListTalentsDetail :many
SELECT
ms.*,
hr.*,
us.*,
jh.*,
bc.*
FROM hr.employee hr
JOIN master.job_role ms
ON hr.emp_joro_id = ms.joro_id
JOIN users.users us
ON hr.emp_entity_id = us.user_entity_id
JOIN jobhire.talent_apply jh
ON us.user_entity_id = jh.taap_entity_id
JOIN bootcamp.batch bc
ON hr.emp_entity_id = bc.batch_entity_id
ORDER BY hr.emp_entity_id
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
		if err := rows.Scan(&i.MasterJobRole.JoroID, &i.MasterJobRole.JoroName, &i.MasterJobRole.JoroModifiedDate,

			&i.HrEmployee.EmpEntityID, &i.HrEmployee.EmpEmpNumber, &i.HrEmployee.EmpNationalID, &i.HrEmployee.EmpBirthDate, &i.HrEmployee.EmpMaritalStatus, &i.HrEmployee.EmpGender, &i.HrEmployee.EmpHireDate, &i.HrEmployee.EmpSalariedFlag, &i.HrEmployee.EmpVacationHours, &i.HrEmployee.EmpSickleaveHours, &i.HrEmployee.EmpCurrentFlag, &i.HrEmployee.EmpModifiedDate, &i.HrEmployee.EmpType, &i.HrEmployee.EmpJoroID, &i.HrEmployee.EmpEmpEntityID,

			&i.UsersUser.UserEntityID, &i.UsersUser.UserName, &i.UsersUser.UserPassword, &i.UsersUser.UserFirstName, &i.UsersUser.UserLastName, &i.UsersUser.UserBirthDate, &i.UsersUser.UserEmailPromotion, &i.UsersUser.UserDemographic, &i.UsersUser.UserModifiedDate, &i.UsersUser.UserPhoto, &i.UsersUser.UserCurrentRole,

			&i.BootcampBatch.BatchID, &i.BootcampBatch.BatchEntityID, &i.BootcampBatch.BatchName, &i.BootcampBatch.BatchDescription, &i.BootcampBatch.BatchStartDate, &i.BootcampBatch.BatchEndDate, &i.BootcampBatch.BatchReason, &i.BootcampBatch.BatchType, &i.BootcampBatch.BatchModifiedDate, &i.BootcampBatch.BatchStatus, &i.BootcampBatch.BatchPicID,

			&i.JobhireTalentApply.TaapUserEntityID, &i.JobhireTalentApply.TaapEntityID, &i.JobhireTalentApply.TaapIntro, &i.JobhireTalentApply.TaapScoring, &i.JobhireTalentApply.TaapModifiedDate, &i.JobhireTalentApply.TaapStatus); err != nil {
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
