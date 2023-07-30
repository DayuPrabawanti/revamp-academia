package services

import (
	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentTopupService struct {
	paymentTopupRepository *repositories.PaymentTopupRepository
}

func NewPaymentTopupService(paymentTopupRepository *repositories.PaymentTopupRepository) *PaymentTopupService {
	return &PaymentTopupService{
		paymentTopupRepository: paymentTopupRepository,
	}
}

func (pts PaymentTopupService) GetListTopupDetail(ctx *gin.Context) ([]*dbContext.TopupDetail, *models.ResponseError) {
	return pts.paymentTopupRepository.GetListTopupDetail(ctx)
}

func (pts PaymentTopupService) GetTopupDetailById(ctx *gin.Context, id int32) ([]*dbContext.TopupDetail, *models.ResponseError) {
	return pts.paymentTopupRepository.GetTopupDetailById(ctx, id)
}

// func (pts PaymentTopupService) CreateTopupProductDto(ctx *gin.Context, topupWithProductDto *models.CreateTopupProductDto) (*dbContext.TopupDetail, *models.ResponseError) {

// 	err := repositories.BeginTransaction(pts.repositoryManager)
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: "Failed to start transaction",
// 			Status:  http.StatusBadRequest,
// 		}
// 	}
// 	//first query statement
// 	response, responseErr := pts.CreateTopup(ctx, (*dbContext.CreateTopupParams)(&topupWithProductDto.CreateTopupDto))
// 	if responseErr != nil {
// 		repositories.RollbackTransaction(pts.repositoryManager)
// 		return nil, responseErr
// 	}
// 	//second query statement
// 	responseErr = pts.DeleteTopup(ctx, int64(response.CategoryID))
// 	if responseErr != nil {
// 		//when delete not succeed, transaction will rollback
// 		repositories.RollbackTransaction(pts.repositoryManager)
// 		return nil, responseErr
// 	}
// 	// if all statement ok, transaction will commit/save to db
// 	repositories.CommitTransaction(pts.repositoryManager)

// 	return nil, &models.ResponseError{
// 		Message: "Data has been created",
// 		Status:  http.StatusOK,
// 	}
// }
