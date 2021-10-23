package repositories

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/domain/repositories"
)

type MedicamentoRepositoryImpl struct {
	MedicamentoDataSource datasource.MedicamentoDataSource
}

func NewMedicamentoRepositoryImpl(MedicamentoDataSource datasource.MedicamentoDataSource) repositories.MedicamentoRepository {
	return &MedicamentoRepositoryImpl{MedicamentoDataSource}
}

func (impl *MedicamentoRepositoryImpl) GetAll() string {
	return impl.MedicamentoDataSource.GetAll()
}
