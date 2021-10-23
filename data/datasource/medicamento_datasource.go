package datasource

import (
	models "github.com/charlieaular/drugstore_backend/models"
	"gorm.io/gorm"
)

type MedicamentoDataSource interface {
	GetAll() ([]models.Medicamento, error)
	Create(newModel models.Medicamento) (models.Medicamento, error)
}

type MedicamentoDataSourceImpl struct {
	db                    *gorm.DB
	MedicamentoDataSource MedicamentoDataSource
}

func NewMedicamentoDataSourceImpl(db *gorm.DB) MedicamentoDataSource {
	return MedicamentoDataSourceImpl{db: db}
}

func (impl MedicamentoDataSourceImpl) GetAll() ([]models.Medicamento, error) {
	var medicamentos []models.Medicamento
	if result := impl.db.Find(&medicamentos); result.Error != nil {
		return nil, result.Error
	}
	return medicamentos, nil
}

func (impl MedicamentoDataSourceImpl) Create(newModel models.Medicamento) (models.Medicamento, error) {
	if result := impl.db.Create(&newModel); result.Error != nil {
		return models.Medicamento{}, result.Error
	}
	return newModel, nil
}
