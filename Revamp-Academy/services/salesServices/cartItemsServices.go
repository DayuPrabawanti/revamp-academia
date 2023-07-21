package salesServices

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/salesRepositories"
	"github.com/gin-gonic/gin"
)

type CartItemsService struct {
	cartItemsRepository *salesRepositories.CartItemsRepository
}

func NewCartItemsRepository(cartItemsRepository *salesRepositories.CartItemsRepository) *CartItemsService {
	return &CartItemsService{
		cartItemsRepository: cartItemsRepository,
	}
}

func (cs CartItemsService) Getcart_items(ctx *gin.Context, id int64) (*models.SalesCartItem, *models.ResponseError) {
	return cs.cartItemsRepository.Getcart_items(ctx, int32(id))
}

func (cs CartItemsService) GetListCartItems(ctx *gin.Context) ([]*models.SalesCartItem, *models.ResponseError) {
	return cs.cartItemsRepository.GetListCartItems(ctx)
}
