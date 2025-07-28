package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickolss/financial_api/config"
	"github.com/nickolss/financial_api/schemas"
)

// @Summary Register a new user
// @Description API to create a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body schemas.Client true "Request Body"
// @Router /register [post]
func RegisterUser(ctx *gin.Context) {
	var client schemas.Client

	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user bad request",
			"error":    err.Error(),
		})
		return
	}

	result := config.DB.Create(&client)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":    result.Error,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"data":    client,
	})
}

func Login(ctx *gin.Context) {
	var loginRequest schemas.ClientLoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user bad request",
			"error":    err.Error(),
		})
		return
	}

	var client schemas.Client
	if err := config.DB.Where("email = ? AND password = ?", loginRequest.Email, loginRequest.Password).First(&client).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
			"data":    nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "login",
		"data": client,
	})
}

func GetUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get user is working",
	})
}
