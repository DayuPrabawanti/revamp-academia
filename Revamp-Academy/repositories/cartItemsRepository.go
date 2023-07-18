package repositories

import (
	// "database/sql"
	// "net/http"

	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
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

	store := dbContext.New(cr.dbHandler)
	cart_items, err := store.Getcart_items(ctx, caitID)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &cart_items, nil
}
