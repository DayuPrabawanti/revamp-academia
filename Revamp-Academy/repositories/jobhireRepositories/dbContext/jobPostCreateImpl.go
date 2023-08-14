package dbContext

import (
	"context"
	"net/http"

	"codeid.revampacademy/models"
)

type CreateJobPostParams struct {
	JobHirePost     models.JobhireJobPost
	JobHirePostDesc models.JobhireJobPostDesc
}

// create to job_post, job_post_desc(description, benefit), master.address, master.city, master.province, master.country, client

const createJobPost = `-- name: CreateJobPost :one

WITH inserted_jobpost AS (
	INSERT INTO jobHire.job_post(jopo_number, jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, jopo_max_salary, jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date, jopo_modified_date, jopo_emp_entity_id, jopo_clit_id, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, jopo_work_code, jopo_edu_code, jopo_indu_code, jopo_status)
	VALUES ($1, $2, NOW(), current_timestamp + interval '30' day, $3, $4, $5, $6, $7, $8, 
	NOW(), current_timestamp + interval '5' day, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
	RETURNING jopo_entity_id
)
INSERT INTO jobHire.job_post_desc(jopo_entity_id, jopo_description, jopo_responsibility, jopo_target, jopo_benefit)
select jopo_entity_id, $19, $20, $21, $22
from inserted_jobpost
`

func (q *Queries) CreateJobPost(ctx context.Context, arg CreateJobPostParams) (*models.CreateJobPost, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createJobPost,
		// arg.JobHirePost,
		// arg.JobHirePostDesc,
		arg.JobHirePost.JopoNumber,
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
	err := row.Scan(
		// &i.JobHirePost,
		// &i.JobHirePostDesc,
		&i.JobHirePost.JopoNumber,
		&i.JobHirePost.JopoTitle,
		&i.JobHirePost.JopoMinSalary,
		&i.JobHirePost.JopoMaxSalary,
		&i.JobHirePost.JopoMinExperience,
		&i.JobHirePost.JopoMaxExperience,
		&i.JobHirePost.JopoPrimarySkill,
		&i.JobHirePost.JopoSecondarySkill,
		&i.JobHirePost.JopoEmpEntityID,
		&i.JobHirePost.JopoClitID,
		&i.JobHirePost.JopoJoroID,
		&i.JobHirePost.JopoJotyID,
		&i.JobHirePost.JopoJocaID,
		&i.JobHirePost.JopoAddrID,
		&i.JobHirePost.JopoWorkCode,
		&i.JobHirePost.JopoEduCode,
		&i.JobHirePost.JopoInduCode,
		&i.JobHirePost.JopoStatus,
		&i.JobHirePostDesc.JopoDescription,
		&i.JobHirePostDesc.JopoResponsibility,
		&i.JobHirePostDesc.JopoTarget,
		&i.JobHirePostDesc.JopoBenefit,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &models.CreateJobPost{
		JobHirePost:     i.JobHirePost,
		JobHirePostDesc: i.JobHirePostDesc,
	}, nil
}
