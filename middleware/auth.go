package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"backend-kking/backend/config"
	"backend-kking/backend/models"
)

func JWTMiddleware(c *gin.Context) {
	tokenStr, err := c.Cookie("Authorization")
	if err != nil || tokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metode signing tidak sesuai: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		c.Abort()
		return
	}

	if exp, ok := claims["exp"].(float64); !ok || float64(time.Now().Unix()) > exp {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token sudah kadaluarsa"})
		c.Abort()
		return
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak mengandung user_id"})
		c.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak mengandung role"})
		c.Abort()
		return
	}

	var user models.User
	if err := config.DB.First(&user, uint(userID)).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		c.Abort()
		return
	}

	c.Set("userID", uint(userID))
	c.Set("userRole", role)
	c.Set("user", user)

	c.Next()
}
