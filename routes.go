package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rajritwika1/codwikz/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Auth Routes
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
	}

	// User Routes
	user := router.Group("/users")
	{
		user.GET("/:id", controllers.GetUserProfile)
		user.PUT("/:id", controllers.UpdateUserProfile)
	}

	// Problem Routes
	problems := router.Group("/problems")
	{
		problems.POST("/", controllers.CreateProblem)
		problems.GET("/", controllers.GetProblems)
		problems.GET("/:id", controllers.GetProblemByID)
		problems.PUT("/:id", controllers.UpdateProblem)
		problems.DELETE("/:id", controllers.DeleteProblem)
	}
}
