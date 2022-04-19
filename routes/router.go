package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://test-finapp.herokuapp.com"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	ResourceGroup(router)
	UserDataGroup(router)
	PrivateGroup(router)
	AuthGroup(router)
	return router
}
