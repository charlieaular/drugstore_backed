package handlers

import (
	"github.com/charlieaular/drugstore_backend/domain/usecases"
	"github.com/gin-gonic/gin"
)

type MedicamentoHandler struct {
	MedicamentoUseCase usecases.MedicamentoUseCase
}

func NewMedicamentoHandler(MedicamentoUseCase usecases.MedicamentoUseCase) *MedicamentoHandler {
	return &MedicamentoHandler{MedicamentoUseCase}
}

func (ctrl *MedicamentoHandler) GetMedicamentos(c *gin.Context) {
	medicamentos, error := ctrl.MedicamentoUseCase.GetAll()
	if error != nil {
		c.JSON(400, gin.H{
			"error": error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"medicamentos": medicamentos,
		})

	}

}
