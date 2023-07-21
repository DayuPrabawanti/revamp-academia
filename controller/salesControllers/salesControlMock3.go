package salescontrollers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	saler "codeid.revampacademy/service/salesService"
	"github.com/gin-gonic/gin"
)

type ControlMock3 struct {
	serviceMock3 *saler.ServiceMock3
}

func NewMockupApplyController(serviceMock3 *saler.ServiceMock3) *ControlMock3 {
	return &ControlMock3{
		serviceMock3: serviceMock3,
	}
}

func (controlMock3 ControlMock3) CreateMergeUsers(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create Gabung request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var mergeParam dbcontext.CreateMergeMock
	err = json.Unmarshal(body, &mergeParam)
	if err != nil {
		log.Println("Error while unmarshaling create Gabung request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := controlMock3.serviceMock3.CreateMergeMocks(ctx, &mergeParam)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controlMock3 ControlMock3) CreateUsers(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create Users request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var userParam dbcontext.CreateUsersParams
	err = json.Unmarshal(body, &userParam)
	if err != nil {
		log.Println("Error while unmarshaling create Users request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := controlMock3.serviceMock3.CreateUsers(ctx, &userParam)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controlMock3 ControlMock3) CreateEducation(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create Education request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var educationParam dbcontext.CreateEducationParams
	err = json.Unmarshal(body, &educationParam)
	if err != nil {
		log.Println("Error while unmarshaling create Education request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := controlMock3.serviceMock3.CreateEducations(ctx, &educationParam)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controlMock3 ControlMock3) CreateMedian(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create media request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var mediaParam dbcontext.CreateMediaParams
	err = json.Unmarshal(body, &mediaParam)
	if err != nil {
		log.Println("Error while unmarshaling create media request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := controlMock3.serviceMock3.CreateMedian(ctx, &mediaParam)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controlMock3 ControlMock3) GetUsers(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, responseErr := controlMock3.serviceMock3.GetUsers(ctx, int64(userId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controlMock3 ControlMock3) GetListGroup(ctx *gin.Context) {
	response, responerr := controlMock3.serviceMock3.ListUserGroup(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
