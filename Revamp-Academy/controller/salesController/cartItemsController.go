package salesController

import (
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/services/salesServices"
	"github.com/gin-gonic/gin"
)

type CartItemsController struct {
	cartItemService *salesServices.CartItemsService
}

func NewCartItemsController(cartItemService *salesServices.CartItemsService) *CartItemsController {
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

func (cartItemsControl CartItemsController) GetListCartItems(ctx *gin.Context) {
	response, resposeErr := cartItemsControl.cartItemService.GetListCartItems(ctx)

	if resposeErr != nil {
		ctx.JSON(resposeErr.Status, resposeErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	// ctx.JSON(http.StatusOK, "Hello gin Framework")
}
