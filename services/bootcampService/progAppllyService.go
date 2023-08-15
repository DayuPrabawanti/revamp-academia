package bootcampService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type ProgAppllyService struct {
	ProgAppllyRepository *bootcampRepository.ProgAppllyRepository
}

func NewProgAppllyRepository(ProgAppllyRepository *bootcampRepository.ProgAppllyRepository) *ProgAppllyService {
	return &ProgAppllyService{
		ProgAppllyRepository: ProgAppllyRepository,
	}
}

func (pro ProgAppllyService) GetListProgApply(ctx *gin.Context) ([]*models.BootcampProgramApply, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetListProgApply(ctx)
}

func (pro ProgAppllyService) UpdateProgApply(ctx *gin.Context, updateParam *dbContext.ProgApply, id int64) *models.ResponseError {

	return pro.ProgAppllyRepository.UpdateProgApply(ctx, updateParam)
}

func validateProgApply(ProgApplly *dbContext.ProgApply) *models.ResponseError {
	if ProgApplly.PrapStatus == "" {
		return &models.ResponseError{
			Message: "Invalid Prap Status",
			Status:  http.StatusBadRequest,
		}
	}
	if ProgApplly.PrapProgEntityID == 0 {
		return &models.ResponseError{
			Message: "Required Name User",
			Status:  http.StatusBadRequest,
		}
	}
	return nil

}
func (pro ProgAppllyService) GetProgApply(ctx *gin.Context, id int64) (*dbContext.ProgApply, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetProgApply(ctx, id)
}

func (pro ProgAppllyService) GetTestScore(ctx *gin.Context, id int64) (*dbContext.UpdateStatus, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetTestScore(ctx, id)
}

func (pro ProgAppllyService) GetReview(ctx *gin.Context, id int64) (*dbContext.UpdateStatus, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetReview(ctx, id)
}

func (pro ProgAppllyService) GetStatus(ctx *gin.Context, id int64) (*dbContext.UpdateStatus, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetStatus(ctx, id)
}

func (pro ProgAppllyService) GetlistProgApplyStatus(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetlistProgApplyStatus(ctx, status)
}

func (pro ProgAppllyService) GetlistProgApplyfiltering(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetlistProgApplyfiltering(ctx, status)
}

func (pro ProgAppllyService) GetlistProgApplycontract(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetlistProgApplycontract(ctx, status)
}

func (pro ProgAppllyService) GetlistProgApplyfailed(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetlistProgApplyfailed(ctx, status)
}

func (pro ProgAppllyService) GetlistProgApplyidle(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetlistProgApplyidle(ctx, status)
}

func (pro ProgAppllyService) UpdateTestScore(ctx *gin.Context, UpdateStatus *dbContext.UpdateStatus, id int32) *models.ResponseError {
	// //responseErr := validateProgApply(UpdateStatus)
	// if responseErr != nil {
	// 	return responseErr
	// }

	return pro.ProgAppllyRepository.UpdateTestScore(ctx, UpdateStatus, id)
}

func (pro ProgAppllyService) UpdatePrapStatus(ctx *gin.Context, UpdateStatus *dbContext.UpdateStatus, id int32) *models.ResponseError {
	// //responseErr := validateProgApply(UpdateStatus)
	// if responseErr != nil {
	// 	return responseErr
	// }

	return pro.ProgAppllyRepository.UpdatePrapStatus(ctx, UpdateStatus, id)
}

func (pro ProgAppllyService) UpdatePrapReview(ctx *gin.Context, UpdateStatus *dbContext.UpdateStatus, id int32) *models.ResponseError {
	// //responseErr := validateProgApply(UpdateStatus)
	// if responseErr != nil {
	// 	return responseErr
	// }

	return pro.ProgAppllyRepository.UpdatePrapReview(ctx, UpdateStatus, id)
}

func (pro ProgAppllyService) GetlistModifiedDate(ctx *gin.Context, metadata *features.Metadata) ([]*dbContext.ProgApplyParams, *models.ResponseError) {
	return pro.ProgAppllyRepository.GetlistModifiedDate(ctx, metadata)
}
