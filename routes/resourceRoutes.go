package routes

import (
	"github.com/bernardn38/financefirst/controllers"
	"github.com/bernardn38/financefirst/middleware"
	"github.com/bernardn38/financefirst/token"
	"github.com/gin-gonic/gin"
)

func ResourceGroup(r *gin.Engine) {
	tokenMaker, _ := token.NewPasetoMaker("lodrtnufepalofeiwnctvuqmifeasklo")
	resources := r.Group("/api/v1", middleware.AuthMiddleware(tokenMaker))
	{
		resources.GET("/", welcomePage)
		resources.GET("/user/:userId/transactions", controllers.GetAllTransactions)
		resources.GET("/user/:userId/transactions/:id", controllers.GetTransaction)
		resources.GET("/user/:userId/transactions/sum", controllers.GetMonthSumsByType)
		resources.GET("/user/:userId/transactions/limit", controllers.GetTransactionsLimit)
	}
}

func welcomePage(c *gin.Context) {
	c.JSON(200, `{"enpoints":["/transactions","transactions/:id","transactions/sum", transactions/limit]}`)
}
