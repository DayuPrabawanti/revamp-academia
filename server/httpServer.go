package server

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config                   *viper.Viper
	router                   *gin.Engine
	paymentFintechController *controllers.PaymentFintechController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	paymentFintechRepository := repositories.NewPaymentFintechRepository(dbHandler)

	paymentFintechService := services.NewPaymentFintechService(paymentFintechRepository)

	paymentFintechController := controllers.NewPaymentFintechController(paymentFintechService)

	paymentAccountRepository := repositories.NewPaymentAccountRepository(dbHandler)

	paymentAccountService := services.NewPaymentAccountService(paymentAccountRepository)

	paymentAccountController := controllers.NewPaymentAccountController(paymentAccountService)

	router := gin.Default()

	// Router (API) end-point Mockup 2
	router.GET("/api/fintech/fintech", paymentFintechController.GetListPaymentFintech)
	router.GET("/api/fintech/fintech/search", paymentFintechController.GetPaymentFintechByName)
	router.POST("/api/fintech/fintech/payment/create", paymentFintechController.CreateNewPaymentFintech)

	router.PUT("/api/fintech/fintech/payment/update/:id", paymentFintechController.UpdatePaymentFintechById)
	router.DELETE("/api/fintech/fintech/payment/delete/:id", paymentFintechController.DeletePaymentFintechById)

	// Router (API) end-point Mockup 3
	router.GET("/api/fintech/accounts", paymentAccountController.GetListPaymentAccount)
	router.GET("/api/fintech/accounts/search", paymentAccountController.GetPaymentAccountByName)
	router.POST("/api/fintech/accounts/payment/create", paymentAccountController.CreateNewPaymentAccount)

	router.PUT("/api/fintech/accounts/payment/update/:id", paymentAccountController.UpdatePaymentAccountById)
	router.DELETE("/api/fintech/accounts/payment/delete/:id", paymentAccountController.DeletePaymentAccountById)

	// Router (API) end-point Mockup 4
	router.GET("/api/fintech/topup", paymentAccountController.GetListPaymentAccount)

	return HttpServer{
		config:                   config,
		router:                   router,
		paymentFintechController: paymentFintechController,
	}
}

// Running gin HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
