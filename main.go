package main

import (
	"fmt"
	"log"
	"time"

	"github.com/charlieaular/drugstore_backend/config"
	"github.com/charlieaular/drugstore_backend/infrastructure/router"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
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

	app.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Minute,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	v1 := app.Group("/v1")

	router.RegisterFacturasRoutes(v1, db)
	router.RegisterMedicamentosRoutes(v1, db)
	router.RegisterPromocionsRoutes(v1, db)

	app.Run(":8000")
}
