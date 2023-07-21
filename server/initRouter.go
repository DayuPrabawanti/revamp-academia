package server

import (
	"codeid.revampacademy/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *controllers.ControllersManager) *gin.Engine {
	router := gin.Default()

	paymentRoute := router.Group("api/fintech")

	// router (API) end-point Mockup 1
	paymentRoute.GET("/bank", controllerMgr.PaymentBankController.GetListPaymentBank)
	paymentRoute.GET("/bank/search", controllerMgr.PaymentBankController.GetPaymentBankByName)
	paymentRoute.POST("/bank/create", controllerMgr.PaymentBankController.CreateNewPaymentBank)

	paymentRoute.PUT("/bank/update/:id", controllerMgr.PaymentBankController.UpdatePaymentBank)
	paymentRoute.DELETE("/bank/delete/:id", controllerMgr.PaymentBankController.DeletePaymentBank)

	// router (API) end-point Mockup 5
	paymentRoute.GET("/transaction", controllerMgr.PaymentTransactionController.GetListPaymentTransaction)
	paymentRoute.GET("/transactio/search", controllerMgr.PaymentTransactionController.GetPaymentTransactionById)
	paymentRoute.POST("/transaction/create", controllerMgr.PaymentTransactionController.CreateNewPaymentTransaction)

	paymentRoute.PUT("/transaction/update/:id", controllerMgr.PaymentTransactionController.UpdatePaymentTransaction)
	paymentRoute.DELETE("/transaction/delete/:id", controllerMgr.PaymentTransactionController.DeletePaymentTransaction)

	return router

}
