package controllers

import (
	"category/internal/repositories"
	use_cases "category/internal/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCase := use_cases.NewListCategoryUseCase(repository)
	categories, useCaseErr := useCase.Execute()

	if useCaseErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   useCaseErr.Error(),
			})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "categories": categories})
}
