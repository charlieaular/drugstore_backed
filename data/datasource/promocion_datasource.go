package datasource

import (
	models "github.com/charlieaular/drugstore_backend/models"
	"gorm.io/gorm"
)

type PromocionDataSource interface {
	GetAll() ([]models.Promocion, error)
	Create(newModel models.Promocion) (models.Promocion, error)
}

type PromocionDataSourceImpl struct {
	db                  *gorm.DB
	PromocionDataSource PromocionDataSource
}

func NewPromocionDataSourceImpl(db *gorm.DB) PromocionDataSource {
	return PromocionDataSourceImpl{db: db}
}

func (impl PromocionDataSourceImpl) GetAll() ([]models.Promocion, error) {
	var promociones []models.Promocion
	if result := impl.db.Find(&promociones); result.Error != nil {
		return nil, result.Error
	}
	return promociones, nil
}

func (impl PromocionDataSourceImpl) Create(newModel models.Promocion) (models.Promocion, error) {
	if result := impl.db.Create(&newModel); result.Error != nil {
		return models.Promocion{}, result.Error
	}
	return newModel, nil
}
