package middlewares

import "github.com/gin-gonic/gin"

var users = gin.Accounts{
	"admin": "password123",
}

func BasicAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(users)
}
