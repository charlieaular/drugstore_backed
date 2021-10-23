package handlers

import (
	"github.com/charlieaular/drugstore_backend/domain/usecases"
	"github.com/gin-gonic/gin"
)

type FacturaHandler struct {
	FacturaUseCase usecases.FacturaUseCase
}

func NewFacturaHandler(FacturaUseCase usecases.FacturaUseCase) *FacturaHandler {
	return &FacturaHandler{FacturaUseCase}
}

func (ctrl *FacturaHandler) GetFacturas(c *gin.Context) {
	message := ctrl.FacturaUseCase.GetAll()
	c.JSON(200, gin.H{
		"facturas": message,
	})

}
