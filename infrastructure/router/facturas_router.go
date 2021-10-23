package router

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/data/repositories"
	"github.com/charlieaular/drugstore_backend/domain/usecases"
	"github.com/charlieaular/drugstore_backend/infrastructure/handlers"
	"github.com/gin-gonic/gin"
)

var facturaHandler = getFacturaHandler()

func RegisterFacturasRoutes(router *gin.Engine) {
	facturasEndpoints := router.Group("/facturas")
	facturasEndpoints.GET("/get-all", facturaHandler.GetFacturas)
}

func getFacturaHandler() handlers.FacturaHandler {
	facturaSource := datasource.NewFacturaDataSourceImpl()
	facturaRepo := repositories.NewFacturaRepositoryImpl(facturaSource)
	facturaUsecase := usecases.NewFacturaUseCase(facturaRepo)
	facturaHandler := handlers.NewFacturaHandler(facturaUsecase)
	return *facturaHandler
}
