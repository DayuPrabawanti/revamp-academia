package curriculumRepositories

import (
	"net/http"

	mod "codeid.revampacademy/models"
	"codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (ped ProgEntityRepository) GetProgEntityDesc(ctx *gin.Context, id int64) (*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {

	store := db.New(ped.dbHandler)
	programEntityDescription, err := store.Getprogram_entity_description(ctx, int32(id))

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return (*mod.CurriculumProgramEntityDescription)(&programEntityDescription), nil
}

func (ped ProgEntityRepository) CreateProgEntityDesc(ctx *gin.Context, programEntityDescriptionParams *db.CreateProgEntityDescParams) (*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {
	store := dbContext.New(ped.dbHandler)
	progEntityDesc, err := store.CreateProgEntityDesc(ctx, *programEntityDescriptionParams)
	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return progEntityDesc, nil
}

func (per ProgEntityRepository) UpdateProgEntityDesc(ctx *gin.Context, progEntityDescParams *db.UpdateProgEntityDescParams) *mod.ResponseError {

	store := dbContext.New(per.dbHandler)
	err := store.UpdateProgEntityDesc(ctx, *progEntityDescParams)

	if err != nil {
		return &mod.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (per ProgEntityRepository) DeleteProgEntityDesc(ctx *gin.Context, id int64) *mod.ResponseError {

	store := dbContext.New(per.dbHandler)
	err := store.DeleteProgEntityDesc(ctx, int32(id))

	if err != nil {
		return &mod.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
