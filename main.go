package main

import (
	"github.com/gin-gonic/gin"

	"backend-kking/backend/config"
	"backend-kking/backend/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
