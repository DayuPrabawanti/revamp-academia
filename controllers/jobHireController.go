package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/dbContext"
	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type JobCategoryController struct {
	JobCategoryService *services.JobCategoryService
}

type JobClientController struct {
    JobClientService  *services.JobClientService
}

type JobPostController struct {
    JobPostService  *services.JobPostService
}

// --------------------------- JOB CATEGORY --------------------------- 

// DECLARE CONSTRUCTOR

func NewJobCategoryController(jobhireController *services.JobCategoryService) *JobCategoryController {
    return &JobCategoryController{
        JobCategoryService: jobhireController,
    }
}

// METHODE JOB CATEGORY 

func (jobCategoryController JobCategoryController) GetJobCategoryHttp(ctx *gin.Context) {

		jobCategoryId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			log.Println("Error while reading paramater id", err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		response, responseErr := jobCategoryController.JobCategoryService.GetJobCategoryService(ctx, int64(jobCategoryId))
		if responseErr != nil {

			ctx.JSON(responseErr.Status, responseErr)
			return
		}

		ctx.JSON(http.StatusOK, response)
}

	func (jobCategoryController JobCategoryController) ListJobCategoryHttp(ctx *gin.Context) {
			response, responseErr := jobCategoryController.JobCategoryService.ListJobCategoryService(ctx)

			if responseErr != nil {
				ctx.JSON(responseErr.Status, responseErr)
				return
			}

			ctx.JSON(http.StatusOK, response)
	}

		func (jobCategoryController JobCategoryController) CreateJobCategoryHttp(ctx *gin.Context) {

				body, err := io.ReadAll(ctx.Request.Body)
				if err != nil {
					log.Println("Error while reading create category request body", err)
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}

				var category dbContext.CreateJobCategoryParams
				err = json.Unmarshal(body, &category)
				if err != nil {
					log.Println("Error while unmarshaling create category request body", err)
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}

				response, responseErr := jobCategoryController.JobCategoryService.CreateJobCategoryService(ctx, &category)
				if responseErr != nil {
					ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
					return
				}

				ctx.JSON(http.StatusOK, response)
		}

			func (jobCategoryController JobCategoryController) UpdateJobCategoryHttp(ctx *gin.Context) {

					categoryId, err := strconv.Atoi(ctx.Param("id"))

					if err != nil {
						log.Println("Error while reading paramater id", err)
						ctx.AbortWithError(http.StatusBadRequest, err)
						return
					}

					body, err := io.ReadAll(ctx.Request.Body)
					if err != nil {
						log.Println("Error while reading update category request body", err)
						ctx.AbortWithError(http.StatusInternalServerError, err)
						return
					}

					var category dbContext.CreateJobCategoryParams
					err = json.Unmarshal(body, &category)
					if err != nil {
						log.Println("Error while unmarshaling update category request body", err)
						ctx.AbortWithError(http.StatusInternalServerError, err)
						return
					}

					response := jobCategoryController.JobCategoryService.UpdateCategoryService(ctx, &category, int64(categoryId))
					if response != nil {
						ctx.AbortWithStatusJSON(response.Status, response)
						return
					}

					ctx.JSON(http.StatusOK, response)
			}

				func (jobCategoryController JobCategoryController) DeleteJobCategoryHttp(ctx *gin.Context) {

						categoryId, err := strconv.Atoi(ctx.Param("id"))

						if err != nil {
							log.Println("Error while reading paramater id", err)
							ctx.AbortWithError(http.StatusBadRequest, err)
							return
						}

						responseErr := jobCategoryController.JobCategoryService.DeleteCategoryService(ctx, int64(categoryId))
						if responseErr != nil {
							ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
							return
						}

						ctx.Status(http.StatusNoContent)
				}

// --------------------------- JOB CLIENT --------------------------- 


func NewJobClientController(jobhireClientController *services.JobClientService) *JobClientController {
		return &JobClientController{
			JobClientService: jobhireClientController,
		}
}

	func (clientController JobClientController) GetJobClientHttp(ctx *gin.Context) {

			categoryId, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				log.Println("Error while reading paramater id", err)
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			response, responseErr := clientController.JobClientService.GetJobClientService(ctx, int64(categoryId))
			if responseErr != nil {

				ctx.JSON(responseErr.Status, responseErr)
				return
			}

			ctx.JSON(http.StatusOK, response)
	}

		func (clientController JobClientController) ListJobClientHttp(ctx *gin.Context) {
			response, responseErr := clientController.JobClientService.ListJobClientService(ctx)

			if responseErr != nil {
				ctx.JSON(responseErr.Status, responseErr)
				return
			}

			ctx.JSON(http.StatusOK, response)
		}	

			func (clientController JobClientController) CreateJobClientHttp(ctx *gin.Context) {

				body, err := io.ReadAll(ctx.Request.Body)
				if err != nil {
					log.Println("Error while reading create category request body", err)
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}

				var category dbContext.CreateJobClientParams
				err = json.Unmarshal(body, &category)
				if err != nil {
					log.Println("Error while unmarshaling create category request body", err)
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}

				response, responseErr := clientController.JobClientService.JobClientServiceRepo.CreateJobClientRepo(ctx, &category)
				if responseErr != nil {
					ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
					return
				}

				ctx.JSON(http.StatusOK, response)
			}

				func (clientController JobClientController) UpdateJobClientHttp(ctx *gin.Context) {

					categoryId, err := strconv.Atoi(ctx.Param("id"))

					if err != nil {
						log.Println("Error while reading paramater id", err)
						ctx.AbortWithError(http.StatusBadRequest, err)
						return
					}

					body, err := io.ReadAll(ctx.Request.Body)
					if err != nil {
						log.Println("Error while reading update category request body", err)
						ctx.AbortWithError(http.StatusInternalServerError, err)
						return
					}

					var category dbContext.CreateJobClientParams
					err = json.Unmarshal(body, &category)
					if err != nil {
						log.Println("Error while unmarshaling update category request body", err)
						ctx.AbortWithError(http.StatusInternalServerError, err)
						return
					}

					response := clientController.JobClientService.UpdateJobClientService(ctx, &category, int64(categoryId))
					if response != nil {
						ctx.AbortWithStatusJSON(response.Status, response)
						return
					}

					ctx.JSON(http.StatusOK, response)
				}

					func (clientController JobClientController) DeleteJobClientHttp(ctx *gin.Context) {

						categoryId, err := strconv.Atoi(ctx.Param("id"))

						if err != nil {
							log.Println("Error while reading paramater id", err)
							ctx.AbortWithError(http.StatusBadRequest, err)
							return
						}

						responseErr := clientController.JobClientService.DeleteJobClientService(ctx, int64(categoryId))
						if responseErr != nil {
							ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
							return
						}

						ctx.Status(http.StatusNoContent)
					}

// --------------------------- JOB POST --------------------------- 

func NewJobPostController(jobhirePostController *services.JobPostService) *JobPostController {
		return &JobPostController{
			JobPostService: jobhirePostController,
		}
}

	func (jobPostController JobPostController) GetJobPostHttp(ctx *gin.Context) {

				categoryId, err := strconv.Atoi(ctx.Param("id"))

				if err != nil {
					log.Println("Error while reading paramater id", err)
					ctx.AbortWithError(http.StatusBadRequest, err)
					return
				}

				response, responseErr := jobPostController.JobPostService.GetJobPostService(ctx, int32(categoryId))
				if responseErr != nil {

					ctx.JSON(responseErr.Status, responseErr)
					return
				}

				ctx.JSON(http.StatusOK, response)
	}

		func (jobPostController JobPostController) ListJobPostHttp(ctx *gin.Context) {
					response, responseErr := jobPostController.JobPostService.ListJobPostService(ctx)

					if responseErr != nil {
						ctx.JSON(responseErr.Status, responseErr)
						return
					}

					ctx.JSON(http.StatusOK, response)
		}