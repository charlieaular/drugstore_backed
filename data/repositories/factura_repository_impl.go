package repositories

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/domain/repositories"
	models "github.com/charlieaular/drugstore_backend/models"
)

type FacturaRepositoryImpl struct {
	FacturaDataSource datasource.FacturaDataSource
}

func NewFacturaRepositoryImpl(FacturaDataSource datasource.FacturaDataSource) repositories.FacturaRepository {
	return &FacturaRepositoryImpl{FacturaDataSource}
}

func (impl *FacturaRepositoryImpl) GetAll() ([]models.Factura, error) {
	return impl.FacturaDataSource.GetAll()
}

func (impl *FacturaRepositoryImpl) Create(newModel models.Factura, medicamentos []int) (models.Factura, error) {
	return impl.FacturaDataSource.Create(newModel, medicamentos)
}

func (impl *FacturaRepositoryImpl) GetGrafica(fechaInicio string, fechaFin string) ([]models.Factura, error) {
	return impl.FacturaDataSource.GetGrafica(fechaInicio, fechaFin)
}

func (impl *FacturaRepositoryImpl) Simular(fecha string, medicamentos []string) (models.Factura, error) {
	return impl.FacturaDataSource.Simular(fecha, medicamentos)
}
