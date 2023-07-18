package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type SpecialOfferRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSpecialOfferRepository(dbHandler *sql.DB) *SpecialOfferRepository {
	return &SpecialOfferRepository{
		dbHandler: dbHandler,
	}
}

func (cr SpecialOfferRepository) ListSpecial_offer(ctx *gin.Context) ([]*models.SalesSpecialOffer, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	specialOffer, err := store.ListSpecial_offer(ctx)

	listSpecialOffer := make([]*models.SalesSpecialOffer, 0)

	for _, v := range specialOffer {
		specialOffer := &models.SalesSpecialOffer{
			SpofDescription:  v.SpofDescription,
			SpofDiscount:     v.SpofDiscount,
			SpofType:         v.SpofType,
			SpofStartDate:    v.SpofStartDate,
			SpofEndDate:      v.SpofEndDate,
			SpofMinQty:       v.SpofMinQty,
			SpofMaxQty:       v.SpofMaxQty,
			SpofModifiedDate: v.SpofModifiedDate,
			SpofCateID:       v.SpofCateID,
		}
		listSpecialOffer = append(listSpecialOffer, specialOffer)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSpecialOffer, nil
}

func (cr SpecialOfferRepository) GetSpecialOffer(ctx *gin.Context, id int64) (*models.SalesSpecialOffer, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	specialoffer, err := store.GetSpecial_offer(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &specialoffer, nil
}

func (cr SpecialOfferRepository) CreateSpecialOffer(ctx *gin.Context, specialofferParams *dbContext.CreateSales_special_offerParams) (*models.SalesSpecialOffer, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	specialoffer, err := store.CreateSales_special_offer(ctx, *specialofferParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return specialoffer, nil
}
