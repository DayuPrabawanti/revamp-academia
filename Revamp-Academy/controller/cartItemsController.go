package controller

import (
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type CartItemsController struct {
	cartItemService *services.CartItemsService
}

func NewCartItemsController(cartItemService *services.CartItemsService) *CartItemsController {
	return &CartItemsController{
		cartItemService: cartItemService,
	}
}

func (cartItems CartItemsController) Getcart_items(ctx *gin.Context) {

	cait_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := cartItems.cartItemService.Getcart_items(ctx, int64(cait_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
