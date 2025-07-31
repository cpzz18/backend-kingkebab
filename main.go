package main

import (
	"github.com/gin-gonic/gin"

	"backend-kking/backend/config"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	
	r.Run(":8080")
}
