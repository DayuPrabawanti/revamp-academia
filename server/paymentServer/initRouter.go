package server

import (
	controllers "codeid.revampacademy/controllers/paymentControllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *controllers.ControllersManager) *gin.Engine {
	router := gin.Default()

	paymentRoute := router.Group("/api/fintech")

	// router (API) end-point Mockup 1
	paymentRoute.GET("/bank", controllerMgr.PaymentBankController.GetListPaymentBank)
	paymentRoute.GET("/bank/search", controllerMgr.PaymentBankController.GetPaymentBankByName)
	paymentRoute.POST("/bank/create", controllerMgr.PaymentBankController.CreateNewPaymentBank)

	paymentRoute.PUT("/bank/update/:id", controllerMgr.PaymentBankController.UpdatePaymentBank)
	paymentRoute.DELETE("/bank/delete/:id", controllerMgr.PaymentBankController.DeletePaymentBank)

	// Router (API) end-point Mockup 2
	paymentRoute.GET("/fintech", controllerMgr.PaymentFintechController.GetListPaymentFintech)
	paymentRoute.GET("/fintech/search", controllerMgr.PaymentFintechController.GetPaymentFintechByName)
	paymentRoute.POST("/fintech/payment/create", controllerMgr.PaymentFintechController.CreateNewPaymentFintech)

	paymentRoute.PUT("/fintech/payment/update/:id", controllerMgr.PaymentFintechController.UpdatePaymentFintechById)
	paymentRoute.DELETE("/fintech/payment/delete/:id", controllerMgr.PaymentFintechController.DeletePaymentFintechById)

	// Router (API) end-point Mockup 3
	paymentRoute.GET("/accounts", controllerMgr.PaymentAccountController.GetListPaymentAccount)
	paymentRoute.GET("/accounts/search", controllerMgr.PaymentAccountController.GetPaymentAccountByName)
	paymentRoute.POST("/accounts/payment/create", controllerMgr.PaymentAccountController.CreateNewPaymentAccount)

	paymentRoute.PUT("/accounts/payment/update/:id", controllerMgr.PaymentAccountController.UpdatePaymentAccountById)
	paymentRoute.DELETE("/accounts/payment/delete/:id", controllerMgr.PaymentAccountController.DeletePaymentAccountById)

	// Router (API) end-point Mockup 4
	paymentRoute.GET("/topup", controllerMgr.PaymentTopupController.GetListTopupDetail)

	// router (API) end-point Mockup 5
	paymentRoute.GET("/transaction", controllerMgr.PaymentTransactionController.GetListPaymentTransaction)
	paymentRoute.GET("/transaction/view", controllerMgr.PaymentTransactionController.GetPaymentTransactionById)
	paymentRoute.POST("/transaction/create", controllerMgr.PaymentTransactionController.CreateNewPaymentTransaction)

	paymentRoute.PUT("/transaction/update/:id", controllerMgr.PaymentTransactionController.UpdatePaymentTransaction)
	paymentRoute.DELETE("/transaction/delete/:id", controllerMgr.PaymentTransactionController.DeletePaymentTransaction)

	return router
}
