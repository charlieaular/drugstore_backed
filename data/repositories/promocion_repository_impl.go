package repositories

import (
	"github.com/charlieaular/drugstore_backend/data/datasource"
	"github.com/charlieaular/drugstore_backend/domain/repositories"
	models "github.com/charlieaular/drugstore_backend/models"
)

type PromocionRepositoryImpl struct {
	PromocionDataSource datasource.PromocionDataSource
}

func NewPromocionRepositoryImpl(PromocionDataSource datasource.PromocionDataSource) repositories.PromocionRepository {
	return &PromocionRepositoryImpl{PromocionDataSource}
}

func (impl *PromocionRepositoryImpl) GetAll() ([]models.Promocion, error) {
	return impl.PromocionDataSource.GetAll()
}

func (impl *PromocionRepositoryImpl) Create(newModel models.Promocion) (models.Promocion, error) {
	return impl.PromocionDataSource.Create(newModel)
}
