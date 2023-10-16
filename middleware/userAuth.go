package middleware

import (
	"firstpro/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("'👌")
		// Retrieve the JWT token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		tokenString := helper.GetTokenFromHeader(authHeader)
		fmt.Println(tokenString, "🎶")

		// Validate the token and extract the user ID
		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				fmt.Printf("💖")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		userID, userEmail, err := helper.ExtractUserIDFromToken(tokenString)
		fmt.Println(userID, "✔")
		if err != nil {
			fmt.Println(err, "✌")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Add the user ID to the Gin context
		fmt.Println(userID, "👌")
		c.Set("user_id", userID)
		c.Set("user_email", userEmail)

		// Call the next handler
		c.Next()
	}
}
