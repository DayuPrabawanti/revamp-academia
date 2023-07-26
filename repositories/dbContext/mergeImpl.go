package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const getJobPostingImpl = `-- name: GetJobPostingImpl 
select  
		jt.jopo_title,
		ne.indu_name,
		jt.jopo_entity_id,
		ne.indu_code_id,
		jt.jopo_start_date,
		jt.jopo_end_date,
		jt.jopo_min_salary,
		jt.jopo_max_salary,
		jt.jopo_min_experience,
		jt.jopo_primary_skill,
		jt.jopo_status,
		jt.jopo_work_code,
		wt.woty_name
from jobhire.job_post jt
join master.industry ne
on jt.jopo_indu_code = ne.indu_code_id
join master.working_type wt
on jt.jopo_work_code = wt.woty_code
WHERE jt.jopo_title like $1
`


func (q *Queries) GetJobPostingImpl(ctx context.Context, title string) (models.JobPosting, error) {
	row := q.db.QueryRowContext(ctx, getJobPostingImpl, title)
	var i models.JobPosting
	err := row.Scan(
			&i.JobhireJobPost.JopoTitle,
			&i.MasterIndustry.InduName,
			&i.JobhireJobPost.JopoEntityID,
			&i.MasterIndustry.InduCodeID,
			&i.JobhireJobPost.JopoStartDate,
			&i.JobhireJobPost.JopoEndDate,
			&i.JobhireJobPost.JopoMinSalary,
			&i.JobhireJobPost.JopoMaxSalary,
			&i.JobhireJobPost.JopoMinExperience,
			&i.JobhireJobPost.JopoPrimarySkill,
			&i.JobhireJobPost.JopoStatus,
			&i.JobhireJobPost.JopoWorkCode,
			&i.MasterWorkingType.WotyName,
	)
	return i, err
}


const listJobPostingImpl = `-- name: ListJobPostingImpl :many
select  
		jt.jopo_title,
		ne.indu_name,
		jt.jopo_entity_id,
		ne.indu_code_id,
		jt.jopo_start_date,
		jt.jopo_end_date,
		jt.jopo_min_salary,
		jt.jopo_max_salary,
		jt.jopo_min_experience,
		jt.jopo_primary_skill,
		jt.jopo_status,
		jt.jopo_work_code,
		wt.woty_name
from jobhire.job_post jt
join master.industry ne
on jt.jopo_indu_code = ne.indu_code_id
join master.working_type wt
on jt.jopo_work_code = wt.woty_code
where jt.jopo_title = $1
`

func (q *Queries) ListJobPostingImpl(ctx context.Context, nama string) ([]models.JobPosting, error) {
	rows, err := q.db.QueryContext(ctx, listJobPostingImpl, nama)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.JobPosting
	for rows.Next() {
		var i models.JobPosting
		if err := rows.Scan(
			&i.JobhireJobPost.JopoTitle,
			&i.MasterIndustry.InduName,
			&i.JobhireJobPost.JopoEntityID,
			&i.MasterIndustry.InduCodeID,
			&i.JobhireJobPost.JopoStartDate,
			&i.JobhireJobPost.JopoEndDate,
			&i.JobhireJobPost.JopoMinSalary,
			&i.JobhireJobPost.JopoMaxSalary,
			&i.JobhireJobPost.JopoMinExperience,
			&i.JobhireJobPost.JopoPrimarySkill,
			&i.JobhireJobPost.JopoStatus,
			&i.JobhireJobPost.JopoWorkCode,
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


