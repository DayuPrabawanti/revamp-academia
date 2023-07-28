package curriculumServices

import (
	"net/http"

	mod "codeid.revampacademy/models"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type CurriculumService struct {
	curriculumRepository *repo.CurriculumRepository
}

func NewCurriculumService(curriculumRepository *repo.CurriculumRepository) *CurriculumService {
	return &CurriculumService{
		curriculumRepository: curriculumRepository,
	}
}

func (cs CurriculumService) GetCurriculum(ctx *gin.Context, id int64) (*[]mod.CurriculumAll, *mod.ResponseError) {
	return cs.curriculumRepository.GetCurriculum(ctx, id)
}

func (cs CurriculumService) UpdateProgramEntity(ctx *gin.Context, programentityParams *db.UpdateprogramentityParams, id int64) *mod.ResponseError {
	responseErr := validateProgramEntity(programentityParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.curriculumRepository.UpdateProgramEntity(ctx, programentityParams)
}

func (cs CurriculumService) Updateprogramentitydescription(ctx *gin.Context, programentitydescParams *db.UpdateprogramentitydescriptionParams, id int64) *mod.ResponseError {
	responseErr := validateProgramEntityDesc(programentitydescParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.curriculumRepository.Updateprogramentitydescription(ctx, programentitydescParams)
}

func (cs CurriculumService) UpdateScore(ctx *gin.Context, scoreParams *db.UpdateScoreParams, id int64) *mod.ResponseError {
	responseErr := validateScore(scoreParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.curriculumRepository.UpdateScore(ctx, scoreParams)
}

// func (cs CurriculumService) UpdateCurriculum(ctx *gin.Context, curriculumParams *db.UpdateCurriculumParams, id int64) *mod.ResponseError {

//		return cs.curriculumRepository.UpdateCurriculum(ctx, curriculumParams)
//	}
func (cs CurriculumService) UpdateCurriculum(ctx *gin.Context, updategabungParams *db.UpdateCurriculum) *mod.ResponseError {
	return cs.curriculumRepository.UpdateCurriculum(ctx, updategabungParams)
}

func validateProgramEntity(progentityParams *db.UpdateprogramentityParams) *mod.ResponseError {
	if progentityParams.ProgEntityID == 0 {
		return &mod.ResponseError{
			Message: "Invalid Program Entity id",
			Status:  http.StatusBadRequest,
		}
	}

	if progentityParams.ProgTitle == "" {
		return &mod.ResponseError{
			Message: "Invalid Program Entity Title",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func validateProgramEntityDesc(progentityParams *db.UpdateprogramentitydescriptionParams) *mod.ResponseError {
	if progentityParams.PredProgEntityID == 0 {
		return &mod.ResponseError{
			Message: "Invalid Program Entity Description id",
			Status:  http.StatusBadRequest,
		}
	}

	if progentityParams.PredItemLearning.String == "" {
		return &mod.ResponseError{
			Message: "Invalid Program Entity Description Learning",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
func validateScore(scoreParams *db.UpdateScoreParams) *mod.ResponseError {
	if scoreParams.SectProgEntityID == 0 {
		return &mod.ResponseError{
			Message: "Invalid Sect Program Enitty ID id",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
