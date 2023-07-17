package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type SalesRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSalesRepository(dbHandler *sql.DB) *SalesRepository {
	return &SalesRepository{
		dbHandler: dbHandler,
	}
}

func (cr SalesRepository) GetListSales(ctx *gin.Context) ([]*models.SalesSpecialOffer, *models.ResponseError) {

	store := dbcontext.New(cr.dbHandler)
	special_offer, err := store.ListSpecial_offer(ctx)

	listSalesOffer := make([]*models.SalesSpecialOffer, 0)

	for _, v := range special_offer {
		sales := &models.SalesSpecialOffer{
			SpofID:           v.SpofID,
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
		listSalesOffer = append(listSalesOffer, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSalesOffer, nil
}
