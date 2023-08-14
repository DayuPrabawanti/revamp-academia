package dbContext

import (
	"context"

	"codeid.revampacademy/models"
	feature "codeid.revampacademy/models/features"
)

const listJobPostSearch = `-- name: ListJobPost :many
SELECT jp.jopo_entity_id, jp.jopo_number, jp.jopo_title, jp.jopo_start_date, jp.jopo_end_date, jp.jopo_min_salary, jp.jopo_max_salary, jp.jopo_min_experience, 
jp.jopo_max_experience, jp.jopo_primary_skill, jp.jopo_secondary_skill, jp.jopo_publish_date, jp.jopo_modified_date, jp.jopo_emp_entity_id, jp.jopo_clit_id, 
jc.clit_name, jc.clit_about, jp.jopo_joro_id, jp.jopo_joty_id, jp.jopo_joca_id, jp.jopo_addr_id, ad.addr_city_id, ac.city_name, jpd.jopo_description, 
jp.jopo_work_code, jp.jopo_edu_code, jp.jopo_indu_code, jp.jopo_status, jr.joro_name, wt.woty_name
FROM jobHire.job_post jp
join master.address ad on jp.jopo_addr_id = ad.addr_id 
join master.city ac on ad.addr_city_id = ac.city_id
join jobHire.job_post_desc jpd on jp.jopo_entity_id = jpd.jopo_entity_id
join jobHire.client jc on jp.jopo_clit_id = jc.clit_id
join master.job_role jr on jp.jopo_joro_id = jr.joro_id
join master.working_type wt on jp.jopo_work_code = wt.woty_code 
where ac.city_name=$1 AND jr.joro_name=$2 AND wt.woty_name = $3
ORDER BY jp.jopo_title
LIMIT $4 OFFSET $5;
`

func (q *Queries) ListJobPostSearch(ctx context.Context, metadata *feature.Metadata) ([]models.MergeJobSearch, error) {
	rows, err := q.db.QueryContext(ctx, listJobPostSearch, metadata.Location, metadata.JobRole, metadata.WorkType, metadata.PageNo, metadata.PageSize)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.MergeJobSearch
	for rows.Next() {
		var i models.MergeJobSearch
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
			&i.MasterJobRole.JoroName,
			&i.MasterWorkingType.WotyName,
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
