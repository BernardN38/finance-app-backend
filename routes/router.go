package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://192.168.0.11:3000", "*"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	ResourceGroup(router)
	UserDataGroup(router)
	PrivateGroup(router)
	AuthGroup(router)
	return router
}
