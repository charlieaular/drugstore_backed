package handlers

import (
	"github.com/charlieaular/drugstore_backend/domain/usecases"
	models "github.com/charlieaular/drugstore_backend/models"
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
			"exito": false,
			"error": error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"exito":        true,
			"medicamentos": medicamentos,
		})

	}
}

func (ctrl *MedicamentoHandler) Create(c *gin.Context) {
	var model models.Medicamento
	c.BindJSON(&model)
	medicamento, error := ctrl.MedicamentoUseCase.Create(model)
	if error != nil {
		c.JSON(400, gin.H{
			"exito": false,
			"error": error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"exito":       true,
			"medicamento": medicamento,
		})

	}
}
