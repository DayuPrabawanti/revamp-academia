package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type TalentsMockupRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewTalentMockupRepository(dbHandler *sql.DB) *TalentsMockupRepository {
	return &TalentsMockupRepository{
		dbHandler: dbHandler,
	}
}

func (tmr TalentsMockupRepository) GetListTalentMockup(ctx *gin.Context) ([]*models.TalentsMockup, *models.ResponseError) {

	store := dbContext.New(tmr.dbHandler)
	talent, err := store.ListTalents(ctx)

	listTalent := make([]*models.TalentsMockup, 0)

	for _, v := range talent {
		talents := &models.TalentsMockup{
			HrEmployee:      v.HrEmployee,
			UsersUser:       v.UsersUser,
			UsersUsersSkill: v.UsersUsersSkill,
			BootcampBatch:   v.BootcampBatch,
			MasterSkillType: v.MasterSkillType,
		}
		listTalent = append(listTalent, talents)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listTalent, nil
}
