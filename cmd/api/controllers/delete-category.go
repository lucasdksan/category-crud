package controllers

import (
	"category/internal/repositories"
	use_cases "category/internal/use-cases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type deleteCategoryInput struct {
	Id string `uri:"id" binding:"required"`
}

func DeleteCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var params deleteCategoryInput

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	index, indexError := strconv.ParseUint(params.Id, 10, 32)

	if indexError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid ID format"})
		return
	}

	useCase := use_cases.NewDeleteCategoryUseCase(repository)

	categories, useCaseErr := useCase.Execute(uint(index))

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
