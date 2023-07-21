package dbContext

import (
	"context"

	"codeid.revampacademy/models"

)

const ListJobPostingImpl = `-- name: ListJobPostingImpl :many
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
		jt.jopo_status
from jobhire.job_post jt
join master.industry ne
on jt.jopo_indu_code = ne.indu_code_id
`

func (q *Queries) ListJobPostingImpl(ctx context.Context) ([]models.JobPosting, error) {
	rows, err := q.db.QueryContext(ctx, ListJobPostingImpl)
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

	