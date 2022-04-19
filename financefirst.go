package main

import (
	"os"

	"github.com/bernardn38/financefirst/db"
	"github.com/bernardn38/financefirst/routes"
	"github.com/gin-gonic/contrib/static"
)

func main() {
	r := routes.SetupRouter()
	r.Use(static.Serve("/", static.LocalFile("./build", true)))
	db.InitDb()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	r.Run(port)
}
