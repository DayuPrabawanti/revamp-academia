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

type UserAddressController struct {
	userAddressService *usersService.UserAddressService
}

// Declare constructor
func NewUseraddressController(userAddressService *usersService.UserAddressService) *UserAddressController {
	return &UserAddressController{
		userAddressService: userAddressService,
	}
}

func (userAddressController UserAddressController) GetListUserAddress(ctx *gin.Context){

	response, responseErr := userAddressController.userAddressService.GetListUserAddress(ctx)

	if responseErr != nil{
		ctx.JSON(responseErr.Status,responseErr)
		return 
	}

	ctx.JSON(http.StatusOK,response)
	
	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}

func (userAddressController UserAddressController) GetAddress(ctx *gin.Context) {

	userAddr, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := userAddressController.userAddressService.GetAddress(ctx, int32(userAddr))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (userAddresscontroller UserAddressController) CreateAddrees(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create address request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var userAddr dbContext.CreateAddreesParams
	err = json.Unmarshal(body, &userAddr)
	if err != nil {
		log.Println("Error while unmarshaling create address request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := userAddresscontroller.userAddressService.CreateAddrees(ctx, &userAddr)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}