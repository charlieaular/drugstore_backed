package model

import (
	"github.com/charlieaular/drugstore_backend/utils"
	"github.com/shopspring/decimal"
)

type Factura struct {
	ID           int32            `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	FechaCrear   utils.CustomDate `gorm:"column:fecha_crear;type:DATE;" json:"fecha_crear"`
	PagoTotal    decimal.Decimal  `gorm:"column:pago_total;type:decimal(6,2);" json:"pago_total" sql:"type:decimal(6,2);"`
	PromocionID  *int             `gorm:"column:promocion_id;type:INT4;" json:"promocion_id,omitempty"`
	Medicamentos *[]Medicamento   `gorm:"many2many:factura_medicamento;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:medicamento_id;jointable_foreignkey:factura_id;" json:"medicamentos,omitempty"`
	Promocion    *Promocion       `gorm:"foreignKey:promocion_id" json:"promocion,omitempty"`
}

func (f *Factura) TableName() string {
	return "factura"
}
