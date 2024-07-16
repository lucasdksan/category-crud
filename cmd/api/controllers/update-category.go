package controllers

import (
	"category/internal/repositories"
	use_cases "category/internal/use-cases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type updateCategoryInput struct {
	Id   string `uri:"id" binding:"required"`
	Name string `json:"name"`
}

func UpdateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var data updateCategoryInput

	if err := ctx.ShouldBindUri(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	idUint, err := strconv.ParseUint(data.Id, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID format"})
		return
	}

	id := uint(idUint)

	useCase := use_cases.NewUpdateCategoryUseCase(repository)

	if useCaseErr := useCase.Execute(id, data.Name); useCaseErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": useCaseErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
