package usecases

import (
	"github.com/charlieaular/drugstore_backend/domain/repositories"
	models "github.com/charlieaular/drugstore_backend/models"
)

type FacturaUseCase struct {
	FacturaRepository repositories.FacturaRepository
}

func NewFacturaUseCase(FacturaRepository repositories.FacturaRepository) FacturaUseCase {
	return FacturaUseCase{FacturaRepository}
}

func (usecase *FacturaUseCase) GetAll() ([]models.Factura, error) {
	return usecase.FacturaRepository.GetAll()
}

func (usecase *FacturaUseCase) Create(newModel models.Factura, medicamentos []int) (models.Factura, error) {
	return usecase.FacturaRepository.Create(newModel, medicamentos)
}
