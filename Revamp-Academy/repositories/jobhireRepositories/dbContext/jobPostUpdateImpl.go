package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const updateJobPost = `-- name: UpdateJobPostGabung :exec
WITH update_jobpost AS (
	update jobHire.job_post set jopo_title = $2, jopo_max_salary = $3, jopo_max_experience = $4, jopo_primary_skill = $5, jopo_secondary_skill = $6, jopo_modified_date = NOW(), jopo_clit_id = $7, jopo_joca_id = $8, jopo_addr_id = $9, jopo_work_code = $10, jopo_edu_code = $11, jopo_indu_code = $12, jopo_status = $13
	where jopo_entity_id = $1
	RETURNING jopo_entity_id
  )
  update jobHire.job_post_desc
  set jopo_description = $14
  where jopo_entity_id IN(
	  select jopo_entity_id from update_jobpost 
  )
`

type UpdateJobPostParams struct {
	JobHirePost     models.JobhireJobPost
	JobHirePostDesc models.JobhireJobPostDesc
}

func (q *Queries) UpdateJobPost(ctx context.Context, arg UpdateJobPostParams) error {
	_, err := q.db.ExecContext(ctx, updateJobPost,
		arg.JobHirePost.JopoEntityID,
		arg.JobHirePost.JopoTitle,
		arg.JobHirePost.JopoMaxSalary,
		arg.JobHirePost.JopoMaxExperience,
		arg.JobHirePost.JopoPrimarySkill,
		arg.JobHirePost.JopoSecondarySkill,
		arg.JobHirePost.JopoClitID,
		arg.JobHirePost.JopoJocaID,
		arg.JobHirePost.JopoAddrID,
		arg.JobHirePost.JopoWorkCode,
		arg.JobHirePost.JopoEduCode,
		arg.JobHirePost.JopoInduCode,
		arg.JobHirePost.JopoStatus,
		arg.JobHirePostDesc.JopoDescription,
	)
	return err
}
