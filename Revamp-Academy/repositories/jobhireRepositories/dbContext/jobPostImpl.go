package dbContext

import (
	"context"

	"codeid.revampacademy/models"
	feature "codeid.revampacademy/models/features"
)

const getlistJobPost = `-- name: ListJobPost :many
SELECT jopo_entity_id, jopo_number, jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, jopo_max_salary, jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date, jopo_modified_date, jopo_emp_entity_id, jopo_clit_id, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, jopo_work_code, jopo_edu_code, jopo_indu_code, jopo_status FROM jobHire.job_post
ORDER BY jopo_title
limit $1 offset $2
`

func (q *Queries) GetListJobPost(ctx context.Context, metadata *feature.Metadata) ([]models.JobhireJobPost, error) {
	rows, err := q.db.QueryContext(ctx, listJobPost, metadata.PageNo, metadata.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.JobhireJobPost
	for rows.Next() {
		var i models.JobhireJobPost
		if err := rows.Scan(
			&i.JopoEntityID,
			&i.JopoNumber,
			&i.JopoTitle,
			&i.JopoStartDate,
			&i.JopoEndDate,
			&i.JopoMinSalary,
			&i.JopoMaxSalary,
			&i.JopoMinExperience,
			&i.JopoMaxExperience,
			&i.JopoPrimarySkill,
			&i.JopoSecondarySkill,
			&i.JopoPublishDate,
			&i.JopoModifiedDate,
			&i.JopoEmpEntityID,
			&i.JopoClitID,
			&i.JopoJoroID,
			&i.JopoJotyID,
			&i.JopoJocaID,
			&i.JopoAddrID,
			&i.JopoWorkCode,
			&i.JopoEduCode,
			&i.JopoInduCode,
			&i.JopoStatus,
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

// list and add location name by addr_id in master
const listJobPost = `-- name: ListJobPost :many
SELECT jopo_entity_id, jopo_number, jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, jopo_max_salary, 
jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date, jopo_modified_date, 
jopo_emp_entity_id, jopo_clit_id, clit_name, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, addr_city_id, city_name, jopo_work_code, jopo_edu_code, jopo_indu_code, jopo_status, emra_range_min, emra_range_max 
FROM jobHire.job_post 
join master.address on jobHire.job_post.jopo_addr_id = master.address.addr_id 
join master.city on master.address.addr_city_id = master.city.city_id
join jobHire.client on jobHire.job_post.jopo_clit_id = jobHire.client.clit_id
join jobHire.employee_range on jobHire.client.clit_emra_id = jobHire.employee_range.emra_id
ORDER BY jopo_title
limit $1 offset $2
`

func (q *Queries) ListJobPost(ctx context.Context, metadata *feature.Metadata) ([]models.MergeJobAndMaster, error) {
	rows, err := q.db.QueryContext(ctx, listJobPost, metadata.PageNo, metadata.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.MergeJobAndMaster
	for rows.Next() {
		var i models.MergeJobAndMaster
		if err := rows.Scan(
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
			&i.JobHirePost.JopoJoroID,
			&i.JobHirePost.JopoJotyID,
			&i.JobHirePost.JopoJocaID,
			&i.JobHirePost.JopoAddrID,
			&i.MasterAddress.AddrCityID,
			&i.MasterCity.CityName,
			&i.JobHirePost.JopoWorkCode,
			&i.JobHirePost.JopoEduCode,
			&i.JobHirePost.JopoInduCode,
			&i.JobHirePost.JopoStatus,
			&i.JobHireEmployeeRange.EmraRangeMin,
			&i.JobHireEmployeeRange.EmraRangeMax,
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
