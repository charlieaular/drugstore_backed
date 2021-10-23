package model

import (
	"database/sql"
	"time"
)

type Factura struct {
	ID          int32           `gorm:"primary_key;column:id;type:INT4;"`
	FechaCrear  time.Time       `gorm:"column:fecha_crear;type:DATE;"`
	PagoTotal   sql.NullFloat64 `gorm:"column:pago_total;type:FLOAT8;"`
	PromocionID sql.NullInt64   `gorm:"column:promocion_id;type:INT4;"`
}

// TableName sets the insert table name for this struct type
func (f *Factura) TableName() string {
	return "factura"
}
