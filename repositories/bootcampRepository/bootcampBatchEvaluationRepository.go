package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type BootcampBatchEvaluationRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewBootcampBatchEvaluationRepository(dbHandler *sql.DB) *BootcampBatchEvaluationRepository {
	return &BootcampBatchEvaluationRepository{
		dbHandler: dbHandler,
	}
}

func (ber BootcampBatchEvaluationRepository) GetListBatch(ctx *gin.Context) ([]*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(ber.dbHandler)
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

func (ber BootcampBatchEvaluationRepository) GetListBootcampBatchEvaluation(ctx *gin.Context, batchId int32) ([]*models.BootcampBatchEvaluationMockup, *models.ResponseError) {

	store := dbContext.New(ber.dbHandler)
	bootcampBatchEvs, err := store.ListBootcampBatchEvaluation(ctx, batchId)

	listBootcampBatchEvaluations := make([]*models.BootcampBatchEvaluationMockup, 0)

	for _, v := range bootcampBatchEvs {
		bootcampBatchEv := &models.BootcampBatchEvaluationMockup{
			BootcampBatch:                  v.BootcampBatch,
			BootcampBatchTraineeEvaluation: v.BootcampBatchTraineeEvaluation,
			CurriculumProgramEntity:        v.CurriculumProgramEntity,
			UsersUser:                      v.UsersUser,
		}
		listBootcampBatchEvaluations = append(listBootcampBatchEvaluations, bootcampBatchEv)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBootcampBatchEvaluations, nil
}

func (ber BootcampBatchEvaluationRepository) GetListBatchTraineeEvaluation(ctx *gin.Context) ([]*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {

	store := dbContext.New(ber.dbHandler)
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

func (ber BootcampBatchEvaluationRepository) GetListProgramEntity(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbContext.New(ber.dbHandler)
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

func (ber BootcampBatchEvaluationRepository) GetListUsers(ctx *gin.Context) ([]*models.UsersUser, *models.ResponseError) {

	store := dbContext.New(ber.dbHandler)
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
