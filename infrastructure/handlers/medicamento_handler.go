package handlers

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"exito": false,
			"error": error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"exito":        true,
			"medicamentos": medicamentos,
		})

	}
}

func (ctrl *MedicamentoHandler) Create(c *gin.Context) {
	var model models.Medicamento
	c.BindJSON(&model)
	_, error := ctrl.MedicamentoUseCase.Create(model)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"exito": false,
			"error": "medicamento no creado",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"exito": true,
			"data":  "medicamento creado",
		})
	}
	return
}
