package main

import (
	"fmt"

	"github.com/nickolss/financial_api/config"
	"github.com/nickolss/financial_api/routes"
)

func main() {
	// Initialize Configs
	err := config.InitConfig()
	if err != nil {
		fmt.Println("Config initialization error:", err)
		return
	}

	// Initialize Routes
	routes.InitRouter()
}
