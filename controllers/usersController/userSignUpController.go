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

type SignUpEmployeeController struct {
	signUpEmployeeService *usersService.SignUpEmployeeService
}

// Declare constructor
func NewSignUpController(signUpService *usersService.SignUpService) *SignUpController {
	return &SignUpController{
		signUpService: signUpService,
	}
}

func NewSignUpEmployeeController(signUpEmployeeService *usersService.SignUpEmployeeService) *SignUpEmployeeController {
	return &SignUpEmployeeController{
		signUpEmployeeService: signUpEmployeeService,
	}
}

func (signupcontroller SignUpController) CreateSignUpUser(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create signup request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user dbContext.CreateUsersParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var email dbContext.CreateEmailParams
	err = json.Unmarshal(body, &email)
	if err != nil {
		log.Println("Error while unmarshaling Email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var phone dbContext.CreatePhonesParams
	err = json.Unmarshal(body, &phone)
	if err != nil {
		log.Println("Error while unmarshaling Phone request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := signupcontroller.signUpService.SignUpUser(ctx, &user, &email, &phone)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (signupEmployeecontroller SignUpEmployeeController) CreateSignUpEmployee(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create signup request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user dbContext.CreateUsersParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var email dbContext.CreateEmailParams
	err = json.Unmarshal(body, &email)
	if err != nil {
		log.Println("Error while unmarshaling Email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var phone dbContext.CreatePhonesParams
	err = json.Unmarshal(body, &phone)
	if err != nil {
		log.Println("Error while unmarshaling Phone request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := signupEmployeecontroller.signUpEmployeeService.SignUpEmplloyee(ctx, &user, &email, &phone)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}