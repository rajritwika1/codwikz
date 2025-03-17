package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rajritwika1/codwikz/controllers"
)

// SetupUserRoutes sets up user-related routes
func SetupUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/:id", controllers.GetUserProfile)    // Example user profile fetch
		userRoutes.PUT("/:id", controllers.UpdateUserProfile) // Example user profile update
	}
}
