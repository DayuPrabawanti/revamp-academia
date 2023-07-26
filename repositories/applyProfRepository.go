package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ApplyProfRepository struct {
	DbHandler   *sql.DB
	Transaction *sql.Tx
	UsersUserRepository
	UsersEduRepository
	UsersPhoneRepository
	UsersMediaRepository
}

func NewApplyProfRepository(dbHandler *sql.DB) *ApplyProfRepository {
	return &ApplyProfRepository{
		DbHandler: dbHandler,
	}
}

func (aprof ApplyProfRepository) ListApplyProfRepo(ctx *gin.Context, nama string) ([]*models.ApplyProf, *models.ResponseError) {

	store := dbContext.New(aprof.DbHandler)
	applyProfGroup, err := store.ListApplyProfImpl(ctx, string(nama))

	applyProfMakeList := make([]*models.ApplyProf, 0)

	for _, v := range applyProfGroup {
		applyProfGroup := &models.ApplyProf{
			UserUser:      v.UserUser,
			UserEducation: v.UserEducation,
			UserPhone:     v.UserPhone,
			UserMedia:     v.UserMedia,
		}
		applyProfMakeList = append(applyProfMakeList, applyProfGroup)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return applyProfMakeList, nil
}
