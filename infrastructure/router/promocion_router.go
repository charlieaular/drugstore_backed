package router

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/data/repositories"
	"github.com/charlieaular/drugstore_backend/domain/usecases"
	"github.com/charlieaular/drugstore_backend/infrastructure/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PromocionHandler struct {
	db *gorm.DB
}

func RegisterPromocionsRoutes(router *gin.Engine, db *gorm.DB) {
	promocionHandler := NewPromocionHandler(db)
	promocionsEndpoints := router.Group("/promocion")
	{
		promocionsEndpoints.GET("/", promocionHandler.GetPromocions)
		promocionsEndpoints.POST("/", promocionHandler.Create)
	}
}

func NewPromocionHandler(db *gorm.DB) handlers.PromocionHandler {
	promocionHandl := PromocionHandler{db: db}
	promocionHandler := promocionHandl.getPromocionHandler()
	return promocionHandler
}

func (handler *PromocionHandler) getPromocionHandler() handlers.PromocionHandler {
	promocionSource := datasource.NewPromocionDataSourceImpl(handler.db)
	promocionRepo := repositories.NewPromocionRepositoryImpl(promocionSource)
	promocionUsecase := usecases.NewPromocionUseCase(promocionRepo)
	promocionHandler := handlers.NewPromocionHandler(promocionUsecase)
	return *promocionHandler
}
