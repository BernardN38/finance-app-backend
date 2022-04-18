package routes

import (
	"github.com/bernardn38/financefirst/controllers"
	"github.com/bernardn38/financefirst/middleware"
	"github.com/bernardn38/financefirst/token"
	"github.com/gin-gonic/gin"
)

func PrivateGroup(r *gin.Engine) {
	tokenMaker, _ := token.NewPasetoMaker("lodrtnufepalofeiwnctvuqmifeasklo")
	private := r.Group("/api/v1/", middleware.AuthMiddleware(tokenMaker))
	{
		private.GET("headers", controllers.ReadHeaders)
		private.GET("profile", controllers.GetUserProfile)
	}
}
