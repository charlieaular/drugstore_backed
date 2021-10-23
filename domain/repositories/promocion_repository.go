package repositories

import models "github.com/charlieaular/drugstore_backend/models"

type PromocionRepository interface {
	GetAll() ([]models.Promocion, error)
	Create(newModel models.Promocion) (models.Promocion, error)
}
