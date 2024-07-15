package main

import (
	"category/cmd/api/controllers"
	"category/internal/repositories"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")
	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()
	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})
	categoryRoutes.GET("/list", func(ctx *gin.Context) {
		controllers.ListCategory(ctx, inMemoryCategoryRepository)
	})
}
