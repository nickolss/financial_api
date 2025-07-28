package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickolss/financial_api/config"
	"github.com/nickolss/financial_api/schemas"
)

func GetAllCategories(ctx *gin.Context) {
	var categories []schemas.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   result.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get all categories",
		"data":    categories,
	})
}

func AddCategory(ctx *gin.Context) {
	var category schemas.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user bad request",
			"error":   err.Error(),
		})
		return
	}

	result := config.DB.Create(&category)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   result.Error,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "category added",
		"data":    category,
	})
}

func DeleteCategoryById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user bad request",
			"data":    nil,
		})
	}

	result := config.DB.Delete(&schemas.Category{}, id)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   result.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted category by id",
	})
}
