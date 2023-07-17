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
	config                *viper.Viper
	router                *gin.Engine
	paymentBankController *controllers.PaymentBankController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	paymentBankRepository := repositories.NewPaymentBankRepository(dbHandler)

	paymentBankService := services.NewPaymentBankService(paymentBankRepository)

	paymentBankController := controllers.NewPaymentBankController(paymentBankService)

	router := gin.Default()

	router.GET("/api/fintech/bank", paymentBankController.GetListPaymentBank)
	router.GET("/api/fintech/bank/search", paymentBankController.GetPaymentBankByName)

	router.POST("/api/fintech/bank/create", paymentBankController.CreateNewPaymentBank)
	router.PUT("/api/fintech/bank/update/:id", paymentBankController.UpdatePaymentBank)
	router.DELETE("api/fintech/bank/delete/:id", paymentBankController.DeletePaymentBank)

	return HttpServer{
		config:                config,
		router:                router,
		paymentBankController: paymentBankController,
	}
}

// Running gin HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
