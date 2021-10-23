package main

import (
	"fmt"
	"log"

	"github.com/charlieaular/drugstore_backend/config"
	"github.com/charlieaular/drugstore_backend/infrastructure/router"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := config.NewConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBPassword, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicln("Failed to Initialized postgres DB:", err)
	}

	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Worlddd")
	})

	router.RegisterFacturasRoutes(app)
	router.RegisterMedicamentosRoutes(app, db)
	router.RegisterPromocionsRoutes(app, db)

	app.Run(":8000")
}
