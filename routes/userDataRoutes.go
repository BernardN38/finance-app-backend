package routes

import (
	"github.com/bernardn38/financefirst/controllers"
	"github.com/bernardn38/financefirst/middleware"
	"github.com/bernardn38/financefirst/token"
	"github.com/gin-gonic/gin"
)

func UserDataGroup(r *gin.Engine) {
	tokenMaker, _ := token.NewPasetoMaker("lodrtnufepalofeiwnctvuqmifeasklo")
	userData := r.Group("/api/v1", middleware.AuthMiddleware(tokenMaker))
	{
		userData.GET("/user/sneekpeek", controllers.GetUserSneakPeek)
	}
}
