package repositories

import model "github.com/charlieaular/drugstore_backend/models"

type MedicamentoRepository interface {
	GetAll() ([]model.Medicamento, error)
}
