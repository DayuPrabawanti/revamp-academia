package salesRepositories

import (
	// "database/sql"
	// "net/http"

	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext/salesContext"
	"github.com/gin-gonic/gin"
)

type CartItemsRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewCartItemsRepository(dbHandler *sql.DB) *CartItemsRepository {
	return &CartItemsRepository{
		dbHandler: dbHandler,
	}
}

func (cr CartItemsRepository) Getcart_items(ctx *gin.Context, caitID int32) (*models.SalesCartItem, *models.ResponseError) {

	store := salesContext.New(cr.dbHandler)
	cart_items, err := store.Getcart_items(ctx, caitID)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &cart_items, nil
}

func (cr CartItemsRepository) GetListCartItems(ctx *gin.Context) ([]*models.SalesCartItem, *models.ResponseError) {

	store := salesContext.New(cr.dbHandler)
	cart_items, err := store.ListCart_item(ctx)

	listCart_item := make([]*models.SalesCartItem, 0)

	for _, v := range cart_items {
		cart_items := &models.SalesCartItem{
			CaitID:           v.CaitID,
			CaitQuantity:     v.CaitQuantity,
			CaitUnitPrice:    v.CaitUnitPrice,
			CaitModifiedDate: v.CaitModifiedDate,
			CaitUserEntityID: v.CaitUserEntityID,
			CaitProgEntityID: v.CaitProgEntityID,
		}
		listCart_item = append(listCart_item, cart_items)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listCart_item, nil
}
