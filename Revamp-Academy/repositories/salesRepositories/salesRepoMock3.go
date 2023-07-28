package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type RepoMockup3 struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewMockupApplyRepo(dbHandler *sql.DB) *RepoMockup3 {
	return &RepoMockup3{
		dbHandler: dbHandler,
	}
}

func (rm RepoMockup3) CreateUser(ctx *gin.Context, userParams *dbcontext.CreateUsersParams) (*models.UsersUser, *models.ResponseError) {
	store := dbcontext.New(rm.dbHandler)
	user, err := store.CreateUsersParams(ctx, *userParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return user, nil
}
func (rm RepoMockup3) CreateEducation(ctx *gin.Context, eductionParams *dbcontext.CreateEducationParam) (*models.UsersUsersEducation, *models.ResponseError) {
	store := dbcontext.New(rm.dbHandler)
	education, err := store.CreateEducationParams(ctx, *eductionParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return education, nil
}

func (rm RepoMockup3) CreateMedia(ctx *gin.Context, mediaParams *dbcontext.CreateMediaParams) (*models.UsersUsersMedium, *models.ResponseError) {
	store := dbcontext.New(rm.dbHandler)
	media, err := store.CreateMediaParams(ctx, *mediaParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return media, nil
}

func (rm RepoMockup3) CreateMergeMock(ctx *gin.Context, mergeParams *dbcontext.CreateMergeMock) (*models.MergeApplyProgress, *models.ResponseError) {
	user, err := rm.CreateUser(ctx, &mergeParams.CreateUsersParams)
	if err != nil {
		return nil, err
	}
	education, err := rm.CreateEducation(ctx, &mergeParams.CreateEducationParam)
	if err != nil {
		return nil, err
	}
	media, err := rm.CreateMedia(ctx, &mergeParams.CreateMediaParams)
	if err != nil {
		return nil, err
	}

	merge := &models.MergeApplyProgress{
		Users:     *user,
		Education: *education,
		Media:     *media,
	}
	return merge, nil
}
func (rm RepoMockup3) GetUsers(ctx *gin.Context, id int64) (*models.UsersUser, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetUsersId(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}

func (rm RepoMockup3) ListUserGroup(ctx *gin.Context) ([]*dbcontext.CreateMergeMock, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	userGroup, err := store.ListUserGroup(ctx)

	listUserGrup := make([]*dbcontext.CreateMergeMock, 0)

	for _, v := range userGroup {
		sales := &dbcontext.CreateMergeMock{
			CreateUsersParams:    v.CreateUsersParams,
			CreateEducationParam: v.CreateEducationParam,
			CreateMediaParams:    v.CreateMediaParams,
		}
		listUserGrup = append(listUserGrup, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUserGrup, nil
}

func (rm RepoMockup3) ListBootcampApplyProgressRepo(ctx *gin.Context) ([]*models.MergeBatchApplyProgress, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	userGroup, err := store.ListApplyProgressMock6(ctx)

	listUserGrup := make([]*models.MergeBatchApplyProgress, 0)

	for _, v := range userGroup {
		sales := &models.MergeBatchApplyProgress{
			ProgramApply:         v.ProgramApply,
			ProgramApplyProgress: v.ProgramApplyProgress,
		}
		listUserGrup = append(listUserGrup, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUserGrup, nil
}
