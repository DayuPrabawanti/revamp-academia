package repositories

// import (
// 	"database/sql"
// 	"net/http"

// 	"codeid.revampacademy/models"
// 	"codeid.revampacademy/repositories/dbContext"
// 	"github.com/gin-gonic/gin"
// )

// type FintechTopupRepository struct {
// 	dbHandler   *sql.DB
// 	transaction *sql.Tx
// }

// func NewFintechTopupRepository(dbHandler *sql.DB) *FintechTopupRepository {
// 	return &FintechTopupRepository{
// 		dbHandler: dbHandler,
// 	}
// }

// // getlist dari data tabel bank
// func (pbr FintechTopupRepository) GetListPaymnetBankRepository(ctx *gin.Context) ([]*models.PaymentBankRepository, *models.ResponseError) {
// 	store := dbContext.New(pbr.dbHandler)
// 	paymentBank, err := store.ListPaymentBank(ctx)
// 	listPaymentBank := make([]*models.PaymentBankRepository, 0)

// 	for _, v := range paymentBank {
// 		paymentBank := &models.PaymentBankRepository{
// 			BankEntityID:     v.BankEntityID,
// 			BankCode:         v.BankCode,
// 			BankName:         v.BankName,
// 			BankModifiedDate: v.BankModifiedDate,
// 		}
// 		listPaymentBank = append(listPaymentBank, paymentBank)
// 	}
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return listPaymentBank, nil
// }

// // getlist dari data tabel fintec
// func (pbr FintechTopupRepository) GetListFintechTopupRepository(ctx *gin.Context) ([]*models.FintechTopup, *models.ResponseError) {
// 	store := dbContext.New(pbr.dbHandler)
// 	fintechTopup, err := store.ListFintechTopup(ctx)
// 	listFintechTopup := make([]*models.FintechTopupRepository, 0)

// 	for _, v := range fintechTopup {
// 		fintechTopups := &models.FintechTopupRepository{
// 			FintEntityID:     v.FintEntityID,
// 			FintCode:         v.FintCode,
// 			FintName:         v.FintName,
// 			FintModifiedDate: v.FintModifiedDate,
// 		}
// 		listFintechTopup = append(listFintechTopup, fintechTopups)
// 	}

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return listFintechTopup, nil
// }

// // getlist dati data tabel user_account
// func (pbr FintechTopupRepository) GetlistUserAccountRepository(ctx *gin.Context) ([]*models.UsersAccount, *models.ResponseError) {
// 	store := dbContext.New(pbr.dbHandler)
// 	userAccountSection, err := store.listUserAccounts(ctx)
// 	listUsersAccount := make([]*models.UsersAccount, 0)

// 	for _, v := range userAccountSection {
// 		userAccountSections := &models.UsersAccountRepository{
// 			UsacBankEntityID:  v.UsacBankEntityID,
// 			UsacUserEntityID:  v.UsacUserEntityID,
// 			UsacAccountNumber: v.UsacAccountNumber,
// 			UsacSaldo:         v.UsacSaldo,
// 			UsacType:          v.UsacType,
// 			UsacStartDate:     v.UsacStartDate,
// 			UsacEndDate:       v.UsacEndDate,
// 			UsacModifiedDate:  v.UsacModifiedDate,
// 			UsacStatus:        v.UsacStatus,
// 		}
// 		listUsersAccount = append(listUsersAccount, userAccountSections)
// 	}
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return listUsersAccount, nil
// }

// // function gabungan dari 3 tabel bank,fintech,user_account
// func (pbr FintechTopupRepository) Gabung(ctx *gin.Context) ([]*models.Gabung, *models.ResponseError) {
// 	// ambil data dari tabel bank
// 	paymentBank, err := pbr.GetListPaymnetBankRepository(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// ambil data dari tabel fintech
// 	fintechTopup, err := pbr.GetListFintechTopupRepository(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// ambil data dari tabel user_account
// 	usersAccounts, err := pbr.GetlistUserAccountRepository(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// buat map utk menyimpan FintechTopup berdasarkan User_account
// 	fintechTopupMap := make(map[int64]*models.FintechTopup)
// 	for _, fintechTopup := range fintechTopups {
// 		fintechTopupMap[fintechTopup.FintEntityID] = fintechTopup
// 	}

// 	userAccountSectionMap := make(map[int64]*models.UserAccount)
// 	for _, userAccount := range userAccounts {
// 		userAccountSectionMap[userAccount.UsacBankEntityID] = userAccount
// 	}

// 	// gabungkan data dari fintech dan useraccount
// 	gabungList := make([]*models.Gabung, 0)
// 	for _, fintechTopup := range fintechTopupMap {

// 	}

// }

// // 1.a fungsi utk ambil create topup
// func (pbr FintechTopupRepository) CreateFintechTopup(ctx *gin.Context, fintechTopupParams *dbContext.CreateFintechTopupParams) (*models.FintechTopup, *models.ResponseError) {
// 	store := dbContext.New(pbr.dbHandler)
// 	fintechTopup, err := store.CreateFintechTopup(ctx, *fintechTopupParams)

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Message,
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return fintechTopup, nil
// }
