package server

import (
	controllers "codeid.revampacademy/controllers/paymentControllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *controllers.ControllersManager) *gin.Engine {
	router := gin.Default()

	paymentRoute := router.Group("/api/fintech")

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
	paymentRoute.GET("/topup/:bankId/:fintId", controllerMgr.PaymentTopupController.GetTopupDetail)

	return router
}
