package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

// Mockup jobDetail
const getJobPost = `-- name: GetJobPost :one
SELECT jp.jopo_entity_id, jp.jopo_number, jp.jopo_title, jp.jopo_start_date, jp.jopo_end_date, jp.jopo_min_salary, jp.jopo_max_salary, jp.jopo_min_experience, jp.jopo_max_experience, jp.jopo_primary_skill, jp.jopo_secondary_skill, jp.jopo_publish_date, jp.jopo_modified_date, jp.jopo_emp_entity_id, jp.jopo_clit_id, jc.clit_name, jc.clit_about, jp.jopo_joro_id, jp.jopo_joty_id, jp.jopo_joca_id, jp.jopo_addr_id, ad.addr_city_id, ac.city_name, jpd.jopo_description, jp.jopo_work_code, jp.jopo_edu_code, jp.jopo_indu_code, jp.jopo_status 
FROM jobHire.job_post jp
join master.address ad on jp.jopo_addr_id = ad.addr_id 
join master.city ac on ad.addr_city_id = ac.city_id
join jobHire.job_post_desc jpd on jp.jopo_entity_id = jpd.jopo_entity_id
join jobHire.client jc on jp.jopo_clit_id = jc.clit_id
WHERE jp.jopo_entity_id = $1
`

func (q *Queries) GetJobPostDetail(ctx context.Context, jopoEntityID int32) (models.MergeJobDetail, error) {
	row := q.db.QueryRowContext(ctx, getJobPost, jopoEntityID)
	var i models.MergeJobDetail
	err := row.Scan(
		&i.JobHirePost.JopoEntityID,
		&i.JobHirePost.JopoNumber,
		&i.JobHirePost.JopoTitle,
		&i.JobHirePost.JopoStartDate,
		&i.JobHirePost.JopoEndDate,
		&i.JobHirePost.JopoMinSalary,
		&i.JobHirePost.JopoMaxSalary,
		&i.JobHirePost.JopoMinExperience,
		&i.JobHirePost.JopoMaxExperience,
		&i.JobHirePost.JopoPrimarySkill,
		&i.JobHirePost.JopoSecondarySkill,
		&i.JobHirePost.JopoPublishDate,
		&i.JobHirePost.JopoModifiedDate,
		&i.JobHirePost.JopoEmpEntityID,
		&i.JobHirePost.JopoClitID,
		&i.JobHireClient.ClitName,
		&i.JobHireClient.ClitAbout,
		&i.JobHirePost.JopoJoroID,
		&i.JobHirePost.JopoJotyID,
		&i.JobHirePost.JopoJocaID,
		&i.JobHirePost.JopoAddrID,
		&i.MasterAddress.AddrCityID,
		&i.MasterCity.CityName,
		&i.JobHireJobPostDesc.JopoDescription,
		&i.JobHirePost.JopoWorkCode,
		&i.JobHirePost.JopoEduCode,
		&i.JobHirePost.JopoInduCode,
		&i.JobHirePost.JopoStatus,
	)
	return i, err
}

// take data from jobhire client
const listClient = `-- name: ListClient :many
SELECT clit_id, clit_name, clit_about, clit_modified_date, clit_addr_id, clit_emra_id FROM jobHire.client
ORDER BY clit_name
`

func (q *Queries) ListClient(ctx context.Context) ([]models.JobhireClient, error) {
	rows, err := q.db.QueryContext(ctx, listClient)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.JobhireClient
	for rows.Next() {
		var i models.JobhireClient
		if err := rows.Scan(
			&i.ClitID,
			&i.ClitName,
			&i.ClitAbout,
			&i.ClitModifiedDate,
			&i.ClitAddrID,
			&i.ClitEmraID,
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

// take data from jobpostdesc
const listJobPostDesc = `-- name: ListJobPostDesc :many
SELECT jopo_entity_id, jopo_description, jopo_responsibility, jopo_target, jopo_benefit FROM jobHire.job_post_desc
ORDER BY jopo_entity_id
`

func (q *Queries) ListJobPostDesc(ctx context.Context) ([]models.JobhireJobPostDesc, error) {
	rows, err := q.db.QueryContext(ctx, listJobPostDesc)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.JobhireJobPostDesc
	for rows.Next() {
		var i models.JobhireJobPostDesc
		if err := rows.Scan(
			&i.JopoEntityID,
			&i.JopoDescription,
			&i.JopoResponsibility,
			&i.JopoTarget,
			&i.JopoBenefit,
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

// take data from jobhireEmployeeRange
const listEmployeesRange = `-- name: ListEmployeesRange :many
SELECT emra_id, emra_range_min, emra_range_max, emra_modified_date FROM jobHire.employee_range
ORDER BY emra_range_max
`

func (q *Queries) ListEmployeesRange(ctx context.Context) ([]models.JobhireEmployeeRange, error) {
	rows, err := q.db.QueryContext(ctx, listEmployeesRange)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.JobhireEmployeeRange
	for rows.Next() {
		var i models.JobhireEmployeeRange
		if err := rows.Scan(
			&i.EmraID,
			&i.EmraRangeMin,
			&i.EmraRangeMax,
			&i.EmraModifiedDate,
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

const listMasterJobRole = `-- name: ListMasterJobRole :many
SELECT joro_id, joro_name, joro_modified_date FROM master.job_role
ORDER BY joro_id
`

func (q *Queries) ListMasterJobRole(ctx context.Context) ([]models.MasterJobRole, error) {
	rows, err := q.db.QueryContext(ctx, listMasterJobRole)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.MasterJobRole
	for rows.Next() {
		var i models.MasterJobRole
		if err := rows.Scan(&i.JoroID, &i.JoroName, &i.JoroModifiedDate); err != nil {
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

const listMasterWorkingType = `-- name: ListMasterWorkingType :many
SELECT woty_code, woty_name FROM master.working_type
ORDER BY woty_code
`

func (q *Queries) ListMasterWorkingType(ctx context.Context) ([]models.MasterWorkingType, error) {
	rows, err := q.db.QueryContext(ctx, listMasterWorkingType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.MasterWorkingType
	for rows.Next() {
		var i models.MasterWorkingType
		if err := rows.Scan(&i.WotyCode, &i.WotyName); err != nil {
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
