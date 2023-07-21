package curriculumRepositories

import (
	"database/sql"
	"net/http"

	mod "codeid.revampacademy/models"
	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

// PROGRAM ENTITY DESCRIPTION

type ProgEntityDescRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgEntityDescRepository(dbHandler *sql.DB) *ProgEntityDescRepository {
	return &ProgEntityDescRepository{
		dbHandler: dbHandler,
	}
}

func (ped ProgEntityDescRepository) GetListProgEntityDesc(ctx *gin.Context) ([]*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {

	store := db.New(ped.dbHandler)
	progEntityDesc, err := store.Listprogram_entity_description(ctx)

	listProgEntityDesc := make([]*mod.CurriculumProgramEntityDescription, 0)

	for _, v := range progEntityDesc {
		progEntityDesc := &mod.CurriculumProgramEntityDescription{
			PredProgEntityID: v.PredProgEntityID,
			PredItemLearning: v.PredItemLearning,
			PredItemInclude:  v.PredItemInclude,
			PredRequirment:   v.PredRequirment,
			PredDescription:  v.PredDescription,
			PredTargetLevel:  v.PredTargetLevel,
		}
		listProgEntityDesc = append(listProgEntityDesc, progEntityDesc)
	}

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgEntityDesc, nil
}

func (ped ProgEntityDescRepository) GetProgEntityDesc(ctx *gin.Context, id int64) (*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {

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
