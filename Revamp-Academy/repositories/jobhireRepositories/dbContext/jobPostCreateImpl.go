package dbContext

import (
	"context"
	"net/http"

	"codeid.revampacademy/models"
)

// create to job_post, job_post_desc(description, benefit), master.address, master.city, master.province, master.country, client

const createJobPost = `-- name: CreateJobPost :one
INSERT INTO jobHire.job_post(jopo_entity_id, jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, jopo_max_salary, 
	jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date, jopo_modified_date, jopo_emp_entity_id, 
	jopo_clit_id, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, jopo_work_code, jopo_edu_code, jopo_indu_code, jopo_status)
	values($1, &3, NOW(), current_timestamp + interval '30' day, $4, $5, $6, $7, $8, $9, 
	NOW(), current_timestamp + interval '5' day, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19);
	insert into jobHire.job_post_desc(jopo_entity_id, jopo_description, jopo_responsibility, jopo_target, jopo_benefit)
	select jopo_entity_id, $20, $21, $22, $23
	from jobHire.job_post
	where jopo_number = $2
`

type CreateJobPostParams struct {
	JobHirePost     models.JobhireJobPost
	JobHirePostDesc models.JobhireJobPostDesc
}

func (q *Queries) CreateJobPost(ctx context.Context, arg CreateJobPostParams) (*models.CreateJobPost, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createJobPost,
		// arg.JobHirePost,
		// arg.JobHirePostDesc,
		arg.JobHirePost.JopoEntityID,
		arg.JobHirePost.JopoNumber,
		arg.JobHirePost.JopoTitle,
		arg.JobHirePost.JopoStartDate,
		arg.JobHirePost.JopoEndDate,
		arg.JobHirePost.JopoMinSalary,
		arg.JobHirePost.JopoMaxSalary,
		arg.JobHirePost.JopoMinExperience,
		arg.JobHirePost.JopoMaxExperience,
		arg.JobHirePost.JopoPrimarySkill,
		arg.JobHirePost.JopoSecondarySkill,
		arg.JobHirePost.JopoPublishDate,
		arg.JobHirePost.JopoModifiedDate,
		arg.JobHirePost.JopoEmpEntityID,
		arg.JobHirePost.JopoClitID,
		arg.JobHirePost.JopoJoroID,
		arg.JobHirePost.JopoJotyID,
		arg.JobHirePost.JopoJocaID,
		arg.JobHirePost.JopoAddrID,
		arg.JobHirePost.JopoStatus,
		arg.JobHirePostDesc.JopoEntityID,
		arg.JobHirePostDesc.JopoDescription,
		arg.JobHirePostDesc.JopoResponsibility,
		arg.JobHirePostDesc.JopoTarget,
		arg.JobHirePostDesc.JopoBenefit,
	)
	i := models.CreateJobPost{}
	err := row.Scan(
		// &i.JobHirePost,
		// &i.JobHirePostDesc,
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
		&i.JobHirePost.JopoJoroID,
		&i.JobHirePost.JopoJotyID,
		&i.JobHirePost.JopoJocaID,
		&i.JobHirePost.JopoAddrID,
		&i.JobHirePost.JopoStatus,
		&i.JobHirePostDesc.JopoEntityID,
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
