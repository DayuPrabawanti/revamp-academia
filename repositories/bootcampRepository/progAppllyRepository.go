package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type ProgAppllyRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgAppllyRepository(dbHandler *sql.DB) *ProgAppllyRepository {
	return &ProgAppllyRepository{
		dbHandler: dbHandler,
	}
}

func (pro ProgAppllyRepository) GetListProgApply(ctx *gin.Context) ([]*models.BootcampProgramApply, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgramApplly, err := store.GetlistProgApply(ctx)

	ProgApplly := make([]*models.BootcampProgramApply, 0)

	for _, v := range ProgramApplly {
		ProgramApplly := &models.BootcampProgramApply{
			PrapUserEntityID: v.PrapUserEntityID,
			PrapProgEntityID: v.PrapProgEntityID,
			PrapTestScore:    v.PrapTestScore,
			PrapGpa:          v.PrapGpa,
			PrapIqTest:       v.PrapIqTest,
			PrapReview:       v.PrapReview,
			PrapModifiedDate: v.PrapModifiedDate,
			PrapStatus:       v.PrapStatus,
		}
		ProgApplly = append(ProgApplly, ProgramApplly)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return ProgApplly, nil
}

func (pro ProgAppllyRepository) UpdateProgApply(ctx *gin.Context, ProgApply *dbContext.ProgApply) *models.ResponseError {

	store := dbContext.New(pro.dbHandler)
	err := store.UpdateProgApply(ctx, *ProgApply)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}
func (pro ProgAppllyRepository) GetProgApply(ctx *gin.Context, id int64) (*dbContext.ProgApply, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	department, err := store.GetProgApply(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &department, nil
}

func (pro ProgAppllyRepository) GetTestScore(ctx *gin.Context, id int64) (*dbContext.UpdateStatus, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgApply, err := store.GetTestScore(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &ProgApply, nil
}

func (pro ProgAppllyRepository) GetReview(ctx *gin.Context, id int64) (*dbContext.UpdateStatus, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgApply, err := store.GetReview(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &ProgApply, nil
}

func (pro ProgAppllyRepository) GetStatus(ctx *gin.Context, id int64) (*dbContext.UpdateStatus, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgApply, err := store.GetStatus(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &ProgApply, nil
}

func (pro ProgAppllyRepository) GetlistProgApplyStatus(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgramApplly, err := store.GetlistProgApplyStatus(ctx, status)

	ProgApplly := make([]*dbContext.ProgApplyParams, 0)

	for _, v := range ProgramApplly {
		ProgramApplly := &dbContext.ProgApplyParams{
			PrapUserEntityID: v.PrapUserEntityID,
			PrapProgEntityID: v.PrapProgEntityID,
			PrapTestScore:    v.PrapTestScore,
			PrapGpa:          v.PrapGpa,
			PrapIqTest:       v.PrapIqTest,
			PrapReview:       v.PrapReview,
			PrapModifiedDate: v.PrapModifiedDate,
			PrapStatus:       v.PrapStatus,
		}
		ProgApplly = append(ProgApplly, ProgramApplly)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return ProgApplly, nil
}

func (pro ProgAppllyRepository) GetlistProgApplyfiltering(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgramApplly, err := store.GetlistProgApplyfiltering(ctx, status)

	ProgApplly := make([]*dbContext.ProgApplyParams, 0)

	for _, v := range ProgramApplly {
		ProgramApplly := &dbContext.ProgApplyParams{
			PrapUserEntityID: v.PrapUserEntityID,
			PrapProgEntityID: v.PrapProgEntityID,
			PrapTestScore:    v.PrapTestScore,
			PrapGpa:          v.PrapGpa,
			PrapIqTest:       v.PrapIqTest,
			PrapReview:       v.PrapReview,
			PrapModifiedDate: v.PrapModifiedDate,
			PrapStatus:       v.PrapStatus,
		}
		ProgApplly = append(ProgApplly, ProgramApplly)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return ProgApplly, nil
}

func (pro ProgAppllyRepository) GetlistProgApplycontract(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgramApplly, err := store.GetlistProgApplycontract(ctx, status)

	ProgApplly := make([]*dbContext.ProgApplyParams, 0)

	for _, v := range ProgramApplly {
		ProgramApplly := &dbContext.ProgApplyParams{
			PrapUserEntityID: v.PrapUserEntityID,
			PrapProgEntityID: v.PrapProgEntityID,
			PrapTestScore:    v.PrapTestScore,
			PrapGpa:          v.PrapGpa,
			PrapIqTest:       v.PrapIqTest,
			PrapReview:       v.PrapReview,
			PrapModifiedDate: v.PrapModifiedDate,
			PrapStatus:       v.PrapStatus,
		}
		ProgApplly = append(ProgApplly, ProgramApplly)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return ProgApplly, nil
}

func (pro ProgAppllyRepository) GetlistProgApplyfailed(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgramApplly, err := store.GetlistProgApplyfailed(ctx, status)

	ProgApplly := make([]*dbContext.ProgApplyParams, 0)

	for _, v := range ProgramApplly {
		ProgramApplly := &dbContext.ProgApplyParams{
			PrapUserEntityID: v.PrapUserEntityID,
			PrapProgEntityID: v.PrapProgEntityID,
			PrapTestScore:    v.PrapTestScore,
			PrapGpa:          v.PrapGpa,
			PrapIqTest:       v.PrapIqTest,
			PrapReview:       v.PrapReview,
			PrapModifiedDate: v.PrapModifiedDate,
			PrapStatus:       v.PrapStatus,
		}
		ProgApplly = append(ProgApplly, ProgramApplly)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return ProgApplly, nil
}

func (pro ProgAppllyRepository) GetlistProgApplyidle(ctx *gin.Context, status string) ([]*dbContext.ProgApplyParams, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgramApplly, err := store.GetlistProgApplyidle(ctx, status)

	ProgApplly := make([]*dbContext.ProgApplyParams, 0)

	for _, v := range ProgramApplly {
		ProgramApplly := &dbContext.ProgApplyParams{
			PrapUserEntityID: v.PrapUserEntityID,
			PrapProgEntityID: v.PrapProgEntityID,
			PrapTestScore:    v.PrapTestScore,
			PrapGpa:          v.PrapGpa,
			PrapIqTest:       v.PrapIqTest,
			PrapReview:       v.PrapReview,
			PrapModifiedDate: v.PrapModifiedDate,
			PrapStatus:       v.PrapStatus,
		}
		ProgApplly = append(ProgApplly, ProgramApplly)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return ProgApplly, nil
}

func (pro ProgAppllyRepository) UpdateTestScore(ctx *gin.Context, UpdateStatus *dbContext.UpdateStatus, id int32) *models.ResponseError {

	store := dbContext.New(pro.dbHandler)
	err := store.UpdateTestScore(ctx, *UpdateStatus, id)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (pro ProgAppllyRepository) UpdatePrapStatus(ctx *gin.Context, UpdateStatus *dbContext.UpdateStatus, id int32) *models.ResponseError {

	store := dbContext.New(pro.dbHandler)
	err := store.UpdatePrapStatus(ctx, *UpdateStatus, id)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (pro ProgAppllyRepository) UpdatePrapReview(ctx *gin.Context, UpdateStatus *dbContext.UpdateStatus, id int32) *models.ResponseError {

	store := dbContext.New(pro.dbHandler)
	err := store.UpdatePrapReview(ctx, *UpdateStatus, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}
func (pro ProgAppllyRepository) GetlistModifiedDate(ctx *gin.Context, metadata *features.Metadata) ([]*dbContext.ProgApplyParams, *models.ResponseError) {

	store := dbContext.New(pro.dbHandler)
	ProgramApplly, err := store.GetlistModifiedDate(ctx, metadata)

	ProgApplly := make([]*dbContext.ProgApplyParams, 0)

	for _, v := range ProgramApplly {
		ProgramApplly := &dbContext.ProgApplyParams{
			PrapUserEntityID: v.PrapUserEntityID,
			PrapProgEntityID: v.PrapProgEntityID,
			PrapTestScore:    v.PrapTestScore,
			PrapGpa:          v.PrapGpa,
			PrapIqTest:       v.PrapIqTest,
			PrapReview:       v.PrapReview,
			PrapModifiedDate: v.PrapModifiedDate,
			PrapStatus:       v.PrapStatus,
		}
		ProgApplly = append(ProgApplly, ProgramApplly)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return ProgApplly, nil
}
