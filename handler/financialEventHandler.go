package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickolss/financial_api/config"
	"github.com/nickolss/financial_api/schemas"
)

func CreateFinancialEvent(ctx *gin.Context) {
	var request schemas.FinancialEventRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user bad request",
			"error":   err.Error(),
		})
		return
	}

	var client schemas.Client
	if err := config.DB.Where("id = ?", request.ClientId).First(&client).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
			"error":   err.Error(),
		})
		return
	}

	event := schemas.FinancialEvent{
		FinancialName: request.FinancialName,
		Amount:        request.Amount,
		ClientId:      request.ClientId,
	}

	if err := config.DB.Create(&event).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "create financial operation is working",
		"data":    event,
	})
}

func GetAllFinancialEvents(ctx *gin.Context) {
	var events []schemas.FinancialEvent
	if err := config.DB.Preload("Client").Find(&events).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get all financial events",
		"data":    events,
	})
}

func GetEventById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user bad request",
			"error":   "invalid id",
		})
		return
	}

	var event schemas.FinancialEvent

	result := config.DB.Find(&event, id)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get operation by id",
		"data":    event,
	})
}

func UpdateFinancialEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user bad request",
			"error":   "invalid id",
		})
		return
	}

	var event schemas.FinancialEvent

	result := config.DB.Find(&event, id)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
		return
	}

	var request schemas.FinancialEventRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user bad request",
			"error":   err.Error(),
		})
		return
	}

	event.FinancialName = request.FinancialName
	event.Amount = request.Amount
	event.ClientId = request.ClientId
	if err := config.DB.Save(&event).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "event updated",
		"data":    event,
	})
}

func DeleteEventById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user bad request",
			"error":   "invalid id",
		})
		return
	}

	if err := config.DB.Delete(&schemas.FinancialEvent{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "event deleted",
	})
}
