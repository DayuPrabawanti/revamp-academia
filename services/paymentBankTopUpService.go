package services

// import (
// 	"net/http"

// 	"codeid.revampacademy/models"
// 	"codeid.revampacademy/repositories"
// 	"codeid.revampacademy/repositories/dbContext"
// 	"github.com/gin-gonic/gin"
// )

// type FintechTopupService struct {
// 	fintechTopupRepository *repositories.FintechTopupRepository
// }

// func NewFintechTopupRepository(FintechTopupRepository *repositories.FintechTopupRepository) *FintechTopupService {
// 	return &FintechTopupService{
// 		fintechTopupRepository: fintechTopupRepository,
// 	}
// }

// // create payment TopUp
// func (pbr FintechTopupService) CreateFintechTopup(ctx *gin.Context, fintechTopupParams *dbContext.CreateFintechTopupParams) (*models.FintechTopup, *models.ResponseError) {
// 	responseErr := validateFintechTopup(fintechTopupParams)
// 	if responseErr != nil {
// 		return nil, responseErr
// 	}
// 	return pbr.fintechTopupRepository.CreateFintechTopup(ctx, fintechTopupParams)
// }

// func validateFintechTopup(fintechTopupParams *dbContext.CreateFintechTopupParams) *models.ResponseError {
// 	if fintechTopupParams.FintEntityID == 0 {
// 		return &models.ResponseError{
// 			Message: "Invalid bankentity id",
// 			Status:  http.StatusBadRequest,
// 		}
// 	}

// 	if fintechTopupParams.FintName == "" {
// 		return &models.ResponseError{
// 			Message: "Invalid category name",
// 			Status:  http.StatusBadRequest,
// 		}
// 	}

// 	return nil

// }
