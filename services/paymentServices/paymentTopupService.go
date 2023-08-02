package services

import (
	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentTopupService struct {
	repositoriesManager *repositories.RepositoriesManager
}

func NewPaymentTopupService(repoMgr *repositories.RepositoriesManager) *PaymentTopupService {
	return &PaymentTopupService{
		repositoriesManager: repoMgr,
	}
}

func (ptts *PaymentTopupService) GetAccountByBankCodeAndAccountNumber(ctx *gin.Context, bankCode string, usacAccountNumber string) (*dbContext.BankAccount, *models.ResponseError) {
	return ptts.repositoriesManager.PaymentTopupRepository.GetAccountByBankCodeAndAccountNumber(ctx, bankCode, usacAccountNumber)
}

func (ptts *PaymentTopupService) GetAccountByFintCodeAndAccountNumber(ctx *gin.Context, fintCode string, usacAccountNumber string) (*dbContext.FintechAccount, *models.ResponseError) {
	return ptts.repositoriesManager.PaymentTopupRepository.GetAccountByFintCodeAndAccountNumber(ctx, fintCode, usacAccountNumber)
}

// func (pts PaymentTopupService) GetListTopupDetail(ctx *gin.Context) ([]*dbContext.TopupDetail, *models.ResponseError) {
// 	return pts.paymentTopupRepository.GetListTopupDetail(ctx)
// }

// func (pts PaymentTopupService) GetTopupDetailById(ctx *gin.Context, id int32) ([]*dbContext.TopupDetail, *models.ResponseError) {
// 	return pts.paymentTopupRepository.GetTopupDetailById(ctx, id)
// }

// func (pts PaymentTopupService) PerformTransfer(ctx *gin.Context, fromAccount string, amount float64, toAccount string, fromUserID int, toUSerID int) *models.ResponseError {
// 	// memanggil PerfomTransfer dari repository
// 	err := pts.paymentTopupRepository.PerformTransfer(ctx, fromAccount, amount, toAccount, fromUserID, toUSerID)

// 	// jika terjadi error, maka return fungsi ini
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

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
