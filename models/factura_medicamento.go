package model

type FacturaMedicamento struct {
	ID            int `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	FacturaID     int `gorm:"column:factura_id;type:INT4;" json:"factura_id"`
	MedicamentoID int `gorm:"column:medicamento_id;type:INT4;" json:"Medicamento_id"`
}

func (f *FacturaMedicamento) TableName() string {
	return "factura_medicamento"
}
