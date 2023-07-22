package salesServices

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/salesRepositories"
	"codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type SpecialOfferService struct {
	specialOfferRepository *salesRepositories.SpecialOfferRepository
}

func NewSpecialOfferService(specialOfferRepository *salesRepositories.SpecialOfferRepository) *SpecialOfferService {
	return &SpecialOfferService{
		specialOfferRepository: specialOfferRepository,
	}
}

func (cs SpecialOfferService) GetListSpecialOffer(ctx *gin.Context) ([]*models.SalesSpecialOffer, *models.ResponseError) {
	return cs.specialOfferRepository.GetListSpecialOffer(ctx)
}

func (cs SpecialOfferService) GetSpecialOffer(ctx *gin.Context, id int64) (*models.SalesSpecialOffer, *models.ResponseError) {
	return cs.specialOfferRepository.GetSpecialOffer(ctx, id)
}

func (cs SpecialOfferService) CreateSales_special_offer(ctx *gin.Context, specialOfferParams *dbContext.CreateSales_special_offerParams) (*models.SalesSpecialOffer, *models.ResponseError) {
	responseErr := validateSpecialOffer(specialOfferParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.specialOfferRepository.CreateSpecialOffer(ctx, specialOfferParams)
}

func validateSpecialOffer(specialOfferParams *dbContext.CreateSales_special_offerParams) *models.ResponseError {
	if specialOfferParams.SpofID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if specialOfferParams.SpofDescription == "" {
		return &models.ResponseError{
			Message: "Invalid special offer name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
