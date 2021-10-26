package repositories

import (
	models "github.com/charlieaular/drugstore_backend/models"
)

type FacturaRepository interface {
	GetAll() ([]models.Factura, error)
	Create(newModel models.Factura, medicamentos []int) (models.Factura, error)
	GetGrafica(fechaInicio string, fechaFin string) ([]models.Factura, error)
}
