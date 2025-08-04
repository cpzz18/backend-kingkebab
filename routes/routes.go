package routes

import (
	"github.com/gin-gonic/gin"

	"backend-kking/backend/controllers"
	"backend-kking/backend/middleware"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	//auth
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	api.GET("/products", controllers.GetProducts)
	api.GET("/products/:id", controllers.GetProductByID)

	//r.admin
	api.POST("/admin/login", controllers.AdminLogin)
	admin := api.Group("/admin")
	admin.Use(middleware.JWTMiddleware, middleware.AdminOnly())
	{
		admin.GET("/products", controllers.GetProducts)
		admin.POST("/products", controllers.CreateProducts)
		admin.PUT("/products/:id", controllers.UpdateProduct)
		admin.DELETE("/products/:id", controllers.DeleteProduct)
	}

}
