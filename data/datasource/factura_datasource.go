package datasource

import (
	"fmt"

	models "github.com/charlieaular/drugstore_backend/models"
	"gorm.io/gorm"
)

type FacturaDataSource interface {
	GetAll() ([]models.Factura, error)
	Create(newModel models.Factura, medicamentos []int) (models.Factura, error)
}

type FacturaDataSourceImpl struct {
	db                *gorm.DB
	FacturaDataSource FacturaDataSource
}

func NewFacturaDataSourceImpl(db *gorm.DB) FacturaDataSource {
	return FacturaDataSourceImpl{db: db}
}

func (impl FacturaDataSourceImpl) GetAll() ([]models.Factura, error) {
	var facturas []models.Factura
	if result := impl.db.Preload("Medicamentos").Preload("Promocion").Find(&facturas); result.Error != nil {
		return nil, result.Error
	}

	return facturas, nil
}

func (impl FacturaDataSourceImpl) Create(newModel models.Factura, medicamentos []int) (models.Factura, error) {
	fmt.Println(newModel.PagoTotal)
	if result := impl.db.Create(&newModel); result.Error != nil {
		return models.Factura{}, result.Error
	}

	if len(medicamentos) > 0 {
		var facturaMedicamentos []models.FacturaMedicamento
		for _, v := range medicamentos {
			facturaMedicamentosModel := models.FacturaMedicamento{
				MedicamentoID: v,
				FacturaID:     int(newModel.ID),
			}
			facturaMedicamentos = append(facturaMedicamentos, facturaMedicamentosModel)
		}

		if result := impl.db.Create(&facturaMedicamentos); result.Error != nil {
			return models.Factura{}, result.Error
		}
	}

	return newModel, nil
}
