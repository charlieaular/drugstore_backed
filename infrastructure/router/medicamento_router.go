package router

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/data/repositories"
	"github.com/charlieaular/drugstore_backend/domain/usecases"
	"github.com/charlieaular/drugstore_backend/infrastructure/handlers"
	"github.com/gin-gonic/gin"
)

var medicamentoHandler = getMedicamentoHandler()

func RegisterMedicamentosRoutes(router *gin.Engine) {
	medicamentosEndpoints := router.Group("/medicamento")
	medicamentosEndpoints.GET("/", medicamentoHandler.GetMedicamentos)
}

func getMedicamentoHandler() handlers.MedicamentoHandler {
	medicamentoSource := datasource.NewMedicamentoDataSourceImpl()
	medicamentoRepo := repositories.NewMedicamentoRepositoryImpl(medicamentoSource)
	medicamentoUsecase := usecases.NewMedicamentoUseCase(medicamentoRepo)
	medicamentoHandler := handlers.NewMedicamentoHandler(medicamentoUsecase)
	return *medicamentoHandler
}
