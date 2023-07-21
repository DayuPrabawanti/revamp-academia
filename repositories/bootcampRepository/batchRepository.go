package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

// BATCH
type BatchRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewBatchRepository(dbHandler *sql.DB) *BatchRepository {
	return &BatchRepository{
		dbHandler: dbHandler,
	}
}

func (br BatchRepository) GetListBatch(ctx *gin.Context) ([]*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batchs, err := store.ListBatchs(ctx)

	listBatchs := make([]*models.BootcampBatch, 0)

	for _, v := range batchs {
		batch := &models.BootcampBatch{
			BatchID:           v.BatchID,
			BatchEntityID:     v.BatchEntityID,
			BatchName:         v.BatchName,
			BatchDescription:  v.BatchDescription,
			BatchStartDate:    v.BatchStartDate,
			BatchEndDate:      v.BatchEndDate,
			BatchReason:       v.BatchReason,
			BatchType:         v.BatchType,
			BatchModifiedDate: v.BatchModifiedDate,
			BatchStatus:       v.BatchStatus,
			BatchPicID:        v.BatchPicID,
		}
		listBatchs = append(listBatchs, batch)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBatchs, nil
}

func (br BatchRepository) GetListBatchTraineeEvaluation(ctx *gin.Context) ([]*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batchTraineeEvs, err := store.ListBatchTraineeEvaluations(ctx)

	listBatchTraineeEvaluations := make([]*models.BootcampBatchTraineeEvaluation, 0)

	for _, v := range batchTraineeEvs {
		batchTraineeEv := &models.BootcampBatchTraineeEvaluation{
			BtevID:              v.BtevID,
			BtevType:            v.BtevType,
			BtevHeader:          v.BtevHeader,
			BtevSection:         v.BtevSection,
			BtevSkill:           v.BtevSkill,
			BtevWeek:            v.BtevWeek,
			BtevSkor:            v.BtevSkor,
			BtevNote:            v.BtevNote,
			BtevModifiedDate:    v.BtevModifiedDate,
			BtevBatchID:         v.BtevBatchID,
			BtevTraineeEntityID: v.BtevTraineeEntityID,
		}
		listBatchTraineeEvaluations = append(listBatchTraineeEvaluations, batchTraineeEv)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBatchTraineeEvaluations, nil
}

func (br BatchRepository) GetListProgramEntity(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	programEntity, err := store.ListProgramEntity(ctx)

	listProgramEntity := make([]*models.CurriculumProgramEntity, 0)

	for _, v := range programEntity {
		programEntity := &models.CurriculumProgramEntity{
			ProgEntityID:     v.ProgEntityID,
			ProgTitle:        v.ProgTitle,
			ProgHeadline:     v.ProgHeadline,
			ProgType:         v.ProgType,
			ProgLearningType: v.ProgLearningType,
			ProgRating:       v.ProgRating,
			ProgBestSeller:   v.ProgBestSeller,
			ProgPrice:        v.ProgPrice,
			ProgLanguage:     v.ProgLanguage,
			ProgModifiedDate: v.ProgModifiedDate,
			ProgDuration:     v.ProgDuration,
			ProgDurationType: v.ProgDurationType,
			ProgTagSkill:     v.ProgTagSkill,
			ProgCityID:       v.ProgCityID,
			ProgCateID:       v.ProgCateID,
			ProgCreatedBy:    v.ProgCreatedBy,
			ProgStatus:       v.ProgStatus,
		}
		listProgramEntity = append(listProgramEntity, programEntity)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgramEntity, nil
}

func (br BatchRepository) GetListUsers(ctx *gin.Context) ([]*models.UsersUser, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	users, err := store.ListUsers(ctx)

	listUsers := make([]*models.UsersUser, 0)

	for _, v := range users {
		user := &models.UsersUser{
			UserEntityID:       v.UserEntityID,
			UserName:           v.UserName,
			UserPassword:       v.UserPassword,
			UserFirstName:      v.UserFirstName,
			UserLastName:       v.UserLastName,
			UserBirthDate:      v.UserBirthDate,
			UserEmailPromotion: v.UserEmailPromotion,
			UserDemographic:    v.UserDemographic,
			UserModifiedDate:   v.UserModifiedDate,
			UserPhoto:          v.UserPhoto,
			UserCurrentRole:    v.UserCurrentRole,
		}
		listUsers = append(listUsers, user)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsers, nil
}

func (br BatchRepository) GetBatch(ctx *gin.Context, id int64) (*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batch, err := store.GetBatch(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &batch, nil
}

func (br BatchRepository) CreateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) (*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batch, err := store.CreateBatch(ctx, *batchParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return batch, nil
}

func (br BatchRepository) UpdateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.UpdateBatch(ctx, *batchParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) DeleteBatch(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.DeleteBatch(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}

// BATCH TRAINEE EVALUATION
// type BatchTraineeEvaluationRepository struct {
// 	dbHandler   *sql.DB
// 	transaction *sql.Tx
// }

// func NewBatchTraineeEvaluationRepository(dbHandler *sql.DB) *BatchTraineeEvaluationRepository {
// 	return &BatchTraineeEvaluationRepository{
// 		dbHandler: dbHandler,
// 	}
// }

// type GroupRepository struct {
// 	dbHandler *sql.DB
// }

// func NewGroupRepository(dbHandler *sql.DB) *GroupRepository {
// 	return &GroupRepository{
// 		dbHandler: dbHandler,
// 	}
// }

// func (gr *GroupRepository) GetGroupByID(batchID int32) (*models.Group, *models.ResponseError) {
// 	query := `
// 		SELECT bb.batch_id, bb.batch_entity_id, bb.batch_name, bb.batch_description,
// 			bb.batch_start_date, bb.batch_end_date, bb.batch_reason, bb.batch_type,
// 			bb.batch_modified_date, bb.batch_status, bb.batch_pic_id,
// 			btev.btev_id, btev.btev_type, btev.btev_header, btev.btev_section,
// 			btev.btev_skill, btev.btev_week, btev.btev_skor, btev.btev_note,
// 			btev.btev_modified_date, btev.btev_trainee_entity_id,
// 			cpe.prog_entity_id, cpe.prog_title, cpe.prog_headline, cpe.prog_type,
// 			cpe.prog_learning_type, cpe.prog_rating, cpe.prog_total_trainee,
// 			cpe.prog_modified_date, cpe.prog_image, cpe.prog_best_seller,
// 			cpe.prog_price, cpe.prog_language, cpe.prog_duration,
// 			cpe.prog_duration_type, cpe.prog_tag_skill,
// 			uu.user_entity_id, uu.user_name, uu.user_password, uu.user_first_name,
// 			uu.user_last_name, uu.user_birth_date, uu.user_email_promotion,
// 			uu.user_demographic, uu.user_modified_date, uu.user_photo,
// 			uu.user_current_role
// 		FROM bootcamp_batch bb
// 		JOIN bootcamp_batch_trainee_evaluation btev ON bb.batch_id = btev.btev_batch_id
// 		JOIN curriculum_program_entity cpe ON btev.btev_batch_id = cpe.prog_entity_id
// 		JOIN users_user uu ON btev.btev_trainee_entity_id = uu.user_entity_id
// 		WHERE bb.batch_id = $1
// 	`

// 	row := gr.dbHandler.QueryRow(query, batchID)

// 	group := &models.Group{}
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
// 		&group.BootcampBatchTraineeEvaluation.BtevModifiedDate,
// 		&group.BootcampBatchTraineeEvaluation.BtevTraineeEntityID,
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
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, &models.ResponseError{
// 				Message: "Group not found",
// 				Status:  http.StatusNotFound,
// 			}
// 		}
// 		return nil, &models.ResponseError{
// 			Message: "Error while querying group",
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return group, nil
// }

// func (br BatchRepository) GetBatchEvaluationView(ctx *gin.Context, batchID int64) (*models.BootcampBatchEvaluationView, *models.ResponseError) {
// 	store := dbContext.New(br.dbHandler)
// 	batch, err := store.GetBatch(ctx, int32(batchID))
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	traineeEvaluation, err := store.GetBatchTraineeEvaluation(ctx, int32(batchID))
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	// curriculumProgramEntity, err := store.GetCurriculumProgramEntity(ctx, traineeEvaluation.EntityID)
// 	// if err != nil {
// 	// 	return nil, &models.ResponseError{
// 	// 		Message: err.Error(),
// 	// 		Status:  http.StatusInternalServerError,
// 	// 	}
// 	// }

// 	// user, err := store.GetUser(ctx, traineeEvaluation.UserID)
// 	// if err != nil {
// 	// 	return nil, &models.ResponseError{
// 	// 		Message: err.Error(),
// 	// 		Status:  http.StatusInternalServerError,
// 	// 	}
// 	// }

// 	// Populate struct BootcampBatchEvaluationView
// 	evaluationView := &models.BootcampBatchEvaluationView{
// 		BootcampBatch:                  *batch,
// 		BootcampBatchTraineeEvaluation: *traineeEvaluation,
// 		// CurriculumProgramEntity:        *curriculumProgramEntity,
// 		// UsersUser:                      *user,
// 	}

// 	return evaluationView, nil
// }
