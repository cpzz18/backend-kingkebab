package routes

import (
	"github.com/gin-gonic/gin"

	"backend-kking/backend/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	//auth
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	
}
