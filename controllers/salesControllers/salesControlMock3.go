package salesControllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	saler "codeid.revampacademy/services/salesService"
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
