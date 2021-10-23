package datasource

import (
	model "github.com/charlieaular/drugstore_backend/models"
	"gorm.io/gorm"
)

type MedicamentoDataSource interface {
	GetAll() ([]model.Medicamento, error)
}

type MedicamentoDataSourceImpl struct {
	db                    *gorm.DB
	MedicamentoDataSource MedicamentoDataSource
}

func NewMedicamentoDataSourceImpl(db *gorm.DB) MedicamentoDataSource {
	return MedicamentoDataSourceImpl{db: db}
}

func (impl MedicamentoDataSourceImpl) GetAll() ([]model.Medicamento, error) {
	var medicamentos []model.Medicamento
	if result := impl.db.Find(&medicamentos); result.Error != nil {
		return nil, result.Error
	}
	return medicamentos, nil
}
