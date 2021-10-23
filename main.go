package main

import (
	"github.com/charlieaular/drugstore_backend/infrastructure/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Worlddd")
	})

	router.RegisterFacturasRoutes(app)

	app.Run(":8000")
}
