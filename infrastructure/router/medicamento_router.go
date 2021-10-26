package router

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/data/repositories"
	"github.com/charlieaular/drugstore_backend/domain/usecases"
	"github.com/charlieaular/drugstore_backend/infrastructure/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MedicamentoHandler struct {
	db *gorm.DB
}

func RegisterMedicamentosRoutes(router *gin.RouterGroup, db *gorm.DB) {
	medicamentoHandler := NewMedicamentoHandler(db)
	medicamentosEndpoints := router.Group("/medicamento")
	{
		medicamentosEndpoints.GET("/", medicamentoHandler.GetMedicamentos)
		medicamentosEndpoints.POST("/", medicamentoHandler.Create)
	}
}

func NewMedicamentoHandler(db *gorm.DB) handlers.MedicamentoHandler {
	medicamentoHandl := MedicamentoHandler{db: db}
	medicamentoHandler := medicamentoHandl.getMedicamentoHandler()
	return medicamentoHandler
}

func (handler *MedicamentoHandler) getMedicamentoHandler() handlers.MedicamentoHandler {
	medicamentoSource := datasource.NewMedicamentoDataSourceImpl(handler.db)
	medicamentoRepo := repositories.NewMedicamentoRepositoryImpl(medicamentoSource)
	medicamentoUsecase := usecases.NewMedicamentoUseCase(medicamentoRepo)
	medicamentoHandler := handlers.NewMedicamentoHandler(medicamentoUsecase)
	return *medicamentoHandler
}
