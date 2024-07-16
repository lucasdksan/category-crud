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

	categoryRoutes.DELETE("/:id", func(ctx *gin.Context) {
		controllers.DeleteCategory(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.PUT("/:id", func(ctx *gin.Context) {
		controllers.UpdateCategory(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.GET("/:id", func(ctx *gin.Context) {
		controllers.ReadCategory(ctx, inMemoryCategoryRepository)
	})
}
