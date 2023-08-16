package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

type CreateJobPostParams struct {
	JobHirePost     models.JobhireJobPost
	JobHirePostDesc models.JobhireJobPostDesc
}

// create to job_post, job_post_desc(description, benefit), master.address, master.city, master.province, master.country, client

const createJobPost = `-- name: CreateJobPost :one

WITH inserted_jobpost AS (
	INSERT INTO jobHire.job_post(jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, jopo_max_salary, jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date, jopo_modified_date, jopo_emp_entity_id, jopo_clit_id, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, jopo_work_code, jopo_edu_code, jopo_indu_code, jopo_status)
	VALUES ($1, NOW(), current_timestamp + interval '30' day, $2, $3, $4, $5, $6, $7, 
	NOW(), current_timestamp + interval '5' day, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	RETURNING jopo_entity_id
)
INSERT INTO jobHire.job_post_desc(jopo_entity_id, jopo_description, jopo_responsibility, jopo_target, jopo_benefit)
select jopo_entity_id, $18, $19, $20, $21
from inserted_jobpost
returning *
`

func (q *Queries) CreateJobPost(ctx context.Context, arg CreateJobPostParams) (*models.CreateJobPost, *models.ResponseError) {
	q.db.QueryRowContext(ctx, createJobPost,
		// arg.JobHirePost,
		// arg.JobHirePostDesc,
		arg.JobHirePost.JopoTitle,
		arg.JobHirePost.JopoMinSalary,
		arg.JobHirePost.JopoMaxSalary,
		arg.JobHirePost.JopoMinExperience,
		arg.JobHirePost.JopoMaxExperience,
		arg.JobHirePost.JopoPrimarySkill,
		arg.JobHirePost.JopoSecondarySkill,
		arg.JobHirePost.JopoEmpEntityID,
		arg.JobHirePost.JopoClitID,
		arg.JobHirePost.JopoJoroID,
		arg.JobHirePost.JopoJotyID,
		arg.JobHirePost.JopoJocaID,
		arg.JobHirePost.JopoAddrID,
		arg.JobHirePost.JopoWorkCode,
		arg.JobHirePost.JopoEduCode,
		arg.JobHirePost.JopoInduCode,
		arg.JobHirePost.JopoStatus,
		arg.JobHirePostDesc.JopoDescription,
		arg.JobHirePostDesc.JopoResponsibility,
		arg.JobHirePostDesc.JopoTarget,
		arg.JobHirePostDesc.JopoBenefit,
	)
	i := models.CreateJobPost{}

	return &models.CreateJobPost{
		JobHirePost:     i.JobHirePost,
		JobHirePostDesc: i.JobHirePostDesc,
	}, nil
}
