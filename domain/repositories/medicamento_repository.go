package repositories

import models "github.com/charlieaular/drugstore_backend/models"

type MedicamentoRepository interface {
	GetAll() ([]models.Medicamento, error)
	Create(newModel models.Medicamento) (models.Medicamento, error)
}
