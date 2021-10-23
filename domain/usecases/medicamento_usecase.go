package usecases

import (
	"github.com/charlieaular/drugstore_backend/domain/repositories"
)

type MedicamentoUseCase struct {
	MedicamentoRepository repositories.MedicamentoRepository
}

func NewMedicamentoUseCase(MedicamentoRepository repositories.MedicamentoRepository) MedicamentoUseCase {
	return MedicamentoUseCase{MedicamentoRepository}
}

func (usecase *MedicamentoUseCase) GetAll() string {
	message := usecase.MedicamentoRepository.GetAll()
	return message
}
