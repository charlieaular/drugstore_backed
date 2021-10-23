package repositories

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/domain/repositories"
)

type FacturaRepositoryImpl struct {
	FacturaDataSource datasource.FacturaDataSource
}

func NewFacturaRepositoryImpl(FacturaDataSource datasource.FacturaDataSource) repositories.FacturaRepository {
	return &FacturaRepositoryImpl{FacturaDataSource}
}

func (impl *FacturaRepositoryImpl) GetAll() string {
	return impl.FacturaDataSource.GetAll()
}
