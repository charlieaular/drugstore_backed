package usecases

import (
	"github.com/charlieaular/drugstore_backend/domain/repositories"
	model "github.com/charlieaular/drugstore_backend/models"
)

type MedicamentoUseCase struct {
	MedicamentoRepository repositories.MedicamentoRepository
}

func NewMedicamentoUseCase(MedicamentoRepository repositories.MedicamentoRepository) MedicamentoUseCase {
	return MedicamentoUseCase{MedicamentoRepository}
}

func (usecase *MedicamentoUseCase) GetAll() ([]model.Medicamento, error) {
	return usecase.MedicamentoRepository.GetAll()
}
