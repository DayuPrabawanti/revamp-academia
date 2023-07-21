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
