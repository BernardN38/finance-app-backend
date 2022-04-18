package routes

import (
	"github.com/bernardn38/financefirst/controllers"
	"github.com/gin-gonic/gin"
)

func AuthGroup(r *gin.Engine) {
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.LoginUser)
	}
}
