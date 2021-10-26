package handlers

import (
	"net/http"

	"github.com/charlieaular/drugstore_backend/domain/usecases"
	models "github.com/charlieaular/drugstore_backend/models"
	"github.com/charlieaular/drugstore_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type FacturaHandler struct {
	FacturaUseCase usecases.FacturaUseCase
}

type CreateDTO struct {
	Fecha        utils.CustomDate `json:"fecha"`
	PagoTotal    decimal.Decimal  `json:"pago_total"`
	Promocion    *int             `json:"promocion,omitempty"`
	Medicamentos []int            `json:"medicamentos"`
}

func NewFacturaHandler(FacturaUseCase usecases.FacturaUseCase) *FacturaHandler {
	return &FacturaHandler{FacturaUseCase}
}

func (ctrl *FacturaHandler) GetFacturas(c *gin.Context) {
	facturas, error := ctrl.FacturaUseCase.GetAll()
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"exito": false,
			"error": error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"exito":    true,
			"facturas": facturas,
		})

	}
}

func (ctrl *FacturaHandler) Create(c *gin.Context) {
	var model CreateDTO
	bindError := c.BindJSON(&model)

	if bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"exito": false,
			"error": bindError.Error(),
		})
		return
	}
	facturaModel := models.Factura{
		FechaCrear:  model.Fecha,
		PagoTotal:   decimal.NewFromFloat(model.PagoTotal.InexactFloat64()),
		PromocionID: model.Promocion,
	}

	medicamentos := model.Medicamentos

	_, error := ctrl.FacturaUseCase.Create(facturaModel, medicamentos)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"exito": false,
			"error": "Factura no creado",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"exito": true,
			"data":  "Factura creado",
		})
	}
	return
}

func (ctrl *FacturaHandler) GetGrafica(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Copy().Query("fecha_fin")

	facturas, error := ctrl.FacturaUseCase.GetGrafica(fechaInicio, fechaFin)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"exito": false,
			"error": error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"exito":    true,
			"facturas": facturas,
		})

	}
}
