package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type GabungRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewGabungRepository(dbHandler *sql.DB) *GabungRepository {
	return &GabungRepository{
		dbHandler: dbHandler,
	}
}

func (gb GabungRepository) GetListGabung(ctx *gin.Context) ([]*models.Gabung, *models.ResponseError) {

	store := dbContext.New(gb.dbHandler)
	Gabung, err := store.Gabung(ctx)

	ListGabung := make([]*models.Gabung, 0)

	for _, v := range Gabung {
		Gabung := &models.Gabung{
			UsersUser:               v.UsersUser,
			UsersUsersEmail:         v.UsersUsersEmail,
			UsersUsersEducation:     v.UsersUsersEducation,
			UsersPhoneNumberType:    v.UsersPhoneNumberType,
			CurriculumProgramEntity: v.CurriculumProgramEntity,
			BootcampProgramApply:    v.BootcampProgramApply,
		}
		ListGabung = append(ListGabung, Gabung)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return ListGabung, nil
}
