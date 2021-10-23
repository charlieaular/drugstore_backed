package usecases

import (
	"github.com/charlieaular/drugstore_backend/domain/repositories"
)

type FacturaUseCase struct {
	FacturaRepository repositories.FacturaRepository
}

func NewFacturaUseCase(FacturaRepository repositories.FacturaRepository) FacturaUseCase {
	return FacturaUseCase{FacturaRepository}
}

func (usecase *FacturaUseCase) GetAll() string {
	message := usecase.FacturaRepository.GetAll()
	return message
}
