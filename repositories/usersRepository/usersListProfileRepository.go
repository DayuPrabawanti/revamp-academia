package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserListProfileRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserListProfileRepository(dbHandler *sql.DB) *UserListProfileRepository {
	return &UserListProfileRepository{
		dbHandler: dbHandler,
	}
}

func (cr UserListProfileRepository) GetProfileUser(ctx *gin.Context, id int32) (*dbContext.UsernameParams, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	user, err := store.GetProfileUser(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &user, nil
}

func (cr UserListProfileRepository) GetProfileEmail(ctx *gin.Context, id int32) ([]*dbContext.ProfileEmailParams, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersEmail, err := store.GetProfileEmail(ctx, int32(id))

	PrfileEmail := make([]*dbContext.ProfileEmailParams, 0)

	for _, v := range usersEmail {
		userEmail := &dbContext.ProfileEmailParams{
			PmailAddress: v.PmailAddress,
		}
		PrfileEmail = append(PrfileEmail, userEmail)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return PrfileEmail, nil
}

func (cr UserListProfileRepository) GetProfile(ctx *gin.Context, id int32) ([]*dbContext.TampilProfile, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	users, err := store.GetProfile(ctx, int32(id))

	Prfile := make([]*dbContext.TampilProfile, 0)

	for _, v := range users {
		user := &dbContext.TampilProfile{
			UserFirstName: v.UserFirstName,
			UserLastName: v.UserLastName,
			UserFullname: v.UserFullname,
			UserPhoto: v.UserPhoto,
			PmailAddress: v.PmailAddress,
			UspoNumber: v.UspoNumber,
			UspoPontyCode: v.UspoPontyCode,
			AddrLine1: v.AddrLine1,
			AddrLine2: v.AddrLine2,
			AddrPostalCode: v.AddrPostalCode,
			CityName: v.CityName,
			UsduSchool: v.UsduSchool,
			UsduDegree: v.UsduDegree,
			UsduFieldStudy: v.UsduFieldStudy,
			UsduGrade: v.UsduGrade,
			UsduStartDate: v.UsduStartDate,
			UsduEndDate: v.UsduEndDate,
			UsduActivities: v.UsduActivities,
			UsduDescription: v.UsduDescription,
			UsexTitle: v.UsexTitle,
			UsexProfileHeadline: v.UsexProfileHeadline,
			UsexCompanyName: v.UsexCompanyName,
			UsexStartDate: v.UsexStartDate,
			UsexEndDate: v.UsexEndDate,
			UsexCityID: v.UsexCityID,
			UsexDescription: v.UsexDescription,
			UskiSktyName: v.UskiSktyName,
		}
		Prfile = append(Prfile, user)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return Prfile, nil
}