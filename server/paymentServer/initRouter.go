package server

import (
	controllers "codeid.revampacademy/controllers/paymentControllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *controllers.ControllersManager) *gin.Engine {
	router := gin.Default()

	paymentRoute := router.Group("/api/payment")

	// router (API) end-point Mockup 1
	paymentRoute.GET("/bank", controllerMgr.PaymentBankController.GetListPaymentBank)
	paymentRoute.GET("/bank/search", controllerMgr.PaymentBankController.GetPaymentBankByName)
	paymentRoute.POST("/bank/create", controllerMgr.PaymentBankController.CreateNewPaymentBank)

	paymentRoute.PUT("/bank/update/:bankEntityID", controllerMgr.PaymentBankController.UpdatePaymentBank)
	paymentRoute.DELETE("/bank/delete/:id", controllerMgr.PaymentBankController.DeletePaymentBank)

	// Router (API) end-point Mockup 2
	paymentRoute.GET("/fintech", controllerMgr.PaymentFintechController.GetListPaymentFintech)
	paymentRoute.GET("/fintech/search", controllerMgr.PaymentFintechController.GetPaymentFintechByName)
	paymentRoute.POST("/fintech/create", controllerMgr.PaymentFintechController.CreateNewPaymentFintech)

	paymentRoute.PUT("/fintech/update/:fintEntityId", controllerMgr.PaymentFintechController.UpdatePaymentFintechById)
	paymentRoute.DELETE("/fintech/delete/:id", controllerMgr.PaymentFintechController.DeletePaymentFintechById)

	// Router (API) end-point Mockup 3
	paymentRoute.GET("/accounts", controllerMgr.PaymentAccountController.GetListPaymentUsers_accountByUserName)
	paymentRoute.GET("/account/search", controllerMgr.PaymentAccountController.GetPaymentAccountByAccountNumber)
	paymentRoute.POST("/account/create", controllerMgr.PaymentAccountController.CreateNewPaymentAccount)

	paymentRoute.PUT("/account/updatePlus", controllerMgr.PaymentAccountController.UpdatePaymentUsers_accountPlus)
	paymentRoute.PUT("/account/updateMinus", controllerMgr.PaymentAccountController.UpdatePaymentUsers_accountMinus)
	paymentRoute.DELETE("/account/deleteAccountNumber", controllerMgr.PaymentAccountController.DeletePaymentAccountByAccNum)
	// Transaction (Debit Saldo)
	paymentRoute.POST("/account/debitSaldo", controllerMgr.PaymentAccountController.DebitSaldo)

	// Router (API) end-point Mockup 4
	paymentRoute.GET("/topup-bank", controllerMgr.PaymentTopupController.GetAccountByBankCodeAndAccountNumber)
	paymentRoute.GET("/topup-fintech", controllerMgr.PaymentTopupController.GetAccountByFintCodeAndAccountNumber)

	// router (API) end-point Mockup 5
	paymentRoute.GET("/transaction", controllerMgr.PaymentTransactionController.GetListPaymentTransaction)
	paymentRoute.GET("/transaction/view", controllerMgr.PaymentTransactionController.GetPaymentTransactionById)
	paymentRoute.POST("/transaction/create", controllerMgr.PaymentTransactionController.CreateNewPaymentTransaction)

	// paymentRoute.PUT("/transaction/update/:id", controllerMgr.PaymentTransactionController.UpdatePaymentTransaction)
	// paymentRoute.DELETE("/transaction/delete/:id", controllerMgr.PaymentTransactionController.DeletePaymentTransaction)

	return router
}
