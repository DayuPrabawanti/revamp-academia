package usersController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"codeid.revampacademy/services/usersService"
	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	signUpService *usersService.SignUpService
}

// Declare constructor
func NewSignUpController(signUpService *usersService.SignUpService) *SignUpController {
	return &SignUpController{
		signUpService: signUpService,
	}
}

func (signupcontroller SignUpController) CreateSignUp(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create signup request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var signup dbContext.SignUpUserParams
	err = json.Unmarshal(body, &signup)
	if err != nil {
		log.Println("Error while unmarshaling create sign request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := signupcontroller.signUpService.SignUpUser(ctx, &signup)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}