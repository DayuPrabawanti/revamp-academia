package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type CartItemsService struct {
	cartItemsRepository *repositories.CartItemsRepository
}

func NewCartItemsRepository(cartItemsRepository *repositories.CartItemsRepository) *CartItemsService {
	return &CartItemsService{
		cartItemsRepository: cartItemsRepository,
	}
}

func (cs CartItemsService) Getcart_items(ctx *gin.Context, id int64) (*models.SalesCartItem, *models.ResponseError) {
	return cs.cartItemsRepository.Getcart_items(ctx, int32(id))
}
