package controllers

import (
	"category/internal/repositories"
	use_cases "category/internal/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})

		return
	}

	useCase := use_cases.NewCreateCategoryUseCase(repository)

	if useCaseErr := useCase.Execute(body.Name); useCaseErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   useCaseErr.Error(),
			})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
