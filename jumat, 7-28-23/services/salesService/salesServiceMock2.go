package salesService

import (
	"codeid.revampacademy/models"
	"github.com/gin-gonic/gin"
)

func (sm ServiceMock) GetMockupId(ctx *gin.Context, id int64) (*models.CurriculumProgramEntity, *models.ResponseError) {
	return sm.repositoryMock.GetMockupId(ctx, id)
}

func (sm ServiceMock) GetListProgramMock(ctx *gin.Context, orderby string) ([]*models.CurriculumProgramEntity, *models.ResponseError) {
	return sm.repositoryMock.GetListProgramMock(ctx, orderby)
}
