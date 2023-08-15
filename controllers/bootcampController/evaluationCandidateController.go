package bootcampController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type EvaluationCandidateController struct {
	evaluationCandidateService *bootcampService.EvaluationCandidateService
}

// declare constructor
func NewEvaluationCandidateController(evaluationCandidateService *bootcampService.EvaluationCandidateService) *EvaluationCandidateController {
	return &EvaluationCandidateController{
		evaluationCandidateService: evaluationCandidateService,
	}
}

func (evaluationCandidateController EvaluationCandidateController) GetEvaluationCandidate(ctx *gin.Context) {
	id := ctx.Query("userentityid") // Mengambil nilai query parameter id dari URL

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user entity id is required"})
		return
	}

	userentityid, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while parsing id", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user entity id"})
		return
	}

	response, responseErr := evaluationCandidateController.evaluationCandidateService.GetEvaluationCandidate(ctx, int32(userentityid))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (evaluationCandidateController EvaluationCandidateController) CreateEvaluationCandidate(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create evaluation candidate request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var evaluationCandidate dbContext.CreateEvaluationCandidateParams
	err = json.Unmarshal(body, &evaluationCandidate)
	if err != nil {
		log.Println("Error while unmarshaling create evaluation candidate request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := evaluationCandidateController.evaluationCandidateService.CreateEvaluationCandidate(ctx, &evaluationCandidate)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (evaluationCandidateController EvaluationCandidateController) UpdateEvaluationCandidate(ctx *gin.Context) {
	userentityid, err := strconv.Atoi(ctx.Query("userentityid")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var evaluationCandidate dbContext.UpdateEvaluationCandidateParams
	err = json.Unmarshal(body, &evaluationCandidate)
	if err != nil {
		log.Println("Error while unmarshaling update batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := evaluationCandidateController.evaluationCandidateService.UpdateEvaluationCandidate(ctx, &evaluationCandidate, int32(userentityid))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// 	id := ctx.Query("userentityid") // Mengambil nilai query parameter id dari URL

// 	if id == "" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user entity id is required"})
// 		return
// 	}

// 	userentityid, err := strconv.Atoi(id)
// 	if err != nil {
// 		log.Println("Error while parsing id", err)
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user entity id"})
// 		return
// 	}

// 	response, responseErr := evaluationCandidateController.evaluationCandidateService.UpdateEvaluationCandidate(ctx, int32(userentityid))
// 	if responseErr != nil {
// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }
