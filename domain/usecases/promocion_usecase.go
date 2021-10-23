package usecases

import (
	"github.com/charlieaular/drugstore_backend/domain/repositories"
	models "github.com/charlieaular/drugstore_backend/models"
)

type PromocionUseCase struct {
	PromocionRepository repositories.PromocionRepository
}

func NewPromocionUseCase(PromocionRepository repositories.PromocionRepository) PromocionUseCase {
	return PromocionUseCase{PromocionRepository}
}

func (usecase *PromocionUseCase) GetAll() ([]models.Promocion, error) {
	return usecase.PromocionRepository.GetAll()
}

func (usecase *PromocionUseCase) Create(newModel models.Promocion) (models.Promocion, error) {
	return usecase.PromocionRepository.Create(newModel)
}
