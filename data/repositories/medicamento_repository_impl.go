package repositories

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/domain/repositories"
	models "github.com/charlieaular/drugstore_backend/models"
)

type MedicamentoRepositoryImpl struct {
	MedicamentoDataSource datasource.MedicamentoDataSource
}

func NewMedicamentoRepositoryImpl(MedicamentoDataSource datasource.MedicamentoDataSource) repositories.MedicamentoRepository {
	return &MedicamentoRepositoryImpl{MedicamentoDataSource}
}

func (impl *MedicamentoRepositoryImpl) GetAll() ([]models.Medicamento, error) {
	return impl.MedicamentoDataSource.GetAll()
}

func (impl *MedicamentoRepositoryImpl) Create(newModel models.Medicamento) (models.Medicamento, error) {
	return impl.MedicamentoDataSource.Create(newModel)
}
