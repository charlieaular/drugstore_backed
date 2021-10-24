package model

import (
	"github.com/charlieaular/drugstore_backend/utils"
)

type Factura struct {
	ID          int32            `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	FechaCrear  utils.CustomDate `gorm:"column:fecha_crear;type:DATE;" json:"fecha_rear"`
	PagoTotal   float64          `gorm:"column:pago_total;type:FLOAT8;" json:"pago_total"`
	PromocionID int              `gorm:"column:promocion_id;type:INT4;" json:"promocion_id"`
}

// TableName sets the insert table name for this struct type
func (f *Factura) TableName() string {
	return "factura"
}
