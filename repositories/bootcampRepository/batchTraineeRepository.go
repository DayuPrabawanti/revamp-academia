package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

// BATCH TRAINEE
type BatchTraineeRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewBatchTraineeRepository(dbHandler *sql.DB) *BatchTraineeRepository {
	return &BatchTraineeRepository{
		dbHandler: dbHandler,
	}
}

func (btr BatchTraineeRepository) GetListBatchTrainee(ctx *gin.Context) ([]*models.BootcampBatchTrainee, *models.ResponseError) {

	store := dbContext.New(btr.dbHandler)
	batchTrainees, err := store.ListBatchTrinee(ctx)

	listBatchTrinee := make([]*models.BootcampBatchTrainee, 0)

	for _, v := range batchTrainees {
		batchTrainee := &models.BootcampBatchTrainee{
			BatrID:               v.BatrID,
			BatrStatus:           v.BatrStatus,
			BatrCertificated:     v.BatrCertificated,
			BatreCertificateLink: v.BatreCertificateLink,
			BatrAccessToken:      v.BatrAccessToken,
			BatrAccessGrant:      v.BatrAccessGrant,
			BatrReview:           v.BatrReview,
			BatrTotalScore:       v.BatrTotalScore,
			BatrModifiedDate:     v.BatrModifiedDate,
			BatrTraineeEntityID:  v.BatrTraineeEntityID,
			BatrBatchID:          v.BatrBatchID,
		}
		listBatchTrinee = append(listBatchTrinee, batchTrainee)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBatchTrinee, nil
}

// GetBatchTrainee retrieves a single batch trainee from the database by its ID.
func (btr BatchTraineeRepository) GetBatchTrainee(ctx *gin.Context, id int64) (*models.BootcampBatchTrainee, *models.ResponseError) {
	store := dbContext.New(btr.dbHandler)
	batchTrainee, err := store.GetBatchTrainee(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &batchTrainee, nil
}

// CreateBatchTrainee creates a new batch trainee in the database.
// func (btr BatchTraineeRepository) CreateBatchTrainee(ctx *gin.Context, batchTrainee *dbContext.CreateBatchTraineeParams) (*models.BootcampBatchTrainee, *models.ResponseError) {

// 	store := dbContext.New(btr.dbHandler)
// 	batchTrainees, err := store.CreateBatchTrainee(ctx, *batchTraineeParams)

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Message,
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return batchTrainees, nil
// }

// // UpdateBatchTrainee updates an existing batch trainee in the database.
// func (btr BatchTraineeRepository) UpdateBatchTrainee(ctx *gin.Context, batchTraineeParams *dbContext.CreateBatchTraineeParams) *models.ResponseError {
// 	store := dbContext.New(btr.dbHandler)
// 	err := store.UpdateBatchTrainee(ctx, *batchTraineeParams)

// 	if err != nil {
// 		return &models.ResponseError{
// 			Message: "error when updating",
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return &models.ResponseError{
// 		Message: "data has been updated",
// 		Status:  http.StatusOK,
// 	}
// }

// // DeleteBatchTrainee deletes a batch trainee from the database.
// func (btr BatchTraineeRepository) DeleteBatchTrainee(ctx *gin.Context, id int64) *models.ResponseError {
// 	store := dbContext.New(btr.dbHandler)
// 	err := store.DeleteBatchTrainee(ctx, int32(id))

// 	if err != nil {
// 		return &models.ResponseError{
// 			Message: "error when deleting",
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return &models.ResponseError{
// 		Message: "data has been deleted",
// 		Status:  http.StatusOK,
// 	}
