package router

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/data/repositories"
	"github.com/charlieaular/drugstore_backend/domain/usecases"
	"github.com/charlieaular/drugstore_backend/infrastructure/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FacturaHandler struct {
	db *gorm.DB
}

func RegisterFacturasRoutes(router *gin.Engine, db *gorm.DB) {
	FacturaHandler := NewFacturaHandler(db)
	FacturasEndpoints := router.Group("/factura")
	{
		FacturasEndpoints.GET("/", FacturaHandler.GetFacturas)
		FacturasEndpoints.POST("/", FacturaHandler.Create)
		FacturasEndpoints.GET("/grafica", FacturaHandler.GetGrafica)
	}
}

func NewFacturaHandler(db *gorm.DB) handlers.FacturaHandler {
	FacturaHandl := FacturaHandler{db: db}
	FacturaHandler := FacturaHandl.getFacturaHandler()
	return FacturaHandler
}

func (handler *FacturaHandler) getFacturaHandler() handlers.FacturaHandler {
	FacturaSource := datasource.NewFacturaDataSourceImpl(handler.db)
	FacturaRepo := repositories.NewFacturaRepositoryImpl(FacturaSource)
	FacturaUsecase := usecases.NewFacturaUseCase(FacturaRepo)
	FacturaHandler := handlers.NewFacturaHandler(FacturaUsecase)
	return *FacturaHandler
}
