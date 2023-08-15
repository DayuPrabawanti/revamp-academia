package bootcampController

import (
	"net/http"

	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type GabungController struct {
	gabungservice *bootcampService.Gabung
}

// Declare constructor
func NewGabungController(gabungService *bootcampService.Gabung) *GabungController {
	return &GabungController{
		gabungservice: gabungService,
	}
}

func (GabungController GabungController) GetListGabung(ctx *gin.Context) {

	response, responseErr := GabungController.gabungservice.GetListGabung(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}
