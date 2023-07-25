package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	userrepo "codeid.revampacademy/repositories/usersRepository"
	dbcontext "codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserSkillService struct {
	UserSkillRepository *userrepo.UserSkillRepository
}

func NewUserSkillService(UserSkillRepository *userrepo.UserSkillRepository) *UserSkillService {
	return &UserSkillService{
		UserSkillRepository: UserSkillRepository,
	}
}

// GetList User Skill
func (cs UserSkillService) GetListUserSkill(ctx *gin.Context) ([]*models.UsersUsersSkill, *models.ResponseError) {
	return cs.UserSkillRepository.GetListUsersSkill(ctx)
}

// Get User Skill
func (cs UserSkillService) GetUserSkill(ctx *gin.Context, id int32) (*models.UsersUsersSkill, *models.ResponseError) {
	return cs.UserSkillRepository.GetUserSkill(ctx, id)
}

// Create User Media
func (cs UserSkillService) CreateUserSkill(ctx *gin.Context, SkillParams *dbcontext.CreateSkillParams) (*models.UsersUsersSkill, *models.ResponseError) {
	responseErr := validateSkill(SkillParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.UserSkillRepository.CreateUserSkill(ctx, SkillParams)
}

// Update Table
func (cs UserSkillService) UpdateSkill(ctx *gin.Context, skillParams *dbcontext.CreateSkillParams, id int64) *models.ResponseError {
	responseErr := validateSkill(skillParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.UserSkillRepository.UpdateSkill(ctx, skillParams)
}

// Validate
func validateSkill(skillParams *dbcontext.CreateSkillParams) *models.ResponseError {
	if skillParams.UskiID == 0 {
		return &models.ResponseError{
			Message: "Invalid Skill ",
			Status:  http.StatusBadRequest,
		}
	}

	if skillParams.UskiEntityID == 0 {
		return &models.ResponseError{
			Message: "Required Skill ",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

// Delete Table
func (cs UserSkillService) DeleteSkill(ctx *gin.Context, id int32) *models.ResponseError {
	return cs.UserSkillRepository.DeleteSkill(ctx, id)
}
