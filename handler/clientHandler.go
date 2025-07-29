package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickolss/financial_api/config"
	"github.com/nickolss/financial_api/schemas"
)

func FindClientById(ctx *gin.Context) {
	var client schemas.Client
	id := ctx.Param("id")

	if err := config.DB.Find(&client, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "client not found",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "client found",
		"data":    client,
	})
}
