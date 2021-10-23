package usecases

import (
	"github.com/charlieaular/drugstore_backend/domain/repositories"
	models "github.com/charlieaular/drugstore_backend/models"
)

type MedicamentoUseCase struct {
	MedicamentoRepository repositories.MedicamentoRepository
}

func NewMedicamentoUseCase(MedicamentoRepository repositories.MedicamentoRepository) MedicamentoUseCase {
	return MedicamentoUseCase{MedicamentoRepository}
}

func (usecase *MedicamentoUseCase) GetAll() ([]models.Medicamento, error) {
	return usecase.MedicamentoRepository.GetAll()
}

func (usecase *MedicamentoUseCase) Create(newModel models.Medicamento) (models.Medicamento, error) {
	return usecase.MedicamentoRepository.Create(newModel)
}
