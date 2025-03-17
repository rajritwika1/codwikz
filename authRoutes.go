package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rajritwika1/codwikz/controllers"
)

// SetupAuthRoutes sets up authentication routes
func SetupAuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", controllers.RegisterUser)
		authRoutes.POST("/login", controllers.LoginUser)
	}
}
