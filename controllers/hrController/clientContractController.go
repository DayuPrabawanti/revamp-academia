package hrController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"codeid.revampacademy/services/hrService"
	"github.com/gin-gonic/gin"
)

type ClientContractController struct {
	clientContractService *hrService.ClientContractService
}

// declare constructor
func NewClientContractController(clientContractService *hrService.ClientContractService) *ClientContractController {
	return &ClientContractController{
		clientContractService: clientContractService,
	}
}

// create method
func (clientContractController ClientContractController) GetListClientContract(ctx *gin.Context) {

	response, responseErr := clientContractController.clientContractService.GetListClientContract(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

func (clientContractController ClientContractController) GetClientContract(ctx *gin.Context) {

	ecco_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := clientContractController.clientContractService.GetClientContract(ctx, int64(ecco_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (clientContractController ClientContractController) CreateClientContract(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create client contract request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var clientContract dbContext.CreateClientContractParams
	err = json.Unmarshal(body, &clientContract)
	if err != nil {
		log.Println("Error while unmarshaling create client contract request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := clientContractController.clientContractService.CreateClientContract(ctx, &clientContract)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (clientContractController ClientContractController) UpdateClientContract(ctx *gin.Context) {

	ecco_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update client contract request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var clientContract dbContext.CreateClientContractParams
	err = json.Unmarshal(body, &clientContract)
	if err != nil {
		log.Println("Error while unmarshaling update client contract request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := clientContractController.clientContractService.UpdateClientContract(ctx, &clientContract, int64(ecco_id))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (clientContractController ClientContractController) DeleteClientContract(ctx *gin.Context) {

	ecco_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := clientContractController.clientContractService.DeleteClientContract(ctx, int64(ecco_id))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
