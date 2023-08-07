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

type UserEmailController struct {
	userEmailService *usersService.UserEmailService
}

// Declare constructor
func NewUserEmailController(userEmailService *usersService.UserEmailService) *UserEmailController {
	return &UserEmailController{
		userEmailService: userEmailService,
	}
}

func (userEmailController UserEmailController) GetListUsersEmail(ctx *gin.Context){

	response, responseErr := userEmailController.userEmailService.GetListUsersEmail(ctx)

	if responseErr != nil{
		ctx.JSON(responseErr.Status,responseErr)
		return 
	}

	ctx.JSON(http.StatusOK,response)
	
	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}

func (userEmailController UserEmailController) GetEmail(ctx *gin.Context) {

	emailId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := userEmailController.userEmailService.GetEmail(ctx, int32(emailId))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (useremailcontroller UserEmailController) CreateEmail(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var email dbContext.CreateEmailParams
	err = json.Unmarshal(body, &email)
	if err != nil {
		log.Println("Error while unmarshaling create email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := useremailcontroller.userEmailService.CreateEmail(ctx, &email)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (userEmailController UserEmailController) UpdateEmail(ctx *gin.Context) {

	emailId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user dbContext.CreateEmailParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling update email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := userEmailController.userEmailService.UpdateEmail(ctx, &user, int64(emailId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (userEmailController UserEmailController) DeleteEmail(ctx *gin.Context) {

	emailId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := userEmailController.userEmailService.DeleteEmail(ctx, int32(emailId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// func (userEmailController UserEmailController) EmailProfile(ctx *gin.Context) {

// 	emailId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

// 	if err != nil {
// 		log.Println("Error while reading paramater id", err)
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	response, responseErr := userEmailController.userEmailService.EmailProfile(ctx, int32(emailId))
// 	if responseErr != nil {

// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }