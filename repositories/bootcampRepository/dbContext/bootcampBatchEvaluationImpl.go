package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listBootcampBatchEvaluations = `-- name: Group :many
SELECT 
    b.batch_id,  
    b.batch_name,  
    pe.prog_title,
    us.user_photo,
    us.user_name,
    b.batch_status,
    bte.btev_skor
FROM 
    bootcamp.batch b
JOIN 
    curriculum.program_entity pe ON b.batch_entity_id = pe.prog_entity_id
JOIN 
    users.users us ON us.user_entity_id = b.batch_entity_id
JOIN 
    bootcamp.batch_trainee_evaluation bte ON b.batch_id = bte.btev_batch_id
ORDER BY
	b.batch_id;
`

func (q *Queries) ListBootcampBatchEvaluation(ctx context.Context, batchId int32) ([]models.BootcampBatchEvaluationMockup, error) {
	rows, err := q.db.QueryContext(ctx, listBootcampBatchEvaluations, batchId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BootcampBatchEvaluationMockup
	for rows.Next() {
		var i models.BootcampBatchEvaluationMockup
		if err := rows.Scan(
			&i.BootcampBatch.BatchName,
			&i.CurriculumProgramEntity.ProgTitle,
			&i.UsersUser.UserPhoto,
			&i.UsersUser.UserName,
			&i.BootcampBatch.BatchStatus,
			&i.BootcampBatchTraineeEvaluation.BtevSkor,
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

// import (
// 	"context"

// 	"codeid.revampacademy/models"
// )

// const listGroupsQuery = `-- name: ListGroups :many
// SELECT
//   b.batch_id,
//   b.batch_entity_id,
//   b.batch_name,
//   b.batch_description,
//   b.batch_start_date,
//   b.batch_end_date,
//   b.batch_reason,
//   b.batch_type,
//   b.batch_modified_date,
//   b.batch_status,
//   b.batch_pic_id,
//   btev.btev_id,
//   btev.btev_type,
//   btev.btev_header,
//   btev.btev_section,
//   btev.btev_skill,
//   btev.btev_week,
//   btev.btev_skor,
//   btev.btev_note,
//   cpe.prog_entity_id,
//   cpe.prog_title,
//   cpe.prog_headline,
//   cpe.prog_type,
//   cpe.prog_learning_type,
//   cpe.prog_rating,
//   cpe.prog_total_trainee,
//   cpe.prog_modified_date,
//   cpe.prog_image,
//   cpe.prog_best_seller,
//   cpe.prog_price,
//   cpe.prog_language,
//   cpe.prog_duration,
//   cpe.prog_duration_type,
//   cpe.prog_tag_skill,
//   uu.user_entity_id,
//   uu.user_name,
//   uu.user_password,
//   uu.user_first_name,
//   uu.user_last_name,
//   uu.user_birth_date,
//   uu.user_email_promotion,
//   uu.user_demographic,
//   uu.user_modified_date,
//   uu.user_photo,
//   uu.user_current_role
// FROM
//   bootcamp.batch b
//   LEFT JOIN bootcamp.bootcamp_batch_trainee_evaluation btev ON b.batch_id = btev.btev_batch_id
//   LEFT JOIN bootcamp.curriculum_program_entity cpe ON b.batch_id = cpe.prog_entity_id
//   LEFT JOIN bootcamp.users_user uu ON b.batch_id = uu.user_entity_id
// ORDER BY
//   b.batch_id
// `

// func (r *Queries) ListGroups(ctx context.Context) ([]*models.Group, error) {
// 	rows, err := r.db.QueryContext(ctx, listGroupsQuery)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var groups []*models.Group

// 	for rows.Next() {
// 		var group models.Group
// 		// Scan the result into the group struct
// 		err := rows.Scan(
// 			&group.BootcampBatch.BatchID,
// 			&group.BootcampBatch.BatchEntityID,
// 			&group.BootcampBatch.BatchName,
// 			&group.BootcampBatch.BatchDescription,
// 			&group.BootcampBatch.BatchStartDate,
// 			&group.BootcampBatch.BatchEndDate,
// 			&group.BootcampBatch.BatchReason,
// 			&group.BootcampBatch.BatchType,
// 			&group.BootcampBatch.BatchModifiedDate,
// 			&group.BootcampBatch.BatchStatus,
// 			&group.BootcampBatch.BatchPicID,
// 			&group.BootcampBatchTraineeEvaluation.BtevID,
// 			&group.BootcampBatchTraineeEvaluation.BtevType,
// 			&group.BootcampBatchTraineeEvaluation.BtevHeader,
// 			&group.BootcampBatchTraineeEvaluation.BtevSection,
// 			&group.BootcampBatchTraineeEvaluation.BtevSkill,
// 			&group.BootcampBatchTraineeEvaluation.BtevWeek,
// 			&group.BootcampBatchTraineeEvaluation.BtevSkor,
// 			&group.BootcampBatchTraineeEvaluation.BtevNote,
// 			&group.CurriculumProgramEntity.ProgEntityID,
// 			&group.CurriculumProgramEntity.ProgTitle,
// 			&group.CurriculumProgramEntity.ProgHeadline,
// 			&group.CurriculumProgramEntity.ProgType,
// 			&group.CurriculumProgramEntity.ProgLearningType,
// 			&group.CurriculumProgramEntity.ProgRating,
// 			&group.CurriculumProgramEntity.ProgTotalTrainee,
// 			&group.CurriculumProgramEntity.ProgModifiedDate,
// 			&group.CurriculumProgramEntity.ProgImage,
// 			&group.CurriculumProgramEntity.ProgBestSeller,
// 			&group.CurriculumProgramEntity.ProgPrice,
// 			&group.CurriculumProgramEntity.ProgLanguage,
// 			&group.CurriculumProgramEntity.ProgDuration,
// 			&group.CurriculumProgramEntity.ProgDurationType,
// 			&group.CurriculumProgramEntity.ProgTagSkill,
// 			&group.UsersUser.UserEntityID,
// 			&group.UsersUser.UserName,
// 			&group.UsersUser.UserPassword,
// 			&group.UsersUser.UserFirstName,
// 			&group.UsersUser.UserLastName,
// 			&group.UsersUser.UserBirthDate,
// 			&group.UsersUser.UserEmailPromotion,
// 			&group.UsersUser.UserDemographic,
// 			&group.UsersUser.UserModifiedDate,
// 			&group.UsersUser.UserPhoto,
// 			&group.UsersUser.UserCurrentRole,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		groups = append(groups, &group)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return groups, nil
// }

// const getGroupByIDQuery = `-- name: GetGroupByID :one
// SELECT
//   b.batch_id,
//   b.batch_entity_id,
//   b.batch_name,
//   b.batch_description,
//   b.batch_start_date,
//   b.batch_end_date,
//   b.batch_reason,
//   b.batch_type,
//   b.batch_modified_date,
//   b.batch_status,
//   b.batch_pic_id,
//   btev.btev_id,
//   btev.btev_type,
//   btev.btev_header,
//   btev.btev_section,
//   btev.btev_skill,
//   btev.btev_week,
//   btev.btev_skor,
//   btev.btev_note,
//   cpe.prog_entity_id,
//   cpe.prog_title,
//   cpe.prog_headline,
//   cpe.prog_type,
//   cpe.prog_learning_type,
//   cpe.prog_rating,
//   cpe.prog_total_trainee,
//   cpe.prog_modified_date,
//   cpe.prog_image,
//   cpe.prog_best_seller,
//   cpe.prog_price,
//   cpe.prog_language,
//   cpe.prog_duration,
//   cpe.prog_duration_type,
//   cpe.prog_tag_skill,
//   uu.user_entity_id,
//   uu.user_name,
//   uu.user_password,
//   uu.user_first_name,
//   uu.user_last_name,
//   uu.user_birth_date,
//   uu.user_email_promotion,
//   uu.user_demographic,
//   uu.user_modified_date,
//   uu.user_photo,
//   uu.user_current_role
// FROM
//   bootcamp.batch b
//   LEFT JOIN bootcamp.bootcamp_batch_trainee_evaluation btev ON b.batch_id = btev.btev_batch_id
//   LEFT JOIN bootcamp.curriculum_program_entity cpe ON b.batch_id = cpe.prog_entity_id
//   LEFT JOIN bootcamp.users_user uu ON b.batch_id = uu.user_entity_id
// WHERE
//   b.batch_id = $1
// `

// func (r *Queries) GetGroupByID(ctx context.Context, batchID int32) (*models.Group, error) {
// 	row := r.db.QueryRowContext(ctx, getGroupByIDQuery, batchID)
// 	var group models.Group
// 	err := row.Scan(
// 		&group.BootcampBatch.BatchID,
// 		&group.BootcampBatch.BatchEntityID,
// 		&group.BootcampBatch.BatchName,
// 		&group.BootcampBatch.BatchDescription,
// 		&group.BootcampBatch.BatchStartDate,
// 		&group.BootcampBatch.BatchEndDate,
// 		&group.BootcampBatch.BatchReason,
// 		&group.BootcampBatch.BatchType,
// 		&group.BootcampBatch.BatchModifiedDate,
// 		&group.BootcampBatch.BatchStatus,
// 		&group.BootcampBatch.BatchPicID,
// 		&group.BootcampBatchTraineeEvaluation.BtevID,
// 		&group.BootcampBatchTraineeEvaluation.BtevType,
// 		&group.BootcampBatchTraineeEvaluation.BtevHeader,
// 		&group.BootcampBatchTraineeEvaluation.BtevSection,
// 		&group.BootcampBatchTraineeEvaluation.BtevSkill,
// 		&group.BootcampBatchTraineeEvaluation.BtevWeek,
// 		&group.BootcampBatchTraineeEvaluation.BtevSkor,
// 		&group.BootcampBatchTraineeEvaluation.BtevNote,
// 		&group.CurriculumProgramEntity.ProgEntityID,
// 		&group.CurriculumProgramEntity.ProgTitle,
// 		&group.CurriculumProgramEntity.ProgHeadline,
// 		&group.CurriculumProgramEntity.ProgType,
// 		&group.CurriculumProgramEntity.ProgLearningType,
// 		&group.CurriculumProgramEntity.ProgRating,
// 		&group.CurriculumProgramEntity.ProgTotalTrainee,
// 		&group.CurriculumProgramEntity.ProgModifiedDate,
// 		&group.CurriculumProgramEntity.ProgImage,
// 		&group.CurriculumProgramEntity.ProgBestSeller,
// 		&group.CurriculumProgramEntity.ProgPrice,
// 		&group.CurriculumProgramEntity.ProgLanguage,
// 		&group.CurriculumProgramEntity.ProgDuration,
// 		&group.CurriculumProgramEntity.ProgDurationType,
// 		&group.CurriculumProgramEntity.ProgTagSkill,
// 		&group.UsersUser.UserEntityID,
// 		&group.UsersUser.UserName,
// 		&group.UsersUser.UserPassword,
// 		&group.UsersUser.UserFirstName,
// 		&group.UsersUser.UserLastName,
// 		&group.UsersUser.UserBirthDate,
// 		&group.UsersUser.UserEmailPromotion,
// 		&group.UsersUser.UserDemographic,
// 		&group.UsersUser.UserModifiedDate,
// 		&group.UsersUser.UserPhoto,
// 		&group.UsersUser.UserCurrentRole,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &group, nil
// }
