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
	dbQueries   dbContext.Queries
}

func NewBootcampBatchEvaluationRepository(dbHandler *sql.DB) *BootcampBatchEvaluationRepository {
	return &BootcampBatchEvaluationRepository{
		dbHandler: dbHandler,
		//add new fields
		dbQueries: *dbContext.New(dbHandler),
	}
}

func (ber BootcampBatchEvaluationRepository) GetBootcampBatchEvaluation(ctx *gin.Context, id int32) ([]*dbContext.BootcampBatchEvaluationMockup, *models.ResponseError) {

	// store := dbContext.New(cr.dbHandler)
	bootcampEvaluations, err := ber.dbQueries.GetBootcampBatchEvaluation(ctx, int32(id))

	bootcampBatch := make([]*dbContext.BootcampBatchEvaluationMockup, 0)

	for _, v := range bootcampEvaluations {
		bootcamp := &dbContext.BootcampBatchEvaluationMockup{
			UserEntityID:        v.UserEntityID,
			BatchID:             v.BatchID,
			BatchName:           v.BatchName,
			ProgTitle:           v.ProgTitle,
			UserPhoto:           v.UserPhoto,
			UserFullname:        v.UserFullname,
			BatrStatus:          v.BatrStatus,
			BtevSkor:            v.BtevSkor,
			BatchEntityID:       v.BatchEntityID,
			BtevTraineeEntityID: v.BtevTraineeEntityID,
		}
		bootcampBatch = append(bootcampBatch, bootcamp)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return bootcampBatch, nil
}

func (ber BootcampBatchEvaluationRepository) GetBatchTraineeReview(ctx *gin.Context, id int32) ([]*dbContext.BootcampBatchTraineeReview, *models.ResponseError) {

	// store := dbContext.New(cr.dbHandler)
	bootcampReviews, err := ber.dbQueries.GetBatchTraineeReview(ctx, int32(id))

	batchReview := make([]*dbContext.BootcampBatchTraineeReview, 0)

	for _, v := range bootcampReviews {
		bootcamp := &dbContext.BootcampBatchTraineeReview{
			UserEntityID: v.UserEntityID,
			UserFullname: v.UserFullname,
			BatrID:       v.BatrID,
			BatrStatus:   v.BatrStatus,
			BatrReview:   v.BatrReview,
		}
		batchReview = append(batchReview, bootcamp)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return batchReview, nil
}

// CreateBatchTrainee creates a new batch trainee in the database.
func (ber BootcampBatchEvaluationRepository) CreateBatchTraineeReview(ctx *gin.Context, batchTraineeReviewParams *dbContext.CreateBatchTraineeReviewParams) (*dbContext.BootcampBatchTraineeReview, *models.ResponseError) {

	// store := dbContext.New(btr.dbHandler)
	bootcampReview, err := ber.dbQueries.CreateBatchTraineeReview(ctx, *batchTraineeReviewParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return bootcampReview, nil
}

// UpdateBatchTrainee updates an existing batch trainee in the database.
func (ber BootcampBatchEvaluationRepository) UpdateBatchTraineeReview(ctx *gin.Context, batchTraineeReviewParams *dbContext.UpdateBatchTraineeReviewParams) (*dbContext.BootcampBatchTraineeReview, *models.ResponseError) {
	// store := dbContext.New(btr.dbHandler)
	bootcampReview, err := ber.dbQueries.UpdateBatchTraineeReview(ctx, *batchTraineeReviewParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return bootcampReview, nil
}

// DeleteBatchTrainee deletes a batch trainee from the database.
func (ber BootcampBatchEvaluationRepository) DeleteBatchTraineeReview(ctx *gin.Context, id int64) *models.ResponseError {
	// store := dbContext.New(btr.dbHandler)
	err := ber.dbQueries.DeleteBatchTraineeReview(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when deleting",
			Status:  http.StatusInternalServerError,
		}
	}

	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
