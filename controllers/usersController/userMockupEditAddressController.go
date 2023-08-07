package usersController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"codeid.revampacademy/services/usersService"
	"github.com/gin-gonic/gin"
)

type EditAddressController struct {
	userAddressService *usersService.AddAddressService
}

// Declare constructor
func NewEditAddressController(userAddressService *usersService.AddAddressService) *EditAddressController {
	return &EditAddressController{
		userAddressService: userAddressService,
	}
}

func (addAddressController EditAddressController) AddAddress(ctx *gin.Context) {

	addressId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create Address request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var addresParams *dbContext.AddAddressParams
	err = json.Unmarshal(body, &addresParams)
	if err != nil {
		log.Println("Error while unmarshaling create User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var addres2Params *dbContext.AddressParams
	err = json.Unmarshal(body, &addres2Params)
	if err != nil {
		log.Println("Error while unmarshaling create User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := addAddressController.userAddressService.AddAddress(ctx, addresParams, addres2Params, int32(addressId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}