package datasource

import (
	models "github.com/charlieaular/drugstore_backend/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type FacturaDataSource interface {
	GetAll() ([]models.Factura, error)
	Create(newModel models.Factura, medicamentos []int) (models.Factura, error)
	GetGrafica(fechaInicio string, fechaFin string) ([]models.Factura, error)
	Simular(fecha string, medicamentos []string) (models.Factura, error)
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

func (impl FacturaDataSourceImpl) GetGrafica(fechaInicio string, fechaFin string) ([]models.Factura, error) {
	var facturas []models.Factura
	if result := impl.db.Where("fecha_crear BETWEEN ? AND ?", fechaInicio, fechaFin).Order("fecha_crear desc").Find(&facturas); result.Error != nil {
		return nil, result.Error
	}

	return facturas, nil
}

func (impl FacturaDataSourceImpl) Simular(fecha string, medicamentos []string) (models.Factura, error) {
	var promocion models.Promocion
	resultPromocion := impl.db.Where("fecha_inicio <= ? and fecha_fin >= ?", fecha, fecha).Find(&promocion)
	if resultPromocion.Error != nil {
		return models.Factura{}, resultPromocion.Error
	}
	var sum int
	resultFactura := impl.db.Table("medicamento").Select("sum(precio) as precio_total").Where("id in (?)", medicamentos).Row().Scan(&sum)

	if resultFactura != nil {
		return models.Factura{}, resultFactura
	}

	restar := (promocion.Porcentaje * float64(sum)) / 100
	precioTotal := float64(sum) - restar

	return models.Factura{
		PagoTotal: decimal.NewFromFloat(precioTotal),
	}, nil

}
