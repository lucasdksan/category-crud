package controllers

import (
	"category/internal/repositories"
	use_cases "category/internal/use-cases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type readCategoryInput struct {
	Id string `uri:"id" binding:"required"`
}

func ReadCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var params readCategoryInput

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	idUint, err := strconv.ParseUint(params.Id, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID format"})
		return
	}

	id := uint(idUint)

	useCase := use_cases.NewReadCategoryUseCase(repository)
	category, useCaseErr := useCase.Execute(id)

	if useCaseErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": useCaseErr.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "categories": category})
}
