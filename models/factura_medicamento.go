package model

type FacturaMedicamento struct {
	ID            int32 `gorm:"primary_key;column:id;type:INT4;"`
	FacturaID     int32 `gorm:"column:factura_id;type:INT4;"`
	MedicamentoID int32 `gorm:"column:medicamento_id;type:INT4;"`
}

func (f *FacturaMedicamento) TableName() string {
	return "factura_medicamento"
}
