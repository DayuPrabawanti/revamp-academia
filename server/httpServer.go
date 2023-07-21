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
	config             *viper.Viper
	router             *gin.Engine
	ControllersManager controllers.ControllersManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	repositoriesManager := repositories.NewRepositoriesManager(dbHandler)

	ServicesManager := services.NewServicesManager(repositoriesManager)

	controllersManager := controllers.NewControllersManager(ServicesManager)

	router := InitRouter(controllersManager)

	return HttpServer{
		config:             config,
		router:             router,
		ControllersManager: *controllersManager,
	}
}

// Running gin HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}

// paymentBankRepository := repositories.NewPaymentBankRepository(dbHandler)
// paymentBankService := services.NewPaymentBankService(paymentBankRepository)
// paymentBankController := controllers.NewPaymentBankController(paymentBankService)

// paymentTransactioRepository := repositories.NewPaymentTransactionRepository(dbHandler)
// paymentTransactionService := services.NewPaymentTransactionService(paymentTransactioRepository)
// paymentTransactionController := controllers.NewPaymentTransactionController(paymentTransactionService)

// router := gin.Default()

// // mockuop 1:bank
// router.GET("/api/fintech/bank", paymentBankController.GetListPaymentBank)
// router.GET("/api/fintech/bank/search", paymentBankController.GetPaymentBankByName)
// router.POST("/api/fintech/bank/create", paymentBankController.CreateNewPaymentBank)
// router.PUT("/api/fintech/bank/update/:id", paymentBankController.UpdatePaymentBank)
// router.DELETE("api/fintech/bank/delete/:id", paymentBankController.DeletePaymentBank)

// // mockuop 7:transacation
// router.GET("/api/fintech/transaction", paymentTransactionController.GetListPaymentTransaction)
// router.GET("api/fintech/transaction/search", paymentTransactionController.GetPaymentTransactionById)
// router.POST("api/fintech/transaction/create", paymentTransactionController.CreateNewPaymentTransaction)
// router.PUT("api/fintech/transaction/update/:id", paymentTransactionController.UpdatePaymentTransaction)
// router.DELETE("api/fintech/transaction/delete/:id", paymentTransactionController.DeletePaymentTransaction)
// return HttpServer{
// 	config:                config,
// 	router:                router,
// 	paymentBankController: paymentBankController,
// }
