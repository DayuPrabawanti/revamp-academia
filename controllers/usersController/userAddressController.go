package usersController

import (
	"net/http"

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
