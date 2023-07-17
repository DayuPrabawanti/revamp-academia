package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcamp/dbContext"
	"github.com/gin-gonic/gin"
)

// BATCH
type BatchRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewBatchRepository(dbHandler *sql.DB) *BatchRepository {
	return &BatchRepository{
		dbHandler: dbHandler,
	}
}

func (br BatchRepository) GetListBatch(ctx *gin.Context) ([]*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batchs, err := store.ListBatchs(ctx)

	listBatchs := make([]*models.BootcampBatch, 0)

	for _, v := range batchs {
		batch := &models.BootcampBatch{
			BatchID:           v.BatchID,
			BatchEntityID:     v.BatchEntityID,
			BatchName:         v.BatchName,
			BatchDescription:  v.BatchDescription,
			BatchStartDate:    v.BatchStartDate,
			BatchEndDate:      v.BatchEndDate,
			BatchReason:       v.BatchReason,
			BatchType:         v.BatchType,
			BatchModifiedDate: v.BatchModifiedDate,
			BatchStatus:       v.BatchStatus,
			BatchPicID:        v.BatchPicID,
		}
		listBatchs = append(listBatchs, batch)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBatchs, nil
}

func (br BatchRepository) GetBatch(ctx *gin.Context, id int64) (*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batch, err := store.GetBatch(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &batch, nil
}

func (br BatchRepository) CreateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) (*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batch, err := store.CreateBatch(ctx, *batchParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return batch, nil
}

func (br BatchRepository) UpdateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.UpdateBatch(ctx, *batchParams)

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

func (br BatchRepository) DeleteBatch(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.DeleteBatch(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}

// BATCH TRAINEE
type BatchTraineeRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewBatchTraineeRepository(dbHandler *sql.DB) *BatchTraineeRepository {
	return &BatchTraineeRepository{
		dbHandler: dbHandler,
	}
}

func (btr BatchTraineeRepository) GetListBatchTrainee(ctx *gin.Context) ([]*models.BootcampBatchTrainee, *models.ResponseError) {

	store := dbContext.New(btr.dbHandler)
	batchTrainees, err := store.ListBatchTrinee(ctx)

	listBatchTrinee := make([]*models.BootcampBatchTrainee, 0)

	for _, v := range batchTrainees {
		batchTrainee := &models.BootcampBatchTrainee{
			BatrID:               v.BatrID,
			BatrStatus:           v.BatrStatus,
			BatrCertificated:     v.BatrCertificated,
			BatreCertificateLink: v.BatreCertificateLink,
			BatrAccessToken:      v.BatrAccessToken,
			BatrAccessGrant:      v.BatrAccessGrant,
			BatrReview:           v.BatrReview,
			BatrTotalScore:       v.BatrTotalScore,
			BatrModifiedDate:     v.BatrModifiedDate,
			BatrTraineeEntityID:  v.BatrTraineeEntityID,
			BatrBatchID:          v.BatrBatchID,
		}
		listBatchTrinee = append(listBatchTrinee, batchTrainee)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBatchTrinee, nil
}

// BATCH TRAINEE EVALUATION
type BatchTraineeEvaluationRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewBatchTraineeEvaluationRepository(dbHandler *sql.DB) *BatchTraineeEvaluationRepository {
	return &BatchTraineeEvaluationRepository{
		dbHandler: dbHandler,
	}
}

func (bter BatchTraineeEvaluationRepository) GetListBatchTraineeEvaluation(ctx *gin.Context) ([]*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {

	store := dbContext.New(bter.dbHandler)
	batchTraineeEvs, err := store.ListBatchTraineeEvaluations(ctx)

	listBatchTraineeEvaluations := make([]*models.BootcampBatchTraineeEvaluation, 0)

	for _, v := range batchTraineeEvs {
		batchTraineeEv := &models.BootcampBatchTraineeEvaluation{
			BtevID:              v.BtevID,
			BtevType:            v.BtevType,
			BtevHeader:          v.BtevHeader,
			BtevSection:         v.BtevSection,
			BtevSkill:           v.BtevSkill,
			BtevWeek:            v.BtevWeek,
			BtevSkor:            v.BtevSkor,
			BtevNote:            v.BtevNote,
			BtevModifiedDate:    v.BtevModifiedDate,
			BtevBatchID:         v.BtevBatchID,
			BtevTraineeEntityID: v.BtevTraineeEntityID,
		}
		listBatchTraineeEvaluations = append(listBatchTraineeEvaluations, batchTraineeEv)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBatchTraineeEvaluations, nil
}

// PROGRAM APPLY
type ProgramApplyRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgramApplyRepository(dbHandler *sql.DB) *ProgramApplyRepository {
	return &ProgramApplyRepository{
		dbHandler: dbHandler,
	}
}

func (par ProgramApplyRepository) GetListProgramApply(ctx *gin.Context) ([]*models.BootcampProgramApply, *models.ResponseError) {

	store := dbContext.New(par.dbHandler)
	progApplies, err := store.ListProgramApplies(ctx)

	listProgramApplies := make([]*models.BootcampProgramApply, 0)

	for _, v := range progApplies {
		programApply := &models.BootcampProgramApply{
			PrapUserEntityID: v.PrapUserEntityID,
			PrapProgEntityID: v.PrapProgEntityID,
			PrapTestScore:    v.PrapTestScore,
			PrapGpa:          v.PrapGpa,
			PrapIqTest:       v.PrapIqTest,
			PrapReview:       v.PrapReview,
			PrapModifiedDate: v.PrapModifiedDate,
			PrapStatus:       v.PrapStatus,
		}
		listProgramApplies = append(listProgramApplies, programApply)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgramApplies, nil
}

// PROGRAM APPLY PROGRESS
type ProgramApplyProgressRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgramApplyProgressRepository(dbHandler *sql.DB) *ProgramApplyProgressRepository {
	return &ProgramApplyProgressRepository{
		dbHandler: dbHandler,
	}
}

func (papr ProgramApplyProgressRepository) GetListProgramApplyProgress(ctx *gin.Context) ([]*models.BootcampProgramApplyProgress, *models.ResponseError) {

	store := dbContext.New(papr.dbHandler)
	progApplyProgresses, err := store.ListProgramApplyProgresses(ctx)

	listProgramApplyProgresses := make([]*models.BootcampProgramApplyProgress, 0)

	for _, v := range progApplyProgresses {
		programApplyProgress := &models.BootcampProgramApplyProgress{
			ParogID:           v.ParogID,
			ParogUserEntityID: v.ParogUserEntityID,
			ParogProgEntityID: v.ParogProgEntityID,
			ParogActionDate:   v.ParogActionDate,
			ParogModifiedDate: v.ParogModifiedDate,
			ParogComment:      v.ParogComment,
			ParogProgressName: v.ParogProgressName,
			ParogEmpEntityID:  v.ParogEmpEntityID,
			ParogStatus:       v.ParogStatus,
		}
		listProgramApplyProgresses = append(listProgramApplyProgresses, programApplyProgress)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgramApplyProgresses, nil
}

// INSTRUCTOR PROGRAMS
type InstructorProgramRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewInstructorProgramRepository(dbHandler *sql.DB) *InstructorProgramRepository {
	return &InstructorProgramRepository{
		dbHandler: dbHandler,
	}
}

func (ipr InstructorProgramRepository) GetListInstructorPrograms(ctx *gin.Context) ([]*models.BootcampInstructorProgram, *models.ResponseError) {

	store := dbContext.New(ipr.dbHandler)
	instructorProgs, err := store.ListInstructorPrograms(ctx)

	listInstructorPrograms := make([]*models.BootcampInstructorProgram, 0)

	for _, v := range instructorProgs {
		instructorProgram := &models.BootcampInstructorProgram{
			BatchID:           v.BatchID,
			InproEntityID:     v.InproEntityID,
			InproEmpEntityID:  v.InproEmpEntityID,
			InproModifiedDate: v.InproModifiedDate,
		}
		listInstructorPrograms = append(listInstructorPrograms, instructorProgram)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listInstructorPrograms, nil
}

// func (br BatchRepository) GetBatchTraineeEvaluation(ctx *gin.Context, id int64) (*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {

// 	store := dbContext.New(br.dbHandler)
// 	btev, err := store.GetBatchTraineeEvaluation(ctx, int32(id))

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return &btev, nil
// }

// func (br BatchRepository) CreateBatchTraineeEvaluation(ctx *gin.Context, batchtev *dbContext.CreateBatchTraineeEvaluationParams) (*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {

// 	store := dbContext.New(br.dbHandler)
// 	btev, err := store.CreateBatchTraineeEvaluation(ctx, *batchtev)

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Message,
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return btev, nil
// }

// func (br BatchRepository) UpdateBatchTraineeEvaluation(ctx *gin.Context, batchtevParams *dbContext.CreateBatchParams) *models.ResponseError {

// 	store := dbContext.New(br.dbHandler)
// 	err := store.UpdateBatch(ctx, *batchtevParams)

// 	if err != nil {
// 		return &models.ResponseError{
// 			Message: "error when update",
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return &models.ResponseError{
// 		Message: "data has been update",
// 		Status:  http.StatusOK,
// 	}
// }

// func (br BatchRepository) DeleteBatchTraineeEvaluation(ctx *gin.Context, id int64) *models.ResponseError {

// 	store := dbContext.New(br.dbHandler)
// 	err := store.DeleteBatchTraineeEvaluation(ctx, int32(id))

// 	if err != nil {
// 		return &models.ResponseError{
// 			Message: "error when update",
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return &models.ResponseError{
// 		Message: "data has been deleted",
// 		Status:  http.StatusOK,
// 	}
// }
