package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nickolss/financial_api/handler"
)

func InitRouter() {
	router := gin.Default()
	// Auth routes
	router.POST("/register", handler.RegisterClient)
	router.POST("/login", handler.Login)
	router.GET("/getUser", handler.GetClient)

	// Client routes
	router.GET("/client/:id", handler.FindClientById)

	// Operation routes
	router.POST("/financialEvent", handler.CreateFinancialEvent)
	router.GET("/financialEvent", handler.GetAllFinancialEvents)
	router.GET("/financialEvent/:id", handler.GetEventById)
	router.PUT("/financialEvent/:id", handler.UpdateFinancialEvent)
	router.DELETE("/financialEvent/:id", handler.DeleteEventById)

	// Category routes
	router.GET("/category", handler.GetAllCategories)
	router.POST("/category", handler.AddCategory)
	router.DELETE("/category/:id", handler.DeleteCategoryById)

	router.Run(":9999")
}
