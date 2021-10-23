package handlers

import (
	"github.com/charlieaular/drugstore_backend/domain/usecases"
	models "github.com/charlieaular/drugstore_backend/models"
	"github.com/gin-gonic/gin"
)

type PromocionHandler struct {
	PromocionUseCase usecases.PromocionUseCase
}

func NewPromocionHandler(PromocionUseCase usecases.PromocionUseCase) *PromocionHandler {
	return &PromocionHandler{PromocionUseCase}
}

func (ctrl *PromocionHandler) GetPromocions(c *gin.Context) {
	promociones, error := ctrl.PromocionUseCase.GetAll()
	if error != nil {
		c.JSON(400, gin.H{
			"exito": false,
			"error": error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"exito":       true,
			"promociones": promociones,
		})

	}
}

func (ctrl *PromocionHandler) Create(c *gin.Context) {
	var model models.Promocion
	bindError := c.ShouldBindJSON(&model)
	if bindError != nil {
		c.JSON(400, gin.H{
			"exito": false,
			"error": bindError.Error(),
		})
		return
	}

	_, err := ctrl.PromocionUseCase.Create(model)
	if err != nil {
		c.JSON(400, gin.H{
			"exito": false,
			"error": "promocion no creada",
		})
	} else {
		c.JSON(200, gin.H{
			"exito": true,
			"data":  "promocion creado",
		})
	}
}
